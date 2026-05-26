// Package varinterval provides a middleware implementation that handles
// announcement intervals according to configured ranges,
// giving different values for each peer so they don't re-announce at the same time.
package varinterval

import (
	"context"
	"errors"
	"sync"

	"github.com/chihaya/chihaya/bittorrent"
	"github.com/chihaya/chihaya/middleware"
)

// Name is the name by which this middleware is registered with Chihaya.
const Name = "interval variation"

func init() {
	middleware.RegisterDriver(Name, driver{})
}

var _ middleware.Driver = driver{}

type driver struct{}

func (d driver) NewHook(optionBytes []byte) (middleware.Hook, error) {
	_ = "STUB: not implemented"
	return *new(middleware.Hook), nil
}

// ErrInvalidModifyResponseProbability is returned for a config with an invalid
// ModifyResponseProbability.
var ErrInvalidModifyResponseProbability = errors.New("invalid modify_response_probability")

// ErrInvalidMaxIncreaseDelta is returned for a config with an invalid
// MaxIncreaseDelta.
var ErrInvalidMaxIncreaseDelta = errors.New("invalid max_increase_delta")

// Config represents the configuration for the varinterval middleware.
type Config struct {
	// ModifyResponseProbability is the probability by which a response will
	// be modified.
	ModifyResponseProbability float32 `yaml:"modify_response_probability"`

	// MaxIncreaseDelta is the amount of seconds that will be added at most.
	MaxIncreaseDelta int `yaml:"max_increase_delta"`

	// ModifyMinInterval specifies whether min_interval should be increased
	// as well.
	ModifyMinInterval bool `yaml:"modify_min_interval"`
}

func checkConfig(cfg Config) error { _ = "STUB: not implemented"; return nil }

type hook struct {
	cfg Config
	sync.Mutex
}

// NewHook creates a middleware to randomly modify the announce interval from
// the given config.
func NewHook(cfg Config) (middleware.Hook, error) {
	_ = "STUB: not implemented"
	return *new(middleware.Hook), nil
}

func (h *hook) HandleAnnounce(ctx context.Context, req *bittorrent.AnnounceRequest, resp *bittorrent.AnnounceResponse) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Generate a probability p < 1.0.

// Generate the increase delta.

func (h *hook) HandleScrape(ctx context.Context, req *bittorrent.ScrapeRequest, resp *bittorrent.ScrapeResponse) (context.Context, error) {
	_ = "STUB: not implemented"
	// Scrapes are not altered.
	return *new(context.Context), nil
}
