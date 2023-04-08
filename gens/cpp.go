package gens

import (
	"os"
	"strings"
)

func GenerateCpp(name, ver string) {

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
	pf, _ := os.Create(name + ".cpp")
	pf.WriteString(`#include <iostream>

int main() {
	std::cout << "Hello World!";
	return 0;
}`)
	pf.Close()

}
