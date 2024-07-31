#!/bin/bash

# Variables
SECRET_NAME=method-service-account-secret
NAMESPACE=default

# Extract the token from the secret
TOKEN=$(kubectl get secret $SECRET_NAME -n $NAMESPACE -o jsonpath='{.data.token}' | base64 --decode)
echo -e "Token\n$TOKEN"

# Get the Kubernetes API server URL
API_SERVER=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}')
echo -e "API Server\n$API_SERVER"

CA_CERT=$(kubectl get secret method-service-account-secret -o jsonpath='{.data.ca\.crt}' | base64 --decode)
echo -e "CA CERT\n$CA_CERT"
