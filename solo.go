/*

Variation of Tim Kay's solo (http://timkay.com/solo/) written in Go (Golang)

*/

package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

var port string
var verbose, silent, noport bool
var sleep int64
var timeout int

const usage string = `

Usage: solo -port=PORT COMMAND

where
	PORT		some arbitrary port number to be used for locking
	COMMAND		shell command to run

options
    -verbose	be verbose
    -silent		be silent
    -sleep      sleep n seconds before running command
    -noport     do not specify port, i.e. do not implement locking

example:

* * * * * solo -port=3801 ./job.pl blah blah
`

func main() {

	flag.Int64Var(&sleep, "sleep", 0, "sleep X seconds before running command")
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.BoolVar(&silent, "silent", false, "omit locking message")
	flag.BoolVar(&noport, "noport", false, "omit port, i.e. do not implement locking")
	flag.StringVar(&port, "port", "", "arbitrary port number")

	flag.Parse()

	if port == "" && !noport {
		log.Fatal(usage)
	}

	if port != "" {
		if verbose {
			log.Println("solo: binding on port", port)
		}
		if lock, listenerr := net.Listen("tcp", ":"+port); listenerr != nil {
			if silent {
				os.Exit(1)
			} else {
				log.Println("solo(", port, ")")
				log.Fatal(listenerr)
			}
		} else {
			defer lock.Close()
		}
	}

	if sleep > 0 {
		time.Sleep(time.Duration(sleep) * time.Second)
	}
	cmd := exec.Command(flag.Args()[0], flag.Args()[1:]...)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
