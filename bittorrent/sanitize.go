package bittorrent

// ErrInvalidIP indicates an invalid IP for an Announce.
var ErrInvalidIP = ClientError("invalid IP")

// ErrInvalidPort indicates an invalid Port for an Announce.
var ErrInvalidPort = ClientError("invalid port")

// SanitizeAnnounce enforces a max and default NumWant and coerces the peer's
// IP address into the proper format.
func SanitizeAnnounce(r *AnnounceRequest, maxNumWant, defaultNumWant uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// implies r.Peer.IP.To4() == nil

// SanitizeScrape enforces a max number of infohashes for a single scrape
// request.
func SanitizeScrape(r *ScrapeRequest, maxScrapeInfoHashes uint32) error {
	_ = "STUB: not implemented"
	return nil
}
