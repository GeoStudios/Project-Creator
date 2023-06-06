// Copyright (c) 2023 Geo-Studios - All Rights Reserved.

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
	makeF, _ := os.Create("Makefile")
	makeF.WriteString(`exec = hello.exe
sources = $(wildcard src/*.cpp)
objects = $(sources:.c=.o)
flags = -g

$(exec): $(objects)
	g++ $(objects) $(flags) -o $(exec)

%.o: %.cpp include/%.hh
	g++ -c $(flags) $< -o $@

clean:
	-rm *.exe
	-rm *.o
	-rm src/*.o`)

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
