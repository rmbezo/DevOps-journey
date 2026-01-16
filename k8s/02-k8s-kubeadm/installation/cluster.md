# CLuster

![alt text](/k8s/images/image.png)


## Calico

wget https://raw.githubusercontent.com/projectcalico/calico/v3.27.3/manifests/calico.yaml
scp calico.yaml root@ip:/tmp/

kubectl apply -f /tmp/calico.yaml


![alt text](/k8s/images/image2.png)


Мы подняли наш кластер:


![alt text](/k8s/images/image3.png)