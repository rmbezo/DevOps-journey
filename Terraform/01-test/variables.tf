variable "libvirt_uri" {
  description = "Libvirt connection URI для удалённого подключения"
  type        = string
}

variable "vm_count" {
  description = "Количество VM"
  type        = number
  default     = 3
}

variable "vm_names" {
  description = "Список имён VM"
  type        = list(string)
  default     = ["k8s-master", "k8s-worker1", "k8s-worker2"]
}

variable "vm_memory" {
  description = "Память VM в MB"
  type        = number
  default     = 2048
}

variable "vm_vcpu" {
  description = "Количество vCPU для VM"
  type        = number
  default     = 2
}

variable "vm_disk_size" {
  description = "Размер диска VM в байтах (5GB = 5368709120)"
  type        = number
  default     = 5368709120  # 5GB
}