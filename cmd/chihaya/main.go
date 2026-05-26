// Package main provides argument parsing and execution logic for the command line tool.
package main

import (
	"runtime"
	"time"

	"github.com/spf13/cobra"

	"github.com/chihaya/chihaya/middleware"
	"github.com/chihaya/chihaya/pkg/log"
	"github.com/chihaya/chihaya/pkg/stop"
	"github.com/chihaya/chihaya/storage"
)

// Run represents the state of a running instance of Chihaya.
type Run struct {
	configFilePath string
	peerStore      storage.PeerStore
	logic          *middleware.Logic
	sg             *stop.Group
}

// NewRun runs an instance of Chihaya.
func NewRun(configFilePath string) (*Run, error) { _ = "STUB: not implemented"; return nil, nil }

// Start begins an instance of Chihaya.
// It is optional to provide an instance of the peer store to avoid the
// creation of a new one.
func (r *Run) Start(ps storage.PeerStore) error { _ = "STUB: not implemented"; return nil }

func combineErrors(prefix string, errs []error) error { _ = "STUB: not implemented"; return nil }

// Stop shuts down an instance of Chihaya.
func (r *Run) Stop(keepPeerStore bool) (storage.PeerStore, error) {
	_ = "STUB: not implemented"
	return *new(storage.PeerStore), nil
}

// RootRunCmdFunc implements a Cobra command that runs an instance of Chihaya
// and handles reloading and shutdown via process signals.
func RootRunCmdFunc(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// RootPreRunCmdFunc handles command line flags for the Run command.
func RootPreRunCmdFunc(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// RootPostRunCmdFunc handles clean up of any state initialized by command line
// flags.
func RootPostRunCmdFunc(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	rootCmd := &cobra.Command{
		Use:                "chihaya",
		Short:              "BitTorrent Tracker",
		Long:               "A customizable, multi-protocol BitTorrent Tracker",
		PersistentPreRunE:  RootPreRunCmdFunc,
		RunE:               RootRunCmdFunc,
		PersistentPostRunE: RootPostRunCmdFunc,
	}

	rootCmd.PersistentFlags().Bool("debug", false, "enable debug logging")
	rootCmd.PersistentFlags().Bool("json", false, "enable json logging")
	if runtime.GOOS == "windows" {
		rootCmd.PersistentFlags().Bool("nocolors", true, "disable log coloring")
	} else {
		rootCmd.PersistentFlags().Bool("nocolors", false, "disable log coloring")
	}

	rootCmd.Flags().String("config", "/etc/chihaya.yaml", "location of configuration file")

	e2eCmd := &cobra.Command{
		Use:   "e2e",
		Short: "exec e2e tests",
		Long:  "Execute the Chihaya end-to-end test suite",
		RunE:  EndToEndRunCmdFunc,
	}

	e2eCmd.Flags().String("httpaddr", "http://127.0.0.1:6969/announce", "address of the HTTP tracker")
	e2eCmd.Flags().String("udpaddr", "udp://127.0.0.1:6969", "address of the UDP tracker")
	e2eCmd.Flags().Duration("delay", time.Second, "delay between announces")

	rootCmd.AddCommand(e2eCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("failed when executing root cobra command: " + err.Error())
	}
}
