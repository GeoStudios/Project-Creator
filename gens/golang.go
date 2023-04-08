package gens

import (
	"os"
	"strings"
)

func GenerateGolang(name, ver string) {

	os.Mkdir(name, 0700)
	os.Chdir(name)
	pj, _ := os.Create("project.json")
	pj.WriteString(strings.ReplaceAll(strings.ReplaceAll(`{
	"ProjectName":"$pn",
	"ProjectVer":"$pv"
}`, "$pn", name), "$pv", ver))
	pj.Close()

	os.Mkdir("src", 0700)
	os.Chdir("src")
	pf, _ := os.Create(name + ".go")
	pf.WriteString(`package main

import (
	"fmt"
	"os"
)
	
func main() {
	OsArgs := os.Args
	fmt.Println(OsArgs)

	fmt.Println("Hello World")

}`)
	pf.Close()
	gm, _ := os.Create("go.mod")
	gm.WriteString(`module ` + name + `

go 1.19`)
	gm.Close()

}
