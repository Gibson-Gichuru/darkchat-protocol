package protocol

import (
	"errors"
	"fmt"
	"io"
)

const (
	MessageType uint8 = iota + 1
	HeartBeat
	Error
	MaxPayloadsize uint32 = 10 << 20
)

var ErrorMaxPayloadSize = errors.New("max payload size exceeded")
var ErrorEmptyHeaders = errors.New("empty headers")
var ErrorUnknownType = errors.New("unknown message type")

type Payload interface {
	fmt.Stringer
	io.WriterTo
	io.ReaderFrom
	Byte() []byte
}

type PayloadHeaders struct {
	Size     uint32
	Type     uint8
	Encoding string
}
