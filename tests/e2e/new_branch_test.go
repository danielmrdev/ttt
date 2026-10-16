package e2e

import (
	"testing"

	"github.com/eugenioenko/ttt/internal/git"
	"github.com/gdamore/tcell/v3"
)

func TestNewBranchCommandCreatesBranch(t *testing.T) {
	h := newTestHarness(t, 100, 30)
	defer h.stop()

	gitIn(t, h.dir, "init", "-q")
	gitIn(t, h.dir, "commit", "--allow-empty", "-q", "-m", "initial")
	h.app.Changes.Dirs = []string{h.dir}

	const branch = "feature/from-changes-panel"
	h.exec("git.newBranch")
	h.assertContains("Branch name")
	for _, r := range branch {
		h.pressRune(r)
	}
	h.pressKey(tcell.KeyEnter, tcell.ModNone)

	result := h.awaitRepoOp()
	if result.Err != nil {
		t.Fatalf("creating branch: %v", result.Err)
	}
	if got := git.BranchName(h.dir); got != branch {
		t.Fatalf("current branch = %q, want %q", got, branch)
	}
	h.assertContains("Created branch: " + branch)
}
