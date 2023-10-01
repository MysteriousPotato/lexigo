package main_test

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"os/exec"
	"testing"
)

//go:generate go build -o lexigogen ./
//go:generate ./lexigogen -p ./test -o ./test/test_tmp.gen.go -pkg test
func TestGoGenerate(t *testing.T) {
	genFilename := "./test/test_tmp.gen.go"
	defer func() {
		if err := os.Remove(genFilename); err != nil {
			t.Logf("failed to remove generated file: %v", err)
		}

		if err := os.Remove("lexigogen"); err != nil {
			t.Logf("failed to remove binary: %v", err)
		}
	}()

	cmd := exec.Command("go", "generate")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("cmd failed: %v", err)
	}

	gen, err := os.ReadFile(genFilename)
	if err != nil {
		t.Fatal(err)
	}

	golden, err := os.ReadFile("./test/test.gen.go")
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(gen, golden); diff != "" {
		t.Fatalf("generated file does not match golden: %v", diff)
	}
}
