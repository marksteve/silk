package silk_kv

import (
	"encoding/json"
	"errors"

	"github.com/charmbracelet/charm/client"
	"github.com/charmbracelet/charm/kv"
	charm "github.com/charmbracelet/charm/proto"
)

// Link

type linkHandler struct {
	tokenSent     chan struct{}
	validToken    chan bool
	success       chan bool // true if the key was already linked
	requestDenied chan struct{}
	timeout       chan struct{}
	err           chan error
}

func newLinkHandler() *linkHandler {
	return &linkHandler{
		tokenSent:     make(chan struct{}),
		validToken:    make(chan bool),
		success:       make(chan bool),
		requestDenied: make(chan struct{}),
		timeout:       make(chan struct{}),
		err:           make(chan error),
	}
}

func (lh *linkHandler) TokenCreated(l *charm.Link) {
	// Not implemented for the link participant
}

func (lh *linkHandler) TokenSent(l *charm.Link) {
	lh.tokenSent <- struct{}{}
}

func (lh *linkHandler) ValidToken(l *charm.Link) {
	lh.validToken <- true
}

func (lh *linkHandler) InvalidToken(l *charm.Link) {
	lh.validToken <- false
}

func (lh *linkHandler) Request(l *charm.Link) bool {
	// Not implemented for the link participant
	return false
}

func (lh *linkHandler) RequestDenied(l *charm.Link) {
	lh.requestDenied <- struct{}{}
}

func (lh *linkHandler) SameUser(l *charm.Link) {
	lh.success <- true
}

func (lh *linkHandler) Success(l *charm.Link) {
	lh.success <- false
}

func (lh *linkHandler) Timeout(l *charm.Link) {
	lh.timeout <- struct{}{}
}

func (lh *linkHandler) Error(l *charm.Link) {
	lh.err <- errors.New("error")
}

func Link(code string) error {
	cc, err := client.NewClientWithDefaults()
	if err != nil {
		return err
	}
	lh := newLinkHandler()
	cc.Link(lh, code)
	return nil
}

// Keys

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

// GetOptions

func GetOptions(name string) (string, error) {
	db, err := openKV(name)
	if err != nil {
		return "failed_to_open", err
	}
	defer db.Close()
	opts := db.DB.Opts()
	return toJSON(opts)
}

// Utils

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
