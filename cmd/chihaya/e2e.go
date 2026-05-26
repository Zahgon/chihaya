package main

import (
	"time"

	"github.com/spf13/cobra"
)

// EndToEndRunCmdFunc implements a Cobra command that runs the end-to-end test
// suite for a Chihaya build.
func EndToEndRunCmdFunc(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Test the HTTP tracker

// Test the UDP tracker.

func generateInfohash() [20]byte { _ = "STUB: not implemented"; return nil }

func test(addr string, delay time.Duration) error { _ = "STUB: not implemented"; return nil }

func testWithInfohash(infoHash [20]byte, url string, delay time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
