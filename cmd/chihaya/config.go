package main

import (
	"github.com/chihaya/chihaya/frontend/http"
	"github.com/chihaya/chihaya/frontend/udp"
	"github.com/chihaya/chihaya/middleware"

	// Imports to register middleware drivers.
	_ "github.com/chihaya/chihaya/middleware/clientapproval"
	_ "github.com/chihaya/chihaya/middleware/jwt"
	_ "github.com/chihaya/chihaya/middleware/torrentapproval"
	_ "github.com/chihaya/chihaya/middleware/varinterval"

	// Imports to register storage drivers.
	_ "github.com/chihaya/chihaya/storage/memory"
	_ "github.com/chihaya/chihaya/storage/redis"
)

type storageConfig struct {
	Name   string      `yaml:"name"`
	Config interface{} `yaml:"config"`
}

// Config represents the configuration used for executing Chihaya.
type Config struct {
	middleware.ResponseConfig `yaml:",inline"`
	MetricsAddr               string                  `yaml:"metrics_addr"`
	HTTPConfig                http.Config             `yaml:"http"`
	UDPConfig                 udp.Config              `yaml:"udp"`
	Storage                   storageConfig           `yaml:"storage"`
	PreHooks                  []middleware.HookConfig `yaml:"prehooks"`
	PostHooks                 []middleware.HookConfig `yaml:"posthooks"`
}

// PreHookNames returns only the names of the configured middleware.
func (cfg Config) PreHookNames() (names []string) { _ = "STUB: not implemented"; return nil }

// PostHookNames returns only the names of the configured middleware.
func (cfg Config) PostHookNames() (names []string) { _ = "STUB: not implemented"; return nil }

// ConfigFile represents a namespaced YAML configation file.
type ConfigFile struct {
	Chihaya Config `yaml:"chihaya"`
}

// ParseConfigFile returns a new ConfigFile given the path to a YAML
// configuration file.
//
// It supports relative and absolute paths and environment variables.
func ParseConfigFile(path string) (*ConfigFile, error) { _ = "STUB: not implemented"; return nil, nil }
