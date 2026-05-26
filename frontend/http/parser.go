package http

import (
	"net"
	"net/http"

	"github.com/chihaya/chihaya/bittorrent"
)

// ParseOptions is the configuration used to parse an Announce Request.
//
// If AllowIPSpoofing is true, IPs provided via BitTorrent params will be used.
// If RealIPHeader is not empty string, the value of the first HTTP Header with
// that name will be used.
type ParseOptions struct {
	AllowIPSpoofing     bool   `yaml:"allow_ip_spoofing"`
	RealIPHeader        string `yaml:"real_ip_header"`
	MaxNumWant          uint32 `yaml:"max_numwant"`
	DefaultNumWant      uint32 `yaml:"default_numwant"`
	MaxScrapeInfoHashes uint32 `yaml:"max_scrape_infohashes"`
}

// Default parser config constants.
const (
	defaultMaxNumWant          = 100
	defaultDefaultNumWant      = 50
	defaultMaxScrapeInfoHashes = 50
)

// ParseAnnounce parses an bittorrent.AnnounceRequest from an http.Request.
func ParseAnnounce(r *http.Request, opts ParseOptions) (*bittorrent.AnnounceRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attempt to parse the event from the request.

// Determine if the client expects a compact response.

// Parse the infohash from the request.

// Parse the PeerID from the request.

// Determine the number of remaining bytes for the client.

// Determine the number of bytes downloaded by the client.

// Determine the number of bytes shared by the client.

// Determine the number of peers the client wants in the response.

// If there were no errors, the user actually provided the numwant.

// Parse the port where the client is listening.

// Parse the IP address where the client is listening.

// ParseScrape parses an bittorrent.ScrapeRequest from an http.Request.
func ParseScrape(r *http.Request, opts ParseOptions) (*bittorrent.ScrapeRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// requestedIP determines the IP address for a BitTorrent client request.
func requestedIP(r *http.Request, p bittorrent.Params, opts ParseOptions) (ip net.IP, provided bool) {
	_ = "STUB: not implemented"
	return *new(net.IP), false
}
