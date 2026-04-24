// Package conndetails provides factory functions for building
// AdditionalConnectionDetailsFn implementations used by Upjet resource
// configurators. Each factory returns a function with the signature
// expected by ujconfig.Resource.Sensitive.AdditionalConnectionDetailsFn.
package conndetails

import "fmt"

// StringKeys returns a connection-detail extractor that reads the given
// attribute keys as strings and maps them to the specified output names.
// Keys whose values are missing or empty are silently skipped.
//
// Example:
//
//	conndetails.StringKeys(map[string]string{"url": "url", "token": "token"})
func StringKeys(mappings map[string]string) func(map[string]any) (map[string][]byte, error) {
	return func(attr map[string]any) (map[string][]byte, error) {
		out := make(map[string][]byte, len(mappings))
		for attrKey, connKey := range mappings {
			if v, ok := attr[attrKey].(string); ok && v != "" {
				out[connKey] = []byte(v)
			}
		}
		return out, nil
	}
}

// IndexedSlice returns a connection-detail extractor for list-type
// attributes. The first element is stored under `prefix`, subsequent
// elements under `prefix_1`, `prefix_2`, etc.
//
// Example:
//
//	conndetails.IndexedSlice("keys", "api_key")
func IndexedSlice(attrKey, prefix string) func(map[string]any) (map[string][]byte, error) {
	return func(attr map[string]any) (map[string][]byte, error) {
		out := map[string][]byte{}
		v, ok := attr[attrKey]
		if !ok {
			return out, nil
		}
		items, ok := v.([]any)
		if !ok {
			return out, nil
		}
		for i, item := range items {
			s, ok := item.(string)
			if !ok || s == "" {
				continue
			}
			key := prefix
			if i > 0 {
				key = fmt.Sprintf("%s_%d", prefix, i)
			}
			out[key] = []byte(s)
		}
		return out, nil
	}
}
