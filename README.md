# Prometheus Example App

This example app serves as an example of how one can easily instrument HTTP handlers with [Prometheus][prometheus] metrics. It uses the Prometheus [go client][client-golang] to create a new Prometheus registry.

Usage is simple, on any request to `/hello` the request will result in a `200` response code. This increments the counter for this response code. Similarly the `/error` endpoint will result in a `404` response code, therefore increments that respective counter.

A Docker image is available at: `quay.io/owensengoku/prometheus-example-app:v0.1.0`

[prometheus]:https://prometheus.io/
[client-golang]:https://github.com/prometheus/client_golang
