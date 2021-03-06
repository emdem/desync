package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/folbricht/desync"
)

const usage = `desync <command> [options]
desync <command> -h

Commands:
extract     - build a blob from a caibx file
verify      - verify the integrity of a local store
list-chunks - list all chunk IDs contained in a caibx
cache       - populate a cache without writing to a blob
chop        - split a blob based on existing caibx and store the chunks
`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
	}

	if len(os.Args) < 2 {
		die(errors.New("No command given. See -h for help."))
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "-h":
		flag.Usage()
		os.Exit(1)
	case "extract":
		extract(args)
	case "verify":
		verify(args)
	case "cache":
		cache(args)
	case "list-chunks":
		list(args)
	case "chop":
		chop(args)
	default:
		die(fmt.Errorf("Unknown command %s", cmd))
	}
}

func readCaibxFile(name string) (c desync.Caibx, err error) {
	f, err := os.Open(name)
	if err != nil {
		return
	}
	defer f.Close()
	return desync.CaibxFromReader(f)
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

type multiArg struct {
	list []string
}

func (a *multiArg) Set(v string) error {
	a.list = append(a.list, v)
	return nil
}

func (a *multiArg) String() string { return "" }
