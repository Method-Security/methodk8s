# Basic Usage

Before you get started, you will need to export K8s credentials that you want methodk8s to utilize as environment variables. 
**********TODO***************

## Binaries

Running as a binary means you don't need to do anything additional for methodaws to leverage the environment variables you have already exported. You can test that things are working properly by running:

```bash
methodk8s pod enumerate --context minikube --path ~/.kube/config
```

## Docker

Running methodk8s within a Docker container requires that you pass the k8s credential environment variables into the container. This can be done with the following command:

```bash
docker run \
```
**********TODO***************

