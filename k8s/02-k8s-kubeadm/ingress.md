# ingress

Traefik:

```
helm repo add traefik https://helm.traefik.io/traefik
helm repo update
```

```
kubectl create namespace traefik
```

```
helm install traefik traefik/traefik \
  --namespace traefik
```