[![Coverage Status](https://coveralls.io/repos/github/kubernetes-sigs/aws-alb-ingress-controller/badge.svg?branch=master)](https://coveralls.io/github/kubernetes-sigs/aws-alb-ingress-controller?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kubernetes-sigs/aws-alb-ingress-controller)](https://goreportcard.com/report/github.com/kubernetes-sigs/aws-alb-ingress-controller)
[![Build Status](https://travis-ci.org/kubernetes-sigs/aws-alb-ingress-controller.svg?branch=master)](https://travis-ci.org/kubernetes-sigs/aws-alb-ingress-controller)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcoreos%2Falb-ingress-controller.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcoreos%2Falb-ingress-controller?ref=badge_shield)

# AWS Load Balancer Controller
We have rebranded "AWS ALB Ingress Controller" to be "AWS Load Balancer Controller" along with our v2.0.0 release.

"AWS ALB Ingress Controller" is now deprecated, we encourage users to migrate to "AWS Load Balancer Controller" by follow our [migration guide](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.2/deploy/upgrade/migrate_v1_v2/).

# AWS ALB Ingress Controller

**NOTE:** The current image version is `v1.1.9`. Please file any issues you find and note the version used.

The AWS ALB Ingress Controller satisfies Kubernetes [ingress resources](https://kubernetes.io/docs/user-guide/ingress) by provisioning [Application Load Balancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/introduction.html).

This project was originated by [Ticketmaster](https://github.com/ticketmaster) and [CoreOS](https://github.com/coreos) as part of Ticketmaster's move to AWS and CoreOS Tectonic. Learn more about Ticketmaster's Kubernetes initiative from Justin Dean's video at [Tectonic Summit](https://www.youtube.com/watch?v=wqXVKneP0Hg).

This project was donated to Kubernetes SIG-AWS to allow AWS, CoreOS, Ticketmaster and other SIG-AWS contributors to officially maintain the project. SIG-AWS reached this consensus on June 1, 2018.

## Documentation

Checkout our [Live Docs](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v1.1/)!

## Getting started

To get started with the controller, see our [walkthrough](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v1.1/guide/walkthrough/echoserver/).

## Setup
- See [controller setup](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v1.1/guide/controller/setup/) on how to install ALB ingress controller
- See [external-dns setup](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v1.1/guide/external-dns/setup/) for how to setup the external-dns to manage route 53 records.

## Building

For details on building this project, see our [building guide](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v1.1/BUILDING/).

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

- [Slack channel](https://kubernetes.slack.com/messages/sig-aws)
- [Mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-aws)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcoreos%2Falb-ingress-controller.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcoreos%2Falb-ingress-controller?ref=badge_large)
