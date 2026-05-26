package redis

import (
	"time"

	"github.com/go-redsync/redsync/v4"
	redigolib "github.com/gomodule/redigo/redis"
)

// redisBackend represents a redis handler.
type redisBackend struct {
	pool    *redigolib.Pool
	redsync *redsync.Redsync
}

// newRedisBackend creates a redisBackend instance.
func newRedisBackend(cfg *Config, u *redisURL, socketPath string) *redisBackend {
	_ = "STUB: not implemented"
	return nil
}

// open returns or creates instance of Redis connection.
func (rb *redisBackend) open() redigolib.Conn {
	_ = "STUB: not implemented"
	return *new(redigolib.Conn)
}

type redisConnector struct {
	URL            *redisURL
	SocketPath     string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	ConnectTimeout time.Duration
}

// NewPool returns a new pool of Redis connections
func (rc *redisConnector) NewPool() *redigolib.Pool { _ = "STUB: not implemented"; return nil }

// PINGs connections that have been idle more than 10 seconds

// Open a new Redis connection
func (rc *redisConnector) open() (redigolib.Conn, error) {
	_ = "STUB: not implemented"
	return *new(redigolib.Conn), nil
}

// A redisURL represents a parsed redisURL
// The general form represented is:
//
//	redis://[password@]host][/][db]
type redisURL struct {
	Host     string
	Password string
	DB       int
}

// parseRedisURL parse rawurl into redisURL
func parseRedisURL(target string) (*redisURL, error) { _ = "STUB: not implemented"; return nil, nil }

// default redis db
