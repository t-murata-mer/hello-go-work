package main

import (
	"config"
	"libs"

	"log"
)

func main() {

	goPracMain()
	config.OreConfig()
	libs.OreLibs()
}

func goPracMain() {
	log.Println("go-prac main.go")
}
