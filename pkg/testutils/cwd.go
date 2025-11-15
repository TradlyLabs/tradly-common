package testutils

import (
	"context"
	"os"
	"testing"

	"github.com/TradlyLabs/tradly-common/pkg/runtime"
)

func ChdirAndRun(t *testing.T, path string) {
	os.Chdir(path)
	t.Cleanup(func() {
		os.Chdir(t.TempDir())
	})
	t.Logf("Changed working directory to %s", path)
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	if err := runtime.DefaultManager.Start(ctx); err != nil {
		t.Fatalf("Failed to start runtime: %v", err)
	}
	t.Cleanup(func() {
		if err := runtime.DefaultManager.Stop(ctx); err != nil {
			t.Fatalf("Failed to stop runtime: %v", err)
		}
	})
}
