/*

Adaptation of Tim Kay's solo (http://timkay.com/solo/) to the Go programming language

*/

package main

import (
	// "net"
	"flag"
	"log"
	"os/exec"
	"time"
)

var port int
var verbose bool
var silent bool
var sleep int64
var timeout int

const usage string = `Usage: solo -port=PORT COMMAND

where
	PORT		some arbitrary port number to be used for locking
	COMMAND		shell command to run

options
	-verbose	be verbose
	-silent		be silent

example:

* * * * * solo -port=3801 ./job.pl blah blah
`

func main() {

	flag.Int64Var(&sleep, "sleep", 0, "sleep X seconds before running command")
	flag.IntVar(&timeout, "timeout", 0, "timeout after X seconds before running command")
	flag.IntVar(&port, "port", 0, "port to bind to for locking")
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.BoolVar(&silent, "silent", false, "omit locking message")

	flag.Parse()
	args := flag.Args()

	if sleep > 0 {
		time.Sleep(time.Duration(sleep) * time.Second)
	}
	command := args[0]
	args = args[1:]
	cmd := exec.Command(command, args...)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
