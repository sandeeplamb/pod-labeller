# pod-labeller Helm Chart

This Helm Chart will create the pod-labeller Pod to create EBS tags on PVC

## Installation Commands

You can install the chart directly

```
helm upgrade --install pod-labeller ./pod-labeller/ --namespace=kube-system --debug --dry-run
helm upgrade --install pod-labeller ./pod-labeller/ --namespace=kube-system --debug 
```

If Helm-Chart is in a repo

```
helm upgrade --install pod-labeller chartmuseum/pod-labeller/ --namespace=kube-system --debug --dry-run
helm upgrade --install pod-labeller chartmuseum/pod-labeller/ --namespace=kube-system --debug 
```