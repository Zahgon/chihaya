package udp

import (
	"io"

	"github.com/chihaya/chihaya/bittorrent"
)

// WriteError writes the failure reason as a null-terminated string.
func WriteError(w io.Writer, txID []byte, err error) {
	_ = "STUB: not implemented"
	// If the client wasn't at fault, acknowledge it.
	return
}

// WriteAnnounce encodes an announce response according to BEP 15.
// The peers returned will be resp.IPv6Peers or resp.IPv4Peers, depending on
// whether v6Peers is set.
// If v6Action is set, the action will be 4, according to
// https://web.archive.org/web/20170503181830/http://opentracker.blog.h3q.com/2007/12/28/the-ipv6-situation/
func WriteAnnounce(w io.Writer, txID []byte, resp *bittorrent.AnnounceResponse, v6Action, v6Peers bool) {
	_ = "STUB: not implemented"
	return
}

// WriteScrape encodes a scrape response according to BEP 15.
func WriteScrape(w io.Writer, txID []byte, resp *bittorrent.ScrapeResponse) {
	_ = "STUB: not implemented"
	return
}

// WriteConnectionID encodes a new connection response according to BEP 15.
func WriteConnectionID(w io.Writer, txID, connID []byte) { _ = "STUB: not implemented"; return }

// writeHeader writes the action and transaction ID to the provided response
// buffer.
func writeHeader(w io.Writer, txID []byte, action uint32) { _ = "STUB: not implemented"; return }
