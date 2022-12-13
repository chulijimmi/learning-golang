package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func initConstant() {
	size := func(b ByteSize) string {
		switch {
		case b >= YB:
			return fmt.Sprintf("%.2fYB", b/YB)
		case b >= ZB:
			return fmt.Sprintf("%.2fZB", b/ZB)
		case b >= EB:
			return fmt.Sprintf("%.2fEB", b/EB)
		case b >= PB:
			return fmt.Sprintf("%.2fPB", b/PB)
		case b >= TB:
			return fmt.Sprintf("%.2fTB", b/TB)
		case b >= GB:
			return fmt.Sprintf("%.2fGB", b/GB)
		case b >= MB:
			return fmt.Sprintf("%.2fMB", b/MB)
		case b >= KB:
			return fmt.Sprintf("%.2fKB", b/KB)
		}
		return fmt.Sprintf("%.2fB", b)
	}
	num := 1e10
	fmt.Printf("num: %v\n", num)
	a := size(1e10)
	fmt.Printf("a: %v\n", a)

}

func initVar() {
	var (
		home   = os.Getenv("HOME")
		user   = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)
	fmt.Println("=============The init function=============")

	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

func main() {
	fmt.Println("=============Constants=============")
	initConstant()
	fmt.Println("=============Variables=============")
	initVar()
}
