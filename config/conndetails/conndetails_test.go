package conndetails_test

import (
	"testing"

	"github.com/crossplane-contrib/provider-checkly/config/conndetails"
)

func TestStringKeys(t *testing.T) {
	fn := conndetails.StringKeys(map[string]string{
		"url":   "url",
		"token": "token",
		"key":   "dashboard_key",
	})

	t.Run("all present", func(t *testing.T) {
		out, err := fn(map[string]any{"url": "https://example.com", "token": "tok", "key": "k"})
		if err != nil {
			t.Fatal(err)
		}
		if string(out["url"]) != "https://example.com" {
			t.Fatalf("unexpected url: %q", out["url"])
		}
		if string(out["token"]) != "tok" {
			t.Fatalf("unexpected token: %q", out["token"])
		}
		if string(out["dashboard_key"]) != "k" {
			t.Fatalf("unexpected dashboard_key: %q", out["dashboard_key"])
		}
	})

	t.Run("partial", func(t *testing.T) {
		out, err := fn(map[string]any{"url": "u"})
		if err != nil {
			t.Fatal(err)
		}
		if len(out) != 1 {
			t.Fatalf("expected 1 entry, got %d", len(out))
		}
	})

	t.Run("empty strings skipped", func(t *testing.T) {
		out, err := fn(map[string]any{"url": ""})
		if err != nil {
			t.Fatal(err)
		}
		if len(out) != 0 {
			t.Fatalf("expected 0 entries, got %d", len(out))
		}
	})

	t.Run("empty attr", func(t *testing.T) {
		out, err := fn(map[string]any{})
		if err != nil {
			t.Fatal(err)
		}
		if len(out) != 0 {
			t.Fatalf("expected 0 entries, got %d", len(out))
		}
	})
}

func TestIndexedSlice(t *testing.T) {
	fn := conndetails.IndexedSlice("keys", "api_key")

	t.Run("single key", func(t *testing.T) {
		out, err := fn(map[string]any{"keys": []any{"abc123"}})
		if err != nil {
			t.Fatal(err)
		}
		if string(out["api_key"]) != "abc123" {
			t.Fatalf("expected api_key=abc123, got %q", out["api_key"])
		}
	})

	t.Run("multiple keys", func(t *testing.T) {
		out, err := fn(map[string]any{"keys": []any{"k0", "k1", "k2"}})
		if err != nil {
			t.Fatal(err)
		}
		if string(out["api_key"]) != "k0" {
			t.Fatalf("expected api_key=k0, got %q", out["api_key"])
		}
		if string(out["api_key_1"]) != "k1" {
			t.Fatalf("expected api_key_1=k1, got %q", out["api_key_1"])
		}
		if string(out["api_key_2"]) != "k2" {
			t.Fatalf("expected api_key_2=k2, got %q", out["api_key_2"])
		}
	})

	t.Run("empty attr", func(t *testing.T) {
		out, err := fn(map[string]any{})
		if err != nil {
			t.Fatal(err)
		}
		if len(out) != 0 {
			t.Fatalf("expected empty map, got %v", out)
		}
	})

	t.Run("empty strings skipped", func(t *testing.T) {
		out, err := fn(map[string]any{"keys": []any{"", ""}})
		if err != nil {
			t.Fatal(err)
		}
		if len(out) != 0 {
			t.Fatalf("expected empty map, got %v", out)
		}
	})
}
