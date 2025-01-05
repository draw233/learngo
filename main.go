package main

import (
	"fmt"
	"learngo/variabledemo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("hello  go")
	variabledemo.VarType()

	var (
		// Stdin  = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
		Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
		// Stderr = os.NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	)
	Stdout.WriteString("Hello 世界!")
	ctrlC()
}

func ctrlC() {
	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
