package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Constants
const (
	ServerPort      = ":8081"
	MinResults      = 1
	MaxResults      = 5
	MinSleepSeconds = 3
	MaxSleepSeconds = 10
)

// ScannerRequest represents the input payload for the scanner
type ScannerRequest struct {
	Provider             string            `json:"provider"`
	Regions              []string          `json:"regions"`
	ComplianceFrameworks []string          `json:"compliance_frameworks"`
	Credentials          map[string]string `json:"credentials"`
	OutputFormat         string            `json:"output_format"`
}

// ProwlerResult represents a single finding from the scanner
type ProwlerResult struct {
	CheckID      string   `json:"check_id"`
	CheckTitle   string   `json:"check_title"`
	Service      string   `json:"service"`
	Region       string   `json:"region"`
	ResourceID   string   `json:"resource_id"`
	Severity     string   `json:"severity"`
	Remediation  string   `json:"remediation"`
	ResourceTags []string `json:"resource_tags,omitempty"`
}

var sampleResults = []ProwlerResult{
	{"iam_user_console_access_unused", "Ensure IAM users with console access have used it in the last 45 days", "iam", "us-east-1", "user1", "LOW", "Disable unused console access", nil},
	{"s3_bucket_public_access", "Ensure S3 buckets have public access blocked", "s3", "eu-west-1", "my-public-bucket", "CRITICAL", "Enable S3 Public Access Block", nil},
	{"ec2_ebs_volume_encryption", "Ensure EBS volumes are encrypted", "ec2", "us-east-1", "vol-12345678", "MEDIUM", "Encrypt EBS volumes", nil},
	{"rds_instance_public_access", "Ensure RDS instances are not publicly accessible", "rds", "eu-west-1", "db-prod", "HIGH", "Disable public access for RDS", nil},
	{"vpc_flow_logs_enabled", "Ensure VPC flow logs are enabled", "vpc", "us-east-1", "vpc-abc123", "LOW", "Enable VPC flow logs", nil},
	{"cloudwatch_log_group_retention", "Ensure CloudWatch log groups have a retention period", "logs", "us-east-1", "/aws/lambda/my-func", "LOW", "Set log group retention", nil},
	{"iam_mfa_enabled", "Ensure MFA is enabled for all IAM users", "iam", "global", "root-account", "CRITICAL", "Enable MFA", nil},
	{"elb_ssl_listeners", "Ensure ELB listeners use SSL", "elb", "eu-west-1", "my-elb", "HIGH", "Add SSL listener", nil},
	{"kms_key_rotation", "Ensure KMS keys have rotation enabled", "kms", "us-east-1", "key-uuid", "MEDIUM", "Enable KMS rotation", nil},
	{"lambda_function_vpc_config", "Ensure Lambda functions are in a VPC", "lambda", "us-east-1", "my-lambda", "MEDIUM", "Configure VPC for Lambda", nil},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Endpoint
	e.POST("/v1/run/scanner", runScanner)

	// Start server
	log.Printf("Starting scanner simulation server on %s...", ServerPort)
	e.Logger.Fatal(e.Start(ServerPort))
}

func runScanner(c echo.Context) error {
	req := new(ScannerRequest)
	if err := c.Bind(req); err != nil {
		log.Printf("Error binding request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// Log request payload
	reqJSON, _ := json.Marshal(req)
	log.Printf("[SCANNER] New Request Received: %s", string(reqJSON))

	// Simulation: Sleep for random seconds
	sleepDuration := rand.Intn(MaxSleepSeconds-MinSleepSeconds+1) + MinSleepSeconds
	log.Printf("[SCANNER] Simulating processing for %d seconds...", sleepDuration)
	time.Sleep(time.Duration(sleepDuration) * time.Second)

	// Simulation: Pick random number of results
	numResults := rand.Intn(MaxResults-MinResults+1) + MinResults
	results := make([]ProwlerResult, numResults)

	// Shuffle and pick
	perm := rand.Perm(len(sampleResults))
	for i := 0; i < numResults; i++ {
		results[i] = sampleResults[perm[i]]
	}

	// Log response
	resJSON, _ := json.Marshal(results)
	log.Printf("[SCANNER] Sending Response: %s", string(resJSON))

	return c.JSON(http.StatusOK, results)
}
