# terraform/variables.tf
variable "project_id" {
  description = "Google Cloud Project ID"
  type        = string
}

variable "region" {
  description = "Google Cloud Region"
  type        = string
  default     = "us-central1"
}

variable "container_image" {
  description = "Container image URL"
  type        = string
}

variable "resy_api_key" {
  description = "Resy API Key"
  type        = string
  sensitive   = true
}

variable "resy_auth_key" {
  description = "Resy Auth Key"
  type        = string
  sensitive   = true
}