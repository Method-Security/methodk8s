<div align="center">
<h1>methodk8s</h1>

[![GitHub Release][release-img]][release]
[![Verify][verify-img]][verify]
[![Go Report Card][go-report-img]][go-report]
[![License: Apache-2.0][license-img]][license]

[![GitHub Downloads][github-downloads-img]][release]
[![Docker Pulls][docker-pulls-img]][docker-pull]

</div>

methodk8s provides security operators with a number of data-rich K8S enumeration capabilities to help them gain visibility into their K8s Cluster. Designed with data-modeling and data-integration needs in mind, methodk8s can be used on its own as an interactive CLI, orchestrated as part of a broader data pipeline, or leveraged from within the Method Platform.

The number of security-relevant K8s resources that methodk8s can enumerate are constantly growing. For the most up to date listing, please see the documentation [here](docs-capabilities)

To learn more about methodk8s, please see the [Documentation site](https://method-security.github.io/methodk8s/) for the most detailed information.

## Quick Start

### Get methodk8s

For the full list of available installation options, please see the [Installation](./docs/getting-started/index.md) page. For convenience, here are some of the most commonly used options:

- `docker run methodsecurity/methodk8s`
- `docker run ghcr.io/method-security/methodk8s:0.0.1`
- Download the latest binary from the [Github Releases](releases) page
- [Installation documentation](./docs/getting-started/index.md)

### Authentication
For authenticated workflows, you need to pass in your kube config file to the docker container

```bash
docker run -v /path/to/your/kubeconfig:/opt/method/methodk8s/kubeconfig methodsecurity/methodk8s
```

### General Usage

```bash
methodk8s <resource> enumerate 
```

#### Examples

```bash
methodk8s pod enumerate --url test-cluster.net
```

```bash
methodk8s node enumerate --context minikube --path ~/.kube/config
```

## Contributing

Interested in contributing to methodk8s? Please see our [Contribution](#) page.

## Want More?

If you're looking for an easy way to tie methodk8s into your broader cybersecurity workflows, or want to leverage some autonomy to improve your overall security posture, you'll love the broader Method Platform.

For more information, see [https://method.security]

## Community

methodk8s is a Method Security open source project.

Learn more about Method's open source source work by checking out our other projects [here](github-org).

Have an idea for a Tool to contribute? Open a Discussion [here](discussion).

[verify]: https://github.com/Method-Security/methodk8s/actions/workflows/verify.yml
[verify-img]: https://github.com/Method-Security/methodk8s/actions/workflows/verify.yml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/Method-Security/methodk8s
[go-report-img]: https://goreportcard.com/badge/github.com/Method-Security/methodk8s
[release]: https://github.com/Method-Security/methodk8s/releases
[releases]: https://github.com/Method-Security/methodk8s/releases/latest
[release-img]: https://img.shields.io/github/release/Method-Security/methodk8s.svg?logo=github
[github-downloads-img]: https://img.shields.io/github/downloads/Method-Security/methodk8s/total?logo=github
[docker-pulls-img]: https://img.shields.io/docker/pulls/methodsecurity/methodk8s?logo=docker&label=docker%20pulls%20%2F%20methodk8s
[docker-pull]: https://hub.docker.com/r/methodsecurity/methodk8s
[license]: https://github.com/Method-Security/methodk8s/blob/main/LICENSE
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[homepage]: https://method.security
[docs-home]: https://method-security.github.io/methodk8s
[docs-capabilities]: https://method-security.github.io/methodk8s/docs/index.html
[discussion]: https://github.com/Method-Security/methodk8s/discussions
[github-org]: https://github.com/Method-Security