package main

import "flag"

// CmdArgs is a structure meant to store command line passed arguments
type CmdArgs struct {
	token string
	proxy string
}

// Parse parses command line arguments and returns them as CmdArgs structure
func Parse() CmdArgs {
	var cmdArgs CmdArgs
	flag.StringVar(&cmdArgs.token, "token", "", "a string")
	flag.StringVar(&cmdArgs.proxy, "proxy", "", "a string")
	flag.Parse()
	return cmdArgs
}
