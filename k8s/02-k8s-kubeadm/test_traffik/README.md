# Test how traffik works

- 2 nginx web servers with ip show:

exec inside pods and write inside /usr/share/nginx/html/index.html ip of node. Or use some nginx configure env's

```
kubectl exec -it <NAME_POD> -n dev -- sh
```


My example:


![example](/k8s/images/image4.png)


![example](/k8s/images/image5.png)