
```
    # установить нужные пакеты
    sudo apt install -y apt-transport-https ca-certificates curl gnupg lsb-release

    # отключить swap 
    sudo swapoff -a
    # чтобы отключить навсегда - закомментировать swap в /etc/fstab:
    sudo sed -i.bak '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab

    # сетевые настройки для k8s
    cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
    net.bridge.bridge-nf-call-iptables  = 1
    net.bridge.bridge-nf-call-ip6tables = 1
    net.ipv4.ip_forward                 = 1
    EOF
sudo sysctl --system
```

```
    # добавить репозиторий Docker (для стабильного containerd)
    curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
    echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] \
    https://download.docker.com/linux/debian $(lsb_release -cs) stable" \
    | sudo tee /etc/apt/sources.list.d/docker.list

    sudo apt update
    sudo apt install -y containerd.io

    # сгенерировать конфиг и включить systemd cgroup
    sudo mkdir -p /etc/containerd
    sudo containerd config default | sudo tee /etc/containerd/config.toml
    sudo sed -i 's/SystemdCgroup = false/SystemdCgroup = true/' /etc/containerd/config.toml

    sudo systemctl restart containerd
    sudo systemctl enable containerd

```

Проверка:
```
    sudo systemctl status containerd
    containerd --version
```