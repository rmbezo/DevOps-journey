# Helm 

- Kubernetes packet manager

## Instalation 

```
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
```

# Chart's

Install chart and start:

```
helm install app 01-first-chart/
```

- With different settings:
```
helm install app1 01-first-chart/ --set container.image=adv4000/k8sphp:version1 --set repclicaCount=3
```

- so word (app1) after install == .Release.Name var in helm

- with --set you can change the values .


List all charts:

```
helm ls
```



## Prod Dev 

```
helm install app test-chart/ -f prod_values
```

.