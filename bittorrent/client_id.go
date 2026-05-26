package bittorrent

// ClientID represents the part of a PeerID that identifies a Peer's client
// software.
type ClientID [6]byte

// NewClientID parses a ClientID from a PeerID.
func NewClientID(pid PeerID) ClientID { _ = "STUB: not implemented"; return *new(ClientID) }
