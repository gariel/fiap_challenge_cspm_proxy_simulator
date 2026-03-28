#!/bin/bash

curl -X POST http://localhost:8081/v1/run/scanner \
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
      "access_key_id": "AKIAXXXXXXXXXXXXXXXX",
      "secret_access_key": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    },
    "output_format": "json"
  }'
