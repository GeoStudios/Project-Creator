package gens

import (
	"os"
	"strings"
)

func GenerateJava(name, ver string) {

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
	pf, _ := os.Create(name + ".java")
	pf.WriteString(`public class ` + name + ` {

	public static void main(String[] args) {
		System.out.println(args)
	
		System.out.println("Hello World");
	
	}
	
}`)
	pf.Close()

}
