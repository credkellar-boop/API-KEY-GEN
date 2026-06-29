package main // Must match main.go to access NewApiKeyGen

import (
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	gen := NewApiKeyGen("test_")
	raw, _ := gen.Generate()
	
	if !strings.HasPrefix(raw, "test_") {
		t.Errorf("Prefix mismatch")
	}
}
