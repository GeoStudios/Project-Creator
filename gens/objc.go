// Copyright (c) 2023 Geo-Studios - All Rights Reserved.

package gens

import (
	"os"
	"strings"
)

func GenerateObjC(name, ver string) {

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
sources = $(wildcard src/*.m)
objects = $(sources:.c=.o)
flags = -O0 -g

$(exec): $(objects)
	gcc $(objects) $(flags) -o $(exec)

%.o: %.cpp include/%.hh
	gcc -c $(flags) $< -o $@

clean:
	-rm *.exe
	-rm *.o
	-rm src/*.o`)

	os.Mkdir("src", 0700)
	os.Chdir("src")
	pf, _ := os.Create(name + ".m")
	pf.WriteString(`@interface SampleClass:NSObject
- (void)sampleMethod;
@end

@implementation SampleClass

- (void)sampleMethod {
   NSLog(@"Hello, World! \n");
}

@end

int main() {
   /* my first program in Objective-C */
   SampleClass *sampleClass = [[SampleClass alloc]init];
   [sampleClass sampleMethod];
   return 0;
}`)
	pf.Close()

}
