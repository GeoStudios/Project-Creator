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

	jarbuilder, _ := os.Create("Compiler.bat")
	jarbuilder.WriteString(`@echo off

set "JarName=jarname" 
set "MainClass=jarclass"
set "KeepSrcOut=false"
	
CLS
MKDIR "bin"
DEL "bin\%JarName%.jar" /Q /S
RMDIR srcOut /Q /S
Xcopy src srcOut /e/h/c/i
cd srcOut
javac *.java
jar cf "%JarName%.jar" *.class *
jar --update --verbose --file "%JarName%.jar" --main-class "%MainClass%"
copy "%JarName%.jar" "../bin/"
DEL "%JarName%.jar"
cd ../
if "%KeepSrcOut%" == "false" RMDIR srcOut /Q /S
CLS
echo Compilation is done and your "%JarName%.jar" file is ready sir.
java -jar "bin/%JarName%.jar"`)

	jarbuilder.Close()

	jbdoc, _ := os.Create("jarbuilder.md")
	jbdoc.WriteString(`# Jar Compiler

### Config
batch
` + "```" + `
set "JarName=jarname" 
set "MainClass=jarclass"
set "KeepSrcOut=false"
` + "```" + `
	
Replace "jarname" with the program name and leave out extentions
and replace "jarclass" with the class that open the program and leave the program to do the rest.

The required objects to make this work are Compiler.bat and the src folder. The program will create and "srcOut" folder with all the compiled files and pack them into a jar while making a bin folder for the jar to stay in. The Jar file will autorun so you can mess with the finished product and check for errors that may occur in the console/terminal.
	
After the jar is created it will delete the srcOut unless you change the KeepSrcOut to true.`)

	os.Mkdir("src", 0700)
	os.Chdir("src")
	pf, _ := os.Create(name + ".java")
	pf.WriteString(`public class ` + name + ` {

	public static void main(String[] args) {
		System.out.println(args);
	
		System.out.println("Hello World");
	
	}
	
}`)
	pf.Close()

}
