package http

import (
	"net/http"

	"github.com/chihaya/chihaya/bittorrent"
	"github.com/chihaya/chihaya/frontend/http/bencode"
)

// WriteError communicates an error to a BitTorrent client over HTTP.
func WriteError(w http.ResponseWriter, err error) error { _ = "STUB: not implemented"; return nil }

// WriteAnnounceResponse communicates the results of an Announce to a
// BitTorrent client over HTTP.
func WriteAnnounceResponse(w http.ResponseWriter, resp *bittorrent.AnnounceResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// Add the peers to the dictionary in the compact format.

// Add the IPv4 peers to the dictionary.

// Add the IPv6 peers to the dictionary.

// Add the peers to the dictionary.

// WriteScrapeResponse communicates the results of a Scrape to a BitTorrent
// client over HTTP.
func WriteScrapeResponse(w http.ResponseWriter, resp *bittorrent.ScrapeResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func compact4(peer bittorrent.Peer) (buf []byte) { _ = "STUB: not implemented"; return nil }

func compact6(peer bittorrent.Peer) (buf []byte) { _ = "STUB: not implemented"; return nil }

func dict(peer bittorrent.Peer) bencode.Dict { _ = "STUB: not implemented"; return *new(bencode.Dict) }
