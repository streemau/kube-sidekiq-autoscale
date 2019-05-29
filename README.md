# kube-sidekiq-autoscale

Thanks to [Maciej Bogus](https://github.com/mbogus) for building [kube-amqp-autoscale](https://github.com/mbogus/kube-amqp-autoscale). This repository is a clone of his code and replaces the AMQP related code with code to monitor Sidekiq instead.

Dynamically scale K8s resources using the amount of enqueued jobs to determine the load on an application/Kubernetes pod.

**NOTICE**

If your application load is not queue-bound but rather CPU-sensitive, make sure to use built-in Kubernetes [Horizontal Pod Autoscaling](http://kubernetes.io/docs/user-guide/horizontal-pod-autoscaling/) instead of this project.

## Status

*Alpha*

## Usage



## Go get

    go get github.com/streemau/kube-sidekiq-autoscale


## Clone from [github](https://github.com/streemau/kube-sidekiq-autoscale)

* Create directory for APT projects `mkdir -p $GOPATH/src/github.com/streemau`
  as typical in [writing go programs](https://golang.org/doc/code.html)
* Clone this project `git clone https://github.com/streemau/kube-sidekiq-autoscale.git`


### Building on Windows
If you have a unix-y shell on Windows ([MSYS2](http://sourceforge.net/p/msys2/wiki/MSYS2%20installation/),
[CYGWIN](https://cygwin.com/install.html) or other), see *Build project* below.


### Build project

The project depends on several external Go projects that can be automatically
downloaded using `make depend` target.

Run `make depend && make [build]`


## Runtime environment variables

* `GOMAXPROCS` limits the number of operating system threads that can execute
  user-level Go code simultaneously


## Runtime command-line arguments

* **`sidekiq-stats-uri`** required, Sidekiq Stats URI, e.g. `http://username:passwd@sidekiq-host/sidekiq/stats`
* **`api-url`** required, Kubernetes API URL, e.g. `http://127.0.0.1:8080`
* `api-user` optional, username for basic authentication on Kubernetes API
* `api-passwd` optional, password for basic authentication on Kubernetes API
* `api-token` optional, path to a bearer token file for OAuth authentication, on a Kubernetes pod usually `/var/run/secrets/kubernetes.io/serviceaccount/token`
* `api-cafile` optional, path to CA certificate file for HTTPS connections to Kubernetes API from within a cluster, typically `/var/run/secrets/kubernetes.io/serviceaccount/ca.crt`
* `api-insecure` optional, set to `true` for connecting to Kubernetes API without verifying TLS certificate; unsafe, use for development only (default `false`)
* `min` lower limit for the number of replicas for a Kubernetes pod that can be set by the autoscaler (default `1`)
* **`max`** required, upper limit for the number of replicate for a Kubernetes pod that can be set by the autoscaler (must be greater than `min`)
* **`name`** required, name of the Kubernetes resource to autoscale
* `kind` type of the Kubernetes resource to autoscale, one of `Deployment`, `ReplicationController`, `ReplicaSet` (default `Deployment`)
* `ns` Kubernetes namespace (default `default`)
* `interval` time interval between Kubernetes resource scale runs in secs (default `30`)
* **`threshold`** required, number of messages on a queue representing maximum load on the autoscaled Kubernetes resource
* `increase-limit` limit number of Kubernetes pods to be provisioned in a single scale iteration to max of the value, set to a number greater than 0, default `unbounded`
* `decrease-limit` limit number of Kubernetes pods to be terminated in a single scale iteration to max of the value, set to a number greater than 0, default `unbounded`
* `stats-interval` time interval between metrics gathering runs in seconds (default `5`)
* `eval-intervals` number of autoscale intervals used to calculate average queue length (default `2`)
* `stats-coverage` required percentage of statistics to calculate average queue length (default `0.75`)
* `db` sqlite3 database filename for storing  queue length statistics (default `file::memory:?cache=shared`)
* `db-dir` directory for sqlite3 statistics database file
* `version` show version
* `metrics-listen-address` the address to listen on for exporting Prometheus metrics (default `:9505`)


## Integration tests

To run integration tests, make sure to configure access to running Sidekiq instance,
export environment variable `SIDEKIQ_STATS_URI=http://username:passwd@sidekiq-host/sidekiq/stats`
and run `go test -tags=integration ./...`
