package main

import (
	"flag"
	"fmt"
	"log"
	//"os/user"
	"path/filepath"

	. "github.com/kkdai/youtube"
)

func main() {
	flag.Parse()
	log.Println("ARGS = 1) url 2) folder 3 filename")
	log.Println(flag.Args())
	//usr, _ := user.Current()
	currentDir := flag.Arg(1) //fmt.Sprintf("%v/Videos/youtube", usr.HomeDir)
	log.Println("download to dir=", currentDir)
	y := NewYoutube(true)
	arg := flag.Arg(0)
	fileName := flag.Arg(2)
	if err := y.DecodeURL(arg); err != nil {
		fmt.Println("err:", err)
	}
	if err := y.StartDownload(filepath.Join(currentDir, fileName)); err != nil {
		fmt.Println("err:", err)
	}

}
