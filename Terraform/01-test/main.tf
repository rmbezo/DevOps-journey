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

# Создаём диски для каждой VM
resource "libvirt_volume" "vm_disks" {
  count  = var.vm_count
  name   = "${var.vm_names[count.index]}-disk.qcow2"
  pool   = "default"
  size   = var.vm_disk_size
  format = "qcow2"
}

# Создаём сами VM
resource "libvirt_domain" "vms" {
  count     = var.vm_count
  name      = var.vm_names[count.index]
  memory    = var.vm_memory
  vcpu      = var.vm_vcpu

  disk {
    volume_id = libvirt_volume.vm_disks[count.index].id
  }

  network_interface {
    bridge = "virbr0"
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