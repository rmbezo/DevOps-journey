# How to launch code:

```
kubectl apply -f namespaces/dev.yaml

kubectl apply -f nginx/deployment.yaml -n dev
kubectl apply -f nginx/service.yaml -n dev


kubectl apply -f nginx/ingress.yaml -n dev
```

before or after that install ingress controler, in my cluster i use traefik. You can see how to do it in ./ingress.md


## To open your nginx in browser from traefik

u have to change /etc/hosts and put an ip of any node cluster nginx.local.lab , then open in browser: nginx.local.lab:<TRAEFIK_PORT>

- to check traefik port use: kubectl get svc traefik -n traefik





# Test how traffik works and linux kernel of course 

- 2 nginx web servers with ip show:

exec inside pods and write inside /usr/share/nginx/html/index.html ip of node. Or use some nginx configure env's

```
kubectl exec -it <NAME_POD> -n dev -- sh
```


My example:


![example](/k8s/images/image4.png)


![example](/k8s/images/image5.png)