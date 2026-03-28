#!/bin/bash

curl -X POST https://fiapchallengecspm.gariel.cloud/v1/run/scanner \
  -H "Content-Type: application/json" \
  -d '{
    "provider": "aws",
    "regions": [
      "us-east-1",
      "eu-west-1"
    ],
    "compliance_frameworks": [
      "cis",
      "pci-dss"
    ],
    "credentials": {
    },
    "output_format": "json"
  }'
