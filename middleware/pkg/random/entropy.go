package random

import (
	"github.com/chihaya/chihaya/bittorrent"
)

// DeriveEntropyFromRequest generates 2*64 bits of pseudo random state from an
// AnnounceRequest.
//
// Calling DeriveEntropyFromRequest multiple times yields the same values.
func DeriveEntropyFromRequest(req *bittorrent.AnnounceRequest) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}
