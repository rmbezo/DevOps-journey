# First pod

Let's create our first pod, with this command:

```
kubectl run nginx-pod \
--image=nginx:alpine \
--restart=Never
```

## deployment

```
kubectl create deployment nginx \
--image=nginx:latest
```

