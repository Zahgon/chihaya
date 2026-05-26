package storage

import (
	"testing"

	"github.com/chihaya/chihaya/bittorrent"
)

type benchData struct {
	infohashes [1000]bittorrent.InfoHash
	peers      [1000]bittorrent.Peer
}

func generateInfohashes() (a [1000]bittorrent.InfoHash) { _ = "STUB: not implemented"; return nil }

func generatePeers() (a [1000]bittorrent.Peer) { _ = "STUB: not implemented"; return nil }

type (
	executionFunc func(int, PeerStore, *benchData) error
	setupFunc     func(PeerStore, *benchData) error
)

func runBenchmark(b *testing.B, ps PeerStore, parallel bool, sf setupFunc, ef executionFunc) {
	_ = "STUB: not implemented"
	return
}

// Nop executes a no-op for each iteration.
// It should produce the same results for each PeerStore.
// This can be used to get an estimate of the impact of the benchmark harness
// on benchmark results and an estimate of the general performance of the system
// benchmarked on.
//
// Nop can run in parallel.
func Nop(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// Put benchmarks the PutSeeder method of a PeerStore by repeatedly Putting the
// same Peer for the same InfoHash.
//
// Put can run in parallel.
func Put(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// Put1k benchmarks the PutSeeder method of a PeerStore by cycling through 1000
// Peers and Putting them into the swarm of one infohash.
//
// Put1k can run in parallel.
func Put1k(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// Put1kInfohash benchmarks the PutSeeder method of a PeerStore by cycling
// through 1000 infohashes and putting the same peer into their swarms.
//
// Put1kInfohash can run in parallel.
func Put1kInfohash(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// Put1kInfohash1k benchmarks the PutSeeder method of a PeerStore by cycling
// through 1000 infohashes and 1000 Peers and calling Put with them.
//
// Put1kInfohash1k can run in parallel.
func Put1kInfohash1k(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// PutDelete benchmarks the PutSeeder and DeleteSeeder methods of a PeerStore by
// calling PutSeeder followed by DeleteSeeder for one Peer and one infohash.
//
// PutDelete can not run in parallel.
func PutDelete(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// PutDelete1k benchmarks the PutSeeder and DeleteSeeder methods in the same way
// PutDelete does, but with one from 1000 Peers per iteration.
//
// PutDelete1k can not run in parallel.
func PutDelete1k(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// PutDelete1kInfohash behaves like PutDelete1k with 1000 infohashes instead of
// 1000 Peers.
//
// PutDelete1kInfohash can not run in parallel.
func PutDelete1kInfohash(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// PutDelete1kInfohash1k behaves like PutDelete1k with 1000 infohashes in
// addition to 1000 Peers.
//
// PutDelete1kInfohash1k can not run in parallel.
func PutDelete1kInfohash1k(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// DeleteNonexist benchmarks the DeleteSeeder method of a PeerStore by
// attempting to delete a Peer that is nonexistent.
//
// DeleteNonexist can run in parallel.
func DeleteNonexist(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// DeleteNonexist1k benchmarks the DeleteSeeder method of a PeerStore by
// attempting to delete one of 1000 nonexistent Peers.
//
// DeleteNonexist can run in parallel.
func DeleteNonexist1k(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// DeleteNonexist1kInfohash benchmarks the DeleteSeeder method of a PeerStore by
// attempting to delete one Peer from one of 1000 infohashes.
//
// DeleteNonexist1kInfohash can run in parallel.
func DeleteNonexist1kInfohash(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// DeleteNonexist1kInfohash1k benchmarks the Delete method of a PeerStore by
// attempting to delete one of 1000 Peers from one of 1000 Infohashes.
//
// DeleteNonexist1kInfohash1k can run in parallel.
func DeleteNonexist1kInfohash1k(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// GradNonexist benchmarks the GraduateLeecher method of a PeerStore by
// attempting to graduate a nonexistent Peer.
//
// GradNonexist can run in parallel.
func GradNonexist(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// GradNonexist1k benchmarks the GraduateLeecher method of a PeerStore by
// attempting to graduate one of 1000 nonexistent Peers.
//
// GradNonexist1k can run in parallel.
func GradNonexist1k(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// GradNonexist1kInfohash benchmarks the GraduateLeecher method of a PeerStore
// by attempting to graduate a nonexistent Peer for one of 100 Infohashes.
//
// GradNonexist1kInfohash can run in parallel.
func GradNonexist1kInfohash(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// GradNonexist1kInfohash1k benchmarks the GraduateLeecher method of a PeerStore
// by attempting to graduate one of 1000 nonexistent Peers for one of 1000
// infohashes.
//
// GradNonexist1kInfohash1k can run in parallel.
func GradNonexist1kInfohash1k(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// PutGradDelete benchmarks the PutLeecher, GraduateLeecher and DeleteSeeder
// methods of a PeerStore by adding one leecher to a swarm, promoting it to a
// seeder and deleting the seeder.
//
// PutGradDelete can not run in parallel.
func PutGradDelete(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// PutGradDelete1k behaves like PutGradDelete with one of 1000 Peers.
//
// PutGradDelete1k can not run in parallel.
func PutGradDelete1k(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// PutGradDelete1kInfohash behaves like PutGradDelete with one of 1000
// infohashes.
//
// PutGradDelete1kInfohash can not run in parallel.
func PutGradDelete1kInfohash(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// PutGradDelete1kInfohash1k behaves like PutGradDelete with one of 1000 Peers
// and one of 1000 infohashes.
//
// PutGradDelete1kInfohash can not run in parallel.
func PutGradDelete1kInfohash1k(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

func putPeers(ps PeerStore, bd *benchData) error { _ = "STUB: not implemented"; return nil }

// AnnounceLeecher benchmarks the AnnouncePeers method of a PeerStore for
// announcing a leecher.
// The swarm announced to has 500 seeders and 500 leechers.
//
// AnnounceLeecher can run in parallel.
func AnnounceLeecher(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// AnnounceLeecher1kInfohash behaves like AnnounceLeecher with one of 1000
// infohashes.
//
// AnnounceLeecher1kInfohash can run in parallel.
func AnnounceLeecher1kInfohash(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// AnnounceSeeder behaves like AnnounceLeecher with a seeder instead of a
// leecher.
//
// AnnounceSeeder can run in parallel.
func AnnounceSeeder(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// AnnounceSeeder1kInfohash behaves like AnnounceSeeder with one of 1000
// infohashes.
//
// AnnounceSeeder1kInfohash can run in parallel.
func AnnounceSeeder1kInfohash(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// ScrapeSwarm benchmarks the ScrapeSwarm method of a PeerStore.
// The swarm scraped has 500 seeders and 500 leechers.
//
// ScrapeSwarm can run in parallel.
func ScrapeSwarm(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }

// ScrapeSwarm1kInfohash behaves like ScrapeSwarm with one of 1000 infohashes.
//
// ScrapeSwarm1kInfohash can run in parallel.
func ScrapeSwarm1kInfohash(b *testing.B, ps PeerStore) { _ = "STUB: not implemented"; return }
