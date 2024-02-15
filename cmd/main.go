package main

import (
	"fmt"
	"go-send/cmd/flags"
)

func main() {
	fmt.Println("Go <=> Amumu")
	fmt.Println()
	amumuCmd := flags.NewAmumuCmd()
	flags.SetCmdExec()
	amumuCmd.Run()
}
