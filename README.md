# Kubernetes Pod Tagger

## What?

It's a simple pod that checks if extra labels needs to be added to Kubernetes pods or not.

## How?

Manually add the pods in the dictionary:

```
var putLabels = map[string]string{
	"client-type": "internal",
	"label-key": "label-value",
}
```

## Deploy


```sh
kubectl apply -f kubernetes-resources.yaml
```