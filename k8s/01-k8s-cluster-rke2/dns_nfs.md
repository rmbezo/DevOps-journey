![alt text](./images/image2.png)

# NFS:

```
    sudo apt update && sudo apt install nfs-kernel-server -y
    sudo mkdir -p /srv/nfs/k8s
    sudo chown nobody:nogroup /srv/nfs/k8s
    sudo chmod 0777 /srv/nfs/k8s


    echo "/srv/nfs/k8s 192.168.0.0/24(rw,sync,no_subtree_check,no_root_squash)" | sudo tee -a /etc/exports


    sudo exportfs -ra
    sudo systemctl restart nfs-kernel-server
    sudo systemctl enable nfs-kernel-server
```

- Проверка:

```
sudo apt install nfs-common -y
sudo mkdir /mnt/nfs-test
sudo mount -t nfs 192.168.0.100:/srv/nfs/k8s /mnt/nfs-test
echo "TEST" > /mnt/nfs-test/test.txt
```

В папке на сервере появится файл test.txt в директории /mnt/nfs-test/


# DNS: 

```
sudo apt install bind9 bind9utils -y
```

Только для ipv4:
```
sudo sed -i 's/OPTIONS="/OPTIONS="-4 /' /etc/default/named
sudo systemctl restart bind9
```

```
sudo nano /etc/bind/named.conf.local
```

внутри:

```
zone "lab.local" {
    type master;
    file "/etc/bind/db.lab.local";
};

```

```
sudo cp /etc/bind/db.local /etc/bind/db.lab.local
sudo nano /etc/bind/db.lab.local
```

внутри:

```
$TTL    604800
@       IN      SOA     ns.lab.local. root.lab.local. (
                              3         ; Serial
                           604800         ; Refresh
                          86400         ; Retry
                        2419200         ; Expire
                         604800 )       ; Negative Cache TTL
;
@       IN      NS      ns.lab.local.
ns          IN      A       192.0.168.0
control     IN      A       192.0.168.0
worker1     IN      A       192.0.168.0
worker2     IN      A       192.0.168.0
nfs         IN      A       192.0.168.0
```

Применить:

```
sudo named-checkconf
sudo named-checkzone lab.local /etc/bind/db.lab.local
sudo systemctl reload bind9
```


На каждой ВМ:

```
echo "nameserver 10.10.10.154" | sudo tee /etc/resolv.conf 

```

С ВМ проверить др. сервера:

```
ping worker1.lab.local
```


Cхема:

    [Физ.сервер  192.168.0.0]
    │ DNS:53  │ NFS         │
    │         │             │
    └─────────┼─────────────┘
              │ сеть 192.168.0.0/24
              ▼
    ┌─────────┼─────────┼─────────┐
    │control  │worker1  │worker2  │
    │192.168.0│20.30.10.│40.50.60.│
    │         │         │         │
    │RKE2-srv │RKE2-agt │RKE2-agt │
    └─────────┼─────────┼─────────┘