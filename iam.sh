#!/bin/bash
cd ./terraform
AWS_REGION=eu-west-1 terraformer import aws --resources=iam --profile="$1"