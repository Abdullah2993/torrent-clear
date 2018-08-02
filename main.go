package main

import (
	"fmt"
	"os"

	"github.com/anacrolix/torrent/metainfo"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "USAGE: %s TORRNET_FILE NEW_TORRENT_FILE", os.Args[0])
		os.Exit(1)
	}

	m, err := metainfo.LoadFromFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open file: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Orignal Info Hash: %s\r\n", m.HashInfoBytes().HexString())
	m.Comment = ""
	m.CreatedBy = ""

	fmt.Printf("After Info Hash: %s\r\n", m.HashInfoBytes().HexString())

	f, err := os.OpenFile(os.Args[2], os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open file for writing: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	err = m.Write(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to write file: %v", err)
		os.Exit(1)
	}
}
