resource "google_container_cluster" "cluster" {
  name               = "stateful-cluster"
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
