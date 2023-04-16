// Copyright (c) 2023 Geo-Studios - All Rights Reserved.

package gens

import (
	"os"
	"strings"
)

func GeneratePython(name, ver string) {

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
	pf, _ := os.Create(name + ".py")
	pf.WriteString(`import sys
	
def main():
	OsArgs = sys.argv

	print(OsArgs)
	
	print("Hello World")
	
main()`)
	pf.Close()

}
