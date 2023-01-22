#!/bin/bash

# docker push nightmanager/random-identity-http-server

docker tag \
    nightmanager/random-identity-http-server \
    "$MY_AWS_ACCOUNT_NUMBER".dkr.ecr.eu-west-1.amazonaws.com/random-identity-http-server

# Login to ECR
aws ecr get-login-password --region eu-west-1 | \
    docker login --username AWS --password-stdin "$MY_AWS_ACCOUNT_NUMBER".dkr.ecr.eu-west-1.amazonaws.com

docker push "$MY_AWS_ACCOUNT_NUMBER".dkr.ecr.eu-west-1.amazonaws.com/random-identity-http-server
