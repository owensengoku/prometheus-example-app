package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	host              string
	httpRequestsTotal *prometheus.CounterVec
	version           prometheus.Gauge
)

func init() {
	var err error
	host, err = os.Hostname()
	if err == nil {
		log.Printf("Service on host:%s", host)
	} else {
		log.Printf("%s", err)
	}
	version = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"version": "v0.1.0",
		},
	})
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
		ConstLabels: map[string]string{
			"version": "v0.1.0",
			"host":    host,
		},
	}, []string{"code", "method"})
}

func main() {
	bind := ""
	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagset.StringVar(&bind, "bind", ":8080", "The socket to bind to.")
	flagset.Parse(os.Args[1:])

	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestsTotal)
	r.MustRegister(version)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello from example application on %s.", host)))
	})
	notfound := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Error from example application on %s.", host)))
	})
	http.Handle("/hello", promhttp.InstrumentHandlerCounter(httpRequestsTotal, handler))
	http.Handle("/error", promhttp.InstrumentHandlerCounter(httpRequestsTotal, notfound))

	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(bind, nil))
}
