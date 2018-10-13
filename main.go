package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/anacrolix/torrent/metainfo"
)

func main() {
	var comment, createdBy string
	var install bool
	flag.StringVar(&comment, "comment", "", "Comment text")
	flag.StringVar(&createdBy, "created", "", "CreatedBy text")
	flag.BoolVar(&install, "install", false, "Install context menu")
	flag.Parse()

	if install {
		exePath, err := os.Executable()
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to get executable path: %v", err)
			os.Exit(1)
		}
		err = Install(exePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to install: %v", err)
			os.Exit(1)
		}
		return
	}

	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "USAGE: %s TORRNET_FILE [NEW_TORRENT_FILE]", os.Args[0])
		os.Exit(1)
	}

	file := flag.Arg(0)
	nfile := flag.Arg(0)

	if len(flag.Args()) == 2 {
		nfile = flag.Arg(1)
	}

	m, err := metainfo.LoadFromFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open file: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Orignal Info Hash: %s\r\n", m.HashInfoBytes().HexString())
	m.Comment = comment
	m.CreatedBy = createdBy

	fmt.Printf("After Info Hash: %s\r\n", m.HashInfoBytes().HexString())

	f, err := os.OpenFile(nfile, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 644)
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
