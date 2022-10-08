# Automatic-scaling-pods
Application that scaling automatically pods and deployments in the Kubernetes.

```
kind.exe create cluster --config .\cluster.yaml

kubectl.exe apply -f .\enable_metrics_kind.yaml
go mod tidy
go run .\main.go
```

DECIMAL SI

Memory:
130792Ki    like    127Mi
115Mi       like    115Mi

CPU:
192106837n  like    190m
1000474931n like    1000m (1CPU)