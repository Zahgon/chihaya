// Package torrentapproval implements a Hook that fails an Announce based on a
// whitelist or blacklist of torrent hash.
package torrentapproval

import (
	"context"

	"github.com/chihaya/chihaya/bittorrent"
	"github.com/chihaya/chihaya/middleware"
)

// Name is the name by which this middleware is registered with Chihaya.
const Name = "torrent approval"

func init() {
	middleware.RegisterDriver(Name, driver{})
}

var _ middleware.Driver = driver{}

type driver struct{}

func (d driver) NewHook(optionBytes []byte) (middleware.Hook, error) {
	_ = "STUB: not implemented"
	return *new(middleware.Hook), nil
}

// ErrTorrentUnapproved is the error returned when a torrent hash is invalid.
var ErrTorrentUnapproved = bittorrent.ClientError("unapproved torrent")

// Config represents all the values required by this middleware to validate
// torrents based on their hash value.
type Config struct {
	Whitelist []string `yaml:"whitelist"`
	Blacklist []string `yaml:"blacklist"`
}

type hook struct {
	approved   map[bittorrent.InfoHash]struct{}
	unapproved map[bittorrent.InfoHash]struct{}
}

// NewHook returns an instance of the torrent approval middleware.
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
