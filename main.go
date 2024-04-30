package main

import (
	"fmt"
	_ "unsafe"

	"github.com/Mahdi-zarei/sing-box-extra/boxbox"
	"github.com/Mahdi-zarei/sing-box-extra/boxmain"
	_ "github.com/Mahdi-zarei/sing-box-extra/distro/all"
)

func main() {
	fmt.Println("sing-box-extra:", boxbox.Version)
	fmt.Println()

	// sing-box
	boxmain.Main()
}
