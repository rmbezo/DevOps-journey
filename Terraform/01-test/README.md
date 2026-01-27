# Terraform IaC

I practice terraform with clouds and libvirt on my server

- for server configuration you need to:
1. Install KVM/Libvirt
2. Install Terraform
3. Practice with terraform 



```
terraform init
terraform plan -var-file="secrets.tfvars"
terraform apply -var-file="secrets.tfvars"
terraform destroy 
```


FOr RU region :

- Change the path of downloading registry.terraform.io

```
cat > ~/.terraformrc << 'EOF'
provider_installation {
  network_mirror {
    url = "https://terraform-mirror.yandexcloud.net/"
    include = ["registry.terraform.io/*/*"]
  }
  direct {
    exclude = ["registry.terraform.io/*/*"]
  }
}
EOF
```