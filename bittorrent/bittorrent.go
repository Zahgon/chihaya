// Package bittorrent implements all of the abstractions used to decouple the
// protocol of a BitTorrent tracker from the logic of handling Announces and
// Scrapes.
package bittorrent

import (
	"net"
	"time"

	"github.com/chihaya/chihaya/pkg/log"
)

// PeerID represents a peer ID.
type PeerID [20]byte

// PeerIDFromBytes creates a PeerID from a byte slice.
//
// It panics if b is not 20 bytes long.
func PeerIDFromBytes(b []byte) PeerID { _ = "STUB: not implemented"; return *new(PeerID) }

// String implements fmt.Stringer, returning the base16 encoded PeerID.
func (p PeerID) String() string { _ = "STUB: not implemented"; return "" }

// RawString returns a 20-byte string of the raw bytes of the ID.
func (p PeerID) RawString() string { _ = "STUB: not implemented"; return "" }

// PeerIDFromString creates a PeerID from a string.
//
// It panics if s is not 20 bytes long.
func PeerIDFromString(s string) PeerID { _ = "STUB: not implemented"; return *new(PeerID) }

// InfoHash represents an infohash.
type InfoHash [20]byte

// InfoHashFromBytes creates an InfoHash from a byte slice.
//
// It panics if b is not 20 bytes long.
func InfoHashFromBytes(b []byte) InfoHash { _ = "STUB: not implemented"; return *new(InfoHash) }

// InfoHashFromString creates an InfoHash from a string.
//
// It panics if s is not 20 bytes long.
func InfoHashFromString(s string) InfoHash { _ = "STUB: not implemented"; return *new(InfoHash) }

// String implements fmt.Stringer, returning the base16 encoded InfoHash.
func (i InfoHash) String() string { _ = "STUB: not implemented"; return "" }

// RawString returns a 20-byte string of the raw bytes of the InfoHash.
func (i InfoHash) RawString() string { _ = "STUB: not implemented"; return "" }

// AnnounceRequest represents the parsed parameters from an announce request.
type AnnounceRequest struct {
	Event           Event
	InfoHash        InfoHash
	Compact         bool
	EventProvided   bool
	NumWantProvided bool
	IPProvided      bool
	NumWant         uint32
	Left            uint64
	Downloaded      uint64
	Uploaded        uint64

	Peer
	Params
}

// LogFields renders the current response as a set of log fields.
func (r AnnounceRequest) LogFields() log.Fields { _ = "STUB: not implemented"; return *new(log.Fields) }

// AnnounceResponse represents the parameters used to create an announce
// response.
type AnnounceResponse struct {
	Compact     bool
	Complete    uint32
	Incomplete  uint32
	Interval    time.Duration
	MinInterval time.Duration
	IPv4Peers   []Peer
	IPv6Peers   []Peer
}

// LogFields renders the current response as a set of log fields.
func (r AnnounceResponse) LogFields() log.Fields {
	_ = "STUB: not implemented"
	return *new(log.Fields)
}

// ScrapeRequest represents the parsed parameters from a scrape request.
type ScrapeRequest struct {
	AddressFamily AddressFamily
	InfoHashes    []InfoHash
	Params        Params
}

// LogFields renders the current response as a set of log fields.
func (r ScrapeRequest) LogFields() log.Fields { _ = "STUB: not implemented"; return *new(log.Fields) }

// ScrapeResponse represents the parameters used to create a scrape response.
//
// The Scrapes must be in the same order as the InfoHashes in the corresponding
// ScrapeRequest.
type ScrapeResponse struct {
	Files []Scrape
}

// LogFields renders the current response as a set of Logrus fields.
func (sr ScrapeResponse) LogFields() log.Fields { _ = "STUB: not implemented"; return *new(log.Fields) }

// Scrape represents the state of a swarm that is returned in a scrape response.
type Scrape struct {
	InfoHash   InfoHash
	Snatches   uint32
	Complete   uint32
	Incomplete uint32
}

// AddressFamily is the address family of an IP address.
type AddressFamily uint8

func (af AddressFamily) String() string { _ = "STUB: not implemented"; return "" }

// AddressFamily constants.
const (
	IPv4 AddressFamily = iota
	IPv6
)

// IP is a net.IP with an AddressFamily.
type IP struct {
	net.IP
	AddressFamily
}

func (ip IP) String() string { _ = "STUB: not implemented"; return "" }

// Peer represents the connection details of a peer that is returned in an
// announce response.
type Peer struct {
	ID   PeerID
	IP   IP
	Port uint16
}

// String implements fmt.Stringer to return a human-readable representation.
// The string will have the format <PeerID>@[<IP>]:<port>, for example
// "0102030405060708090a0b0c0d0e0f1011121314@[10.11.12.13]:1234"
func (p Peer) String() string { _ = "STUB: not implemented"; return "" }

// LogFields renders the current peer as a set of Logrus fields.
func (p Peer) LogFields() log.Fields { _ = "STUB: not implemented"; return *new(log.Fields) }

// Equal reports whether p and x are the same.
func (p Peer) Equal(x Peer) bool { _ = "STUB: not implemented"; return false }

// EqualEndpoint reports whether p and x have the same endpoint.
func (p Peer) EqualEndpoint(x Peer) bool { _ = "STUB: not implemented"; return false }

// ClientError represents an error that should be exposed to the client over
// the BitTorrent protocol implementation.
type ClientError string

// Error implements the error interface for ClientError.
func (c ClientError) Error() string { _ = "STUB: not implemented"; return "" }
