# How to launch code:

```
kubectl apply -f namespaces/dev.yaml

kubectl apply -f nginx/deployment.yaml -n dev
kubectl apply -f nginx/service.yaml -n dev


kubectl apply -f nginx/ingress.yaml -n dev
```

before or after that install ingress controler, in my cluster i use traefik. You can see how to do it in ./ingress.md