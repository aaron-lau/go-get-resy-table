# terraform/main.tf
terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_cloud_run_service" "go_get_resy_table" {
  name     = "go-get-resy-table"
  location = var.region

  template {
    spec {
      containers {
        image = var.container_image
        env {
          name  = "RESY_API_KEY"
          value = var.resy_api_key
        }
        env {
          name  = "RESY_AUTH_KEY"
          value = var.resy_auth_key
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

# IAM policy to make the service public
resource "google_cloud_run_service_iam_member" "public" {
  service  = google_cloud_run_service.go_get_resy_table.name
  location = google_cloud_run_service.go_get_resy_table.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}