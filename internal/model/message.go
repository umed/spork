package model

import "bytes"

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// HandleMessageSpaces handles spaces and newlines in byte sequences
func HandleMessageSpaces(message []byte) []byte {
	return bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
}
