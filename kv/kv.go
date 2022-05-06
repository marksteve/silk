package kv

import (
	"github.com/charmbracelet/charm/client"
)

func Link(code string) error {
	cc, err := client.NewClientWithDefaults()
	if err != nil {
		return err
	}
	lh := newLinkHandler()
	cc.Link(lh, code)
	return nil
}

func Keys(name string) (string, error) {
	db, err := openKV(name)
	if err != nil {
		return "failed_to_open", err
	}
	defer db.Close()

	err = db.Sync()
	if err != nil {
		return "failed_to_sync", err
	}

	b, err := db.Keys()
	if err != nil {
		return "failed_to_get_keys", err
	}

	var keys []string
	for _, k := range b {
		keys = append(keys, string(k))
	}
	return toJSON(keys)
}

func GetOptions(name string) (string, error) {
	db, err := openKV(name)
	if err != nil {
		return "failed_to_open", err
	}
	defer db.Close()
	opts := db.DB.Opts()
	return toJSON(opts)
}
