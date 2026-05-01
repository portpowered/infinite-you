package smoke_test

import (
	"os"
	"testing"

	"github.com/portpowered/agent-factory/pkg/interfaces"
)

func setWorkingDirectory(t *testing.T, dir string) {
	t.Helper()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd(): %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("Chdir(%q): %v", dir, err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Fatalf("restore working directory: %v", err)
		}
	})
}

func assertDispatchHistoryContainsWorkstation(
	t *testing.T,
	history []interfaces.CompletedDispatch,
	workstationName string,
	terminalPlace string,
	workID string,
) {
	t.Helper()

	for _, dispatch := range history {
		if dispatch.WorkstationName != workstationName {
			continue
		}
		for _, mutation := range dispatch.OutputMutations {
			if mutation.ToPlace != terminalPlace || mutation.Token == nil {
				continue
			}
			if mutation.Token.Color.WorkID == workID {
				return
			}
		}
	}

	t.Fatalf(
		"dispatch history missing %q route to %q for work %q: %#v",
		workstationName,
		terminalPlace,
		workID,
		history,
	)
}
