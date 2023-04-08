package gens

import (
	"os"
	"strings"
)

func GenerateC(name, ver string) {

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
sources = $(wildcard src/*.c)
objects = $(sources:.c=.o)
flags = -g

$(exec): $(objects)
	gcc $(objects) $(flags) -o $(exec)

%.o: %.cpp include/%.h
	gcc -c $(flags) $< -o $@

clean:
	-rm *.exe
	-rm *.o
	-rm src/*.o`)

	os.Mkdir("src", 0700)
	os.Chdir("src")
	pf, _ := os.Create(name + ".c")
	pf.WriteString(`#include <stdio.h>

int main(int argc, char const *argv[])
{
	printf(argc);
	printf("Hello World");
	return 0;
}`)
	pf.Close()

}
