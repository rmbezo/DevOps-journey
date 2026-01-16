# Monitoring

```
kubectl create ns observability

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm upgrade --install kps prometheus-community/kube-prometheus-stack \
  -n observability
```

- Пароль для Grafana admin:
```
kubectl -n observability get secret kps-grafana -o jsonpath='{.data.admin-password}{"\n"}' | base64 -d ; echo
```

- Зайти на наш grafana: 

http://grafana.local.lab:<PORT_TRAEFIK> перед этим у себя в /etc/hosts/ добавить любую ноду 192.0.2.12   grafana.local.lab



## Grafana Loki

```
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

helm upgrade --install loki grafana/loki-stack \
  -n observability \
  --set grafana.enabled=false \
  --set promtail.enabled=true
```