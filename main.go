package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/blang/vfs"
	"github.com/blang/vfs/memfs"
)

const (
	hw = "Hello World!\n"
	hwPath = "/tmp/helloworld.txt"
)

func main() {
	fmt.Print("Instantiating memfs with tmp directory...")
	mfs := memfs.Create()
	mfs.Mkdir("/tmp", 0700)
	fmt.Println("Done.")

	fmt.Print("Instantiating read-only wrapper of memfs...")
	rofs := vfs.ReadOnly(mfs)
	fmt.Println("Done.")

	fmt.Print("Opening hello world file...")
	wf, err := mfs.OpenFile(hwPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_SYNC, 0600)
	check(err)
	fmt.Println("Done.")

	fmt.Print("Writing contents to file and closing...")
	n, err := wf.Write([]byte(hw))
	check(err)
	err = wf.Close()
	check(err)
	fmt.Println("Done.")
	fmt.Printf("Bytes written: %d\n", n)

	fmt.Print("Obtaining and displaying file meta...")
	rfStat, err := rofs.Stat(hwPath)
	check(err)
	fmt.Println("Done.")
	fmt.Printf("ROFS file size: %d\n", rfStat.Size())

	fmt.Print("Reading contents of hello world file...")
	contents, err := vfs.ReadFile(rofs, hwPath)
	check(err)
	fmt.Println("Done.")
	fmt.Printf("Bytes Read: %d\n", len(contents))

	fmt.Println("File Contents:")
	fmt.Print(strings.Trim(string(contents), " "))
}

func check(err error) {
	if err != nil {
		fmt.Println("ERROR")
		log.Panicln(err)
	}
}
