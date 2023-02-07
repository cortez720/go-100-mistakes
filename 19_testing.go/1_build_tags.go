//go:build integration
//go:build !integration runs only if integration flag is not enabled.

package tests

import "testing"

// go test -v // runs tests without tags
// go test -v --tags=integration runs tests without tags and with tag integration

// go test -short
func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long tests")
	}
}
