terraform {
  required_providers {
    libvirt = {
      source  = "dmacvicar/libvirt"
      version = "~> 0.7.6"
    }
  }
}

provider "libvirt" {
  uri = var.libvirt_uri
}

# Создаём диски для каждой VM на основе базового образа
resource "libvirt_volume" "vm_disks" {
  count          = var.vm_count
  name           = "${var.vm_names[count.index]}-disk.qcow2"
  pool           = "default"

  # Ссылаемся на точное имя файла, который уже лежит в /var/lib/libvirt/images/ на сервере
  base_volume_name = "jammy-server-cloudimg-amd64.img" 
  
  size           = var.vm_disk_size
  format         = "qcow2"
}

# Cloud-init конфигурация
#data "template_file" "user_data" {
 # count    = var.vm_count
#  template = file("${path.module}/cloud_init.cfg")
 # vars = {
 #   hostname = var.vm_names[count.index]
#  }
#}

#resource "libvirt_cloudinit_disk" "commoninit" {
#  count     = var.vm_count
#  name      = "${var.vm_names[count.index]}-cloudinit.iso"
#  user_data = data.template_file.user_data[count.index].rendered
#  pool      = "default"
#}

# Создаём сами VM
resource "libvirt_domain" "vms" {
  count     = var.vm_count
  name      = var.vm_names[count.index]
  memory    = var.vm_memory
  vcpu      = var.vm_vcpu

 # cloudinit = libvirt_cloudinit_disk.commoninit[count.index].id

  disk {
    volume_id = libvirt_volume.vm_disks[count.index].id
  }

  network_interface {
    bridge = "virbr0"
    wait_for_lease = false
  }

  console {
    type        = "pty"
    target_type = "serial"
    target_port = "0"
  }

  graphics {
    type     = "spice"
    autoport = true
  }

  autostart = true
}

# Вывод IP-адресов
output "vm_ips" {
  value = {
    for idx, name in var.vm_names :
    name => libvirt_domain.vms[idx].network_interface[0].addresses
  }
}