// (cd demo; make); go run coypud.go -r demo/coypu/
package main

// Thanks to https://go.dev/src/net/example_test.go
//
// Demo:
//
// (sleep 2; echo '31------------compensation                                            ') | telnet localhost 31525
//
// (sleep 2; echo '30                                                                    ') | telnet localhost 31525

import (
	"flag"
	"log"
	"net"
	"os"
	"path/filepath"
	//"fmt"
	"strings"
)

const RequestLen = 64
const ReplyLen = 1024

var LISTEN = flag.String("l", ":31525", "TCP port to bind to and listen on")
var ROOT = flag.String("r", ".", "Filesystem root to serve")

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *LISTEN)
	if err != nil {
		log.Fatalf("cannot Listen: tcp %q: %v", *LISTEN, err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			defer c.Close()
			Handle(c)
		}(conn)
	}
}

func Handle(conn net.Conn) {
	bb := make([]byte, RequestLen)
	_, err := conn.Read(bb)
	if err != nil {
		log.Panic(err)
	}

	// Ignore the IP and Port at bb[2:14].  Join type and name with "-".
	sel := string(bb[:2]) + "-" + string(bb[14:])

	// Clear spaces and crap from the right.
	sel = strings.TrimRight(sel, " \000\032") // \032 is Control-Z

	if strings.Contains(sel, "/") {
		log.Panicf("Selector should not contain '/': %q", sel)
	}
	if strings.Contains(sel, ":") {
		log.Panicf("Selector should not contain ':': %q", sel)
	}
	if strings.Contains(sel, ".") {
		log.Panicf("Selector should not contain '.': %q", sel)
	}

	filename := filepath.Join(*ROOT, sel)
	r, err := os.Open(filename)
	if err != nil {
		log.Panicf("Cannot open selector %q: %v", filename, err)
	}
	defer r.Close()

	guts := make([]byte, ReplyLen)
	for i, _ := range guts {
		guts[i] = 26 // control-Z
	}
	_, _ = r.Read(guts) // ignore if error, just return control-Zs.
	conn.Write(guts)
}
