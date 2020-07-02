package model

import "bytes"

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// MessageType represents type of data within @Message
type MessageType int

const (
	// Text messsage type
	Text MessageType = iota
	// Binary message type
	Binary MessageType = iota
)

// Message represents message that will be sent to/from client
type Message struct {
	client *Client

	dataType MessageType
	data     []byte
}

// HandleMessageSpaces handles spaces and newlines in byte sequences
func HandleMessageSpaces(message []byte) []byte {
	return bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
}
