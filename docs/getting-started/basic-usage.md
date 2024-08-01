# Basic Usage

Before you get started, you will need to export K8s credentials that you want methodk8s to utilize as environment variables. 

## Binaries

Running as a binary means you don't need to do anything additional for methodk8s to leverage the environment variables you have already exported. You can test that things are working properly by running:

```bash
methodk8s pod enumerate --context minikube --path ~/.kube/config
```

## Docker

Running an authenticated workflow with methodk8s as a Docker container requires that you pass k8s credentials to the container. This can be done in following 2 ways
1. Either by creating a service account and passing in the secrets

```bash
helm install my-release ./k8s-helm-chart
docker run -e SERVICE_ACCOUNT_TOKEN="XXXX" -e SERVER_API="https://ClusterURL.com" -e CA_CERT="XXXX" methodsecurity/methodk8s
```

2. By mounting the k8s config file

```bash
docker run -v /path/to/your/kubeconfig:/opt/method/methodk8s/kubeconfig methodsecurity/methodk8s
```
