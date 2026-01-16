# Observability Grafana + Prometheus + Loki

- 1. Namespaces (namespaces.yaml)
- 2. Traefik (traefik-values.yaml):
```
    helm repo add traefik https://traefik.github.io/charts
    helm repo update

    helm upgrade --install traefik traefik/traefik -n traefik -f k8s/observability/traefik-values.yaml
```
- 3. Kube-prometheus-stack helm (prometheus + grafana + alertmanager + exporters) :
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```

```
helm upgrade --install kps prometheus-community/kube-prometheus-stack \
  -n observability -f kps-values.yaml
```

- 4. Ingress for Grafana:

```
kubectl apply -f ingress-grafana.yaml
```

- 5. web grafana enter :

change /etc/hosts : <ANY_NODE_IP>   grafana.local.lab
in url type: http://grafana.local.lab:<TRAEFIK_PORT>


- 6. Loki + Promtail :

```
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
```

```
helm upgrade --install loki grafana/loki -n observability -f loki-values.yaml
```

- 7. Promtail :

```
helm upgrade --install promtail grafana/promtail -n observability -f promtail-values.yaml
```

