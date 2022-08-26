# gke-stateful (WIP)

An example of a stateful cluster with a load balancer. Packed with terraform
and skaffold for buttery smooth developer experience.

This cluster scales in and out automatically as per the terraform
configuration using the vertical autoscaler. GKE offers it as an out-of-the-box
solution.

## Disclaimer

Same as [here](https://github.com/piotrostr/gke-gpu-tf/).

## Installation

Provisioning of the resources happens using

```sh
$ terraform apply
```

in the `terraform` directory. Having created the cluster, get the credentials

```sh
$ gcloud container clusters get-credentials cluster --region=us-central1-a
```

and run

```sh
skaffold apply
```

to create the resources in the GKE cluster.

## License

MIT
