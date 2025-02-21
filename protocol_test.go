package protocol

import (
	"bytes"
	"testing"
)

// TestPayloadEncoding tests that a message can be encoded with a message type header
// correctly. It creates a message, encodes it, and checks that the decoded headers
// are equal to the original message type.
func TestHeartBeatPayloadEncoding(t *testing.T) {

	buf := new(bytes.Buffer)

	var payload = new(Beat)

	_, err := Encode(buf, payload, HeartBeat)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// TestHearBeatPayloadDecoding tests that a heartbeat message can be encoded and decoded
// correctly. It creates a heartbeat message, encodes it, decodes it, and checks that the
// decoding was successful.
func TestHearBeatPayloadDecoding(t *testing.T) {
	buf := new(bytes.Buffer)

	var payload = new(Beat)

	_, err := Encode(buf, payload, HeartBeat)
	if err != nil {
		t.Errorf("Expected no Encoding error, got %v", err)
	}

	_, err = Decode(buf)
	if err != nil {
		t.Errorf("Expected no Decoding error, got %v", err)
	}

}

// TestPayloadDecode tests that a message can be encoded and then decoded
// correctly. It creates a message, encodes it, decodes it, and checks that the
// decoded message is equal to the original message.
func TestMessagePayloadEncoding(t *testing.T) {
	message := Message{
		Message: "Hello, world",
		To:      "John Doe",
		From:    "Jane Doe",
	}

	buf := new(bytes.Buffer)

	_, err := Encode(buf, &message, MessageType)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// TestMessagePayloadDecoding tests that a message can be decoded correctly.
// It creates a message, encodes it, decodes it, and checks that the decoded
// message is equal to the original message.
func TestMessagePayloadDecoding(t *testing.T) {
	message := Message{
		Message: "Hello, world",
		To:      "John Doe",
		From:    "Jane Doe",
	}

	buf := new(bytes.Buffer)

	_, err := Encode(buf, &message, MessageType)

	if err != nil {
		t.Errorf("Expected no Encoding error, got %v", err)
	}

	m, err := Decode(buf)

	if err != nil {
		t.Errorf("Expected not Decoding error got: %v", err)
	}

	if message.String() != m.String() {
		t.Errorf("Expected m to be : %s got : %s", message.String(), m.String())
	}
}

// TestErrorPayloadEncoding tests that a error message can be encoded
// correctly. It creates a error message, encodes it, and checks that the
// encoding was successful.
func TestErrorPayloadEncoding(t *testing.T) {
	errPayload := Error_("Something went wrong")

	buf := new(bytes.Buffer)

	_, err := Encode(buf, &errPayload, Error)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

}

// TestDecodeErrorPayload tests that a error message can be encoded and then decoded
// correctly. It creates a error message, encodes it, decodes it, and checks that the
// decoded error message is equal to the original error message.
func TestDecodeErrorPayload(t *testing.T) {
	errPayload := Error_("Something went wrong")

	buf := new(bytes.Buffer)

	_, err := Encode(buf, &errPayload, Error)

	if err != nil {
		t.Errorf("Expected no Encoding error, got %v", err)
	}

	m, err := Decode(buf)

	if err != nil {
		t.Errorf("Expected not Decoding error got: %v", err)
	}

	if errPayload.String() != m.String() {
		t.Errorf("Expected m to be : %s got : %s", errPayload.String(), m.String())
	}
}
