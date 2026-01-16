# kubeadm kubelet kubectl установка

```
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.32/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg

echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.32/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
```

## Обновить и установить + фиксация версии
```
sudo apt update
sudo apt install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

Если установить kubelet kubeadm kubectl не получается используем VPN или прокси.

Я сделал просто, через свой пк установил пакеты и переместил их с помощью scp:
scp ~/k8s-packages/*.deb root@ip:/tmp/

## 4. Включить и перезапустить kubelet (оно может быть неактивно до установки kubeadm)
sudo systemctl enable kubelet
sudo systemctl start kubelet