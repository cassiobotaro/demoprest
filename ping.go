package main

var (
	// HTTPVars route variables for the current request
	HTTPVars map[string]string
	// URLQuery parses RawQuery and returns the corresponding values
	URLQuery map[string][]string
)

// GETPingHandler plugin
//
// BUILD:
// go build -o plugins/ping.so -buildmode=plugin ping.go
func GETPingHandler() (ret string) {
	return "Pong!"
}
