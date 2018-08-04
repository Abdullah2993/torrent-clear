package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/anacrolix/torrent/metainfo"
)

func main() {
	var comment, createdBy string
	flag.StringVar(&comment, "comment", "", "Comment text")
	flag.StringVar(&createdBy, "created", "", "CreatedBy text")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s TORRNET_FILE NEW_TORRENT_FILE", os.Args[0])
		os.Exit(1)
	}

	m, err := metainfo.LoadFromFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open file: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Orignal Info Hash: %s\r\n", m.HashInfoBytes().HexString())
	m.Comment = comment
	m.CreatedBy = createdBy

	fmt.Printf("After Info Hash: %s\r\n", m.HashInfoBytes().HexString())

	f, err := os.OpenFile(flag.Arg(1), os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 644)
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
