package main

import (
    "flag"
    "fmt"
    "os"
)

var name string
/*
	更灵活地定制命令参数容器，定制不会影响到全局变量flag.CommandLine
 */
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
    cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	//os.Args[1:]指给定的命令参数
    cmdLine.Parse(os.Args[1:])
    fmt.Printf("Hello, %s!\n", name)
}