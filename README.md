# Common Network Policy Operator

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/bells17/common-network-policy-operator)](https://goreportcard.com/report/github.com/bells17/common-network-policy-operator)

Common Policy Operator is auto create common network policy to all namespaces.

## Requirement

- kubectl
- Kustomize

## Installation

```
$ kubectl config current-context <TARGET CLUSTER CONTEXT>
$ kubectl apply -f https://raw.githubusercontent.com/bells17/common-network-policy-operator/1.0.2/config/deploy.yaml
```

If install is success, you can check resources as below:

```
$ kubectl get po -n common-network-policy-operator-system
NAME                                                     READY     STATUS    RESTARTS   AGE
common-network-policy-operator-controller-manager-0   1/1       Running   0          1m

$ kubectl get crd
NAME                                                           CREATED AT
commonnetworkpolicies.commonnetworkpolicies.bells17.io   2018-11-09T00:00:00Z
```

## Usage

You can apply to your cluster using the following example.

```yaml
apiVersion: commonnetworkpolicies.bells17.io/v1alpha1
kind: CommonNetworkPolicy
metadata:
  labels:
    controller-tools.k8s.io: "1.0"
  name: sample-networkpolicy
spec:
  namePrefix: common
  excludeNamespaces:
  - kube-system
  - common-network-policy-operator-system
  policySpec:
    podSelector: {}
    policyTypes:
    - Egress

---
apiVersion: commonnetworkpolicies.bells17.io/v1alpha1
kind: CommonNetworkPolicy
metadata:
  labels:
    controller-tools.k8s.io: "1.0"
  name: sample-networkpolicy2
spec:
  namePrefix: common
  excludeNamespaces:
  - kube-system
  - common-network-policy-operator-system
  policySpec:
    podSelector: {}
    ingress:
    - {}
```

After apply, create commonnetworkpolicies and networkpolicies such as below:

```
$ kubectl get commonnetworkpolicies
NAME                    CREATED AT
sample-networkpolicy    1m
sample-networkpolicy2   1m

$ kubectl get networkpolicies --all-namespaces=true
NAMESPACE     NAME                              POD-SELECTOR   AGE
default       common-sample-networkpolicy    <none>         1m
default       common-sample-networkpolicy2   <none>         1m
docker        common-sample-networkpolicy    <none>         1m
docker        common-sample-networkpolicy2   <none>         1m
kube-public   common-sample-networkpolicy    <none>         1m
kube-public   common-sample-networkpolicy2   <none>         1m
```

## LICENSE

Copyright 2018 bells17.

Licensed under the MIT License.
