package main

import (
	"os"
)

//goland:noinspection GoTestName
func ExampleVersion() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"releasecontroller-mcp-server", "--version"}
	main()
	// Output: 0.0.0
}
