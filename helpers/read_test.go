package helpers

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInputReader_Lines_ReturnsAllLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")
	content := "first line\nsecond line\nthird line\n"
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("failed writing temp file: %v", err)
	}

	ir := NewInputReader(path)
	got := ir.Lines()
	want := []string{"first line", "second line", "third line"}
	if len(got) != len(want) {
		t.Fatalf("Lines() length = %d, want %d; got=%v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Lines()[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestInputReader_Lines_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.txt")
	if err := os.WriteFile(path, []byte(""), 0o600); err != nil {
		t.Fatalf("failed writing temp file: %v", err)
	}

	ir := NewInputReader(path)
	got := ir.Lines()
	if len(got) != 0 {
		t.Fatalf("Lines() length = %d, want 0; got=%v", len(got), got)
	}
}
