package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const defaultPackage = "./tests/functional/default/..."

func main() {
	budget := flag.Duration("budget", 10*time.Second, "maximum allowed runtime for the default functional lane")
	goBinary := flag.String("go-binary", os.Getenv("GO"), "go toolchain binary/path used to run the default functional lane")
	timeout := flag.String("timeout", "300s", "go test timeout passed to the default functional lane")
	flag.Parse()

	elapsed, err := runDefaultFunctionalLane(resolveGoBinary(*goBinary), *timeout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if elapsed > *budget {
		fmt.Fprintf(os.Stderr, "[agent-factory:functional-runtime] default lane exceeded budget: %s > %s\n", elapsed.Round(time.Millisecond), budget.Round(time.Millisecond))
		os.Exit(1)
	}

	fmt.Printf("[agent-factory:functional-runtime] default lane completed in %s (budget %s)\n", elapsed.Round(time.Millisecond), budget.Round(time.Millisecond))
}

func resolveGoBinary(configured string) string {
	if configured == "" {
		return "go"
	}
	return configured
}

func runDefaultFunctionalLane(goBinary string, timeout string) (time.Duration, error) {
	cmd := exec.Command(goBinary, "test", defaultPackage, "-count=1", "-timeout", timeout)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	start := time.Now()
	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("run default functional lane: %w", err)
	}
	return time.Since(start), nil
}
