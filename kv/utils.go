package kv

import (
	"encoding/json"

	"github.com/charmbracelet/charm/kv"
)

func toJSON(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "failed_to_marshal", err
	}
	return string(b), nil
}

func openKV(name string) (*kv.KV, error) {
	if name == "" {
		name = "silk"
	}
	return kv.OpenWithDefaults(name)
}
