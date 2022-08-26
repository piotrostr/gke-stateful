data "google_client_config" "provider" {}

resource "google_container_cluster" "cluster" {
  name               = "cluster"
  location           = "us-central1-a"
  initial_node_count = 1

  cluster_autoscaling {
    enabled = true

    resource_limits {
      resource_type = "cpu"
      minimum = 1
      maximum = 8
    }
    resource_limits {
      resource_type = "memory"
      minimum = 2
      maximum = 16
    }
  }
}

provider "kubernetes" {
  host  = "https://${google_container_cluster.cluster.endpoint}"
  token = data.google_client_config.provider.access_token
  cluster_ca_certificate = base64decode(
    google_container_cluster.cluster.master_auth.0.cluster_ca_certificate
  )
}
