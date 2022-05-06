package store

import (
	"github.com/charmbracelet/charm/kv"
)

func openKV(name string) (*kv.KV, error) {
	if name == "" {
		name = "silk"
	}
	return kv.OpenWithDefaults(name)
}
