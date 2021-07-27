package designPattern

type Status uint8

const (
	Stopped Status = iota
	Started
)

type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}

type Body struct {
	Items []string
}

type Message struct {
	Header *Header
	Body   *Body
}