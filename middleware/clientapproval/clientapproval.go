// Package clientapproval implements a Hook that fails an Announce based on a
// whitelist or blacklist of BitTorrent client IDs.
package clientapproval

import (
	"context"

	"github.com/chihaya/chihaya/bittorrent"
	"github.com/chihaya/chihaya/middleware"
)

// Name is the name by which this middleware is registered with Chihaya.
const Name = "client approval"

func init() {
	middleware.RegisterDriver(Name, driver{})
}

var _ middleware.Driver = driver{}

type driver struct{}

func (d driver) NewHook(optionBytes []byte) (middleware.Hook, error) {
	_ = "STUB: not implemented"
	return *new(middleware.Hook), nil
}

// ErrClientUnapproved is the error returned when a client's PeerID is invalid.
var ErrClientUnapproved = bittorrent.ClientError("unapproved client")

// Config represents all the values required by this middleware to validate
// peers based on their BitTorrent client ID.
type Config struct {
	Whitelist []string `yaml:"whitelist"`
	Blacklist []string `yaml:"blacklist"`
}

type hook struct {
	approved   map[bittorrent.ClientID]struct{}
	unapproved map[bittorrent.ClientID]struct{}
}

// NewHook returns an instance of the client approval middleware.
func NewHook(cfg Config) (middleware.Hook, error) {
	_ = "STUB: not implemented"
	return *new(middleware.Hook), nil
}

func (h *hook) HandleAnnounce(ctx context.Context, req *bittorrent.AnnounceRequest, resp *bittorrent.AnnounceResponse) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *hook) HandleScrape(ctx context.Context, req *bittorrent.ScrapeRequest, resp *bittorrent.ScrapeResponse) (context.Context, error) {
	_ = "STUB: not implemented"
	// Scrapes don't require any protection.
	return *new(context.Context), nil
}
