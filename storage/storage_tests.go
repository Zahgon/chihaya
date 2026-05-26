package storage

import (
	"testing"

	"github.com/chihaya/chihaya/bittorrent"
)

// PeerEqualityFunc is the boolean function to use to check two Peers for
// equality.
// Depending on the implementation of the PeerStore, this can be changed to
// use (Peer).EqualEndpoint instead.
var PeerEqualityFunc = func(p1, p2 bittorrent.Peer) bool { return p1.Equal(p2) }

// TestPeerStore tests a PeerStore implementation against the interface.
func TestPeerStore(t *testing.T, p PeerStore) { _ = "STUB: not implemented"; return }

// Test ErrDNE for non-existent swarms.

// Test empty scrape response for non-existent swarms.

// Insert dummy Peer to keep swarm active
// Has the same address family as c.peer

// Test ErrDNE for non-existent seeder.

// Test PutLeecher -> Announce -> DeleteLeecher -> Announce

// non-seeder announce should still return the leecher

// Test PutSeeder -> Announce -> DeleteSeeder -> Announce

// Should be leecher to see the seeder

// Test PutLeecher -> Graduate -> Announce -> DeleteLeecher -> Announce

// Has to be leecher to see the graduated seeder

// Deleting the Peer as a Leecher should have no effect

// Verify it's still there

// Clean up

// Test ErrDNE for missing leecher

func containsPeer(peers []bittorrent.Peer, p bittorrent.Peer) bool {
	_ = "STUB: not implemented"
	return false
}
