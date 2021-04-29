# Knative Sample Controller

[![GoDoc](https://godoc.org/knative.dev/sample-controller?status.svg)](https://godoc.org/knative.dev/sample-controller)
[![Go Report Card](https://goreportcard.com/badge/knative/sample-controller)](https://goreportcard.com/report/knative/sample-controller)

Knative `sample-controller` defines a few simple resources that are validated by
webhook and managed by a controller to demonstrate the canonical style in which
Knative writes controllers.

To learn more about Knative, please visit our
[Knative docs](https://github.com/knative/docs) repository.

If you are interested in contributing, see [CONTRIBUTING.md](./CONTRIBUTING.md)
and [DEVELOPMENT.md](./DEVELOPMENT.md).



cat <<EOF | kubectl apply -f -
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: factor-viewer
  labels:
    duck.tableflip.dev/results: "true"
rules:
- apiGroups:
  - samples.knative.dev
  resources:
  - factors
  - factors/status
  verbs:
  - get
  - list
  - watch
EOF