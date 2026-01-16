# Monitoring

```
kubectl create ns observability

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm upgrade --install kps prometheus-community/kube-prometheus-stack \
  -n observability
```