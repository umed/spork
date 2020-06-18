package config

import "flag"

// ParseAddr parses service address
func ParseAddr() *string {
	var addr = flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()
	return addr
}
