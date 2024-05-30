package main

import (
	"fmt"
	_ "unsafe"

	"github.com/johnayman0/sing-box-extra/boxbox"
	"github.com/johnayman0/sing-box-extra/boxmain"
	_ "github.com/johnayman0/sing-box-extra/distro/all"
)

func main() {
	fmt.Println("sing-box-extra:", boxbox.Version)
	fmt.Println()

	// sing-box
	boxmain.Main()
}
