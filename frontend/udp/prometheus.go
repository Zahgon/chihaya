package udp

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/chihaya/chihaya/bittorrent"
)

func init() {
	prometheus.MustRegister(promResponseDurationMilliseconds)
}

var promResponseDurationMilliseconds = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "chihaya_udp_response_duration_milliseconds",
		Help:    "The duration of time it takes to receive and write a response to an API request",
		Buckets: prometheus.ExponentialBuckets(9.375, 2, 10),
	},
	[]string{"action", "address_family", "error"},
)

// recordResponseDuration records the duration of time to respond to a UDP
// Request in milliseconds.
func recordResponseDuration(action string, af *bittorrent.AddressFamily, err error, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}
