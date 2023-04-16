// Copyright (c) 2023 Geo-Studios - All Rights Reserved.

package gens

import (
	"os"
	"strings"
)

func GenerateNodejs(name, ver string) {

	os.Mkdir(name, 0700)
	os.Chdir(name)
	pj, _ := os.Create("project.json")
	pj.WriteString(strings.ReplaceAll(strings.ReplaceAll(`{
	"ProjectName":"$pn",
	"ProjectVer":"$pv"
}`, "$pn", name), "$pv", ver))
	pj.Close()
	ba, _ := os.Create("run.bat")
	ba.WriteString(`node ./src/` + name + ".js")
	ba.Close()
	os.Mkdir("src", 0700)
	os.Chdir("src")
	pf, _ := os.Create(name + ".js")
	pf.WriteString(`function main(){
		console.log("Hello World")
	}
	main()`)
	pf.Close()

}
