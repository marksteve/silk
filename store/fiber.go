package store

type Fiber struct {
	TS       string `json:"ts"`
	Data     []byte `json:"data"`
	Mimetype string `json:"mimetype"`
}
