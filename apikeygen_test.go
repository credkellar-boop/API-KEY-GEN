package main

import (
	"strings"
	"testing"
)

func TestGeneratePrefix(t *testing.T) {
	generator := NewApiKeyGen("test_go_")
	data, err := generator.Generate()
	
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !strings.HasPrefix(data.RawKey, "test_go_") {
		t.Errorf("Expected raw key to start with 'test_go_', got %s", data.RawKey)
	}
}

func TestGenerateUniqueness(t *testing.T) {
	generator := NewApiKeyGen("sk_live_")
	key1, _ := generator.Generate()
	key2, _ := generator.Generate()
	
	if key1.RawKey == key2.RawKey {
		t.Errorf("Expected unique keys, got identical keys: %s", key1.RawKey)
	}
}

func TestGenerateHashLength(t *testing.T) {
	generator := NewApiKeyGen("sk_live_")
	data, _ := generator.Generate()
	
	if len(data.DBHash) != 64 {
		t.Errorf("Expected SHA-256 hash to be 64 characters, got %d", len(data.DBHash))
	}
	if data.RawKey == data.DBHash {
		t.Errorf("Raw key and DB hash must not be identical")
	}
}
