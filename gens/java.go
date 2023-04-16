// Copyright (c) 2023 Geo-Studios - All Rights Reserved.

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
	
	set "javaDir="
	set "experiments=false"
	
	@REM Only works if experiments are enabled
	set "javaBaseDir="
	set "javaJdkVer=jdk-20"
	
	
	
	@Rem IMPORTANT DO NOT CHANGE
	set "java=%javaBaseDir%\%javaJdkVer%\bin\"
	
	CLS
	MKDIR "bin"
	DEL "bin\%JarName%.jar" /Q /S
	RMDIR srcOut /Q /S
	Xcopy src srcOut /e/h/c/i
	cd srcOut
	if "%experiments%" == "false" "%javaDir%javac" *.java
	if "%experiments%" == "false" "%javaDir%jar" cf "%JarName%.jar" *.class *
	if "%experiments%" == "false" "%javaDir%jar" --update --verbose --file "%JarName%.jar" --main-class "%MainClass%"
	
	if "%experiments%" == "true" "%java%javac" *.java
	if "%experiments%" == "true" "%java%jar" cf "%JarName%.jar" *.class *
	if "%experiments%" == "true" "%java%jar" --update --verbose --file "%JarName%.jar" --main-class "%MainClass%"
	copy "%JarName%.jar" "../bin/"
	DEL "%JarName%.jar"
	cd ../
	if "%KeepSrcOut%" == "false" RMDIR srcOut /Q /S
	CLS
	echo Compilation is done and your "%JarName%.jar" is ready sir.
	if "%experiments%" == "false" "%java%java" -jar "bin/%JarName%.jar"
	
	if "%experiments%" == "true" "%java%java" -jar "bin/%JarName%.jar"`)

	jarbuilder.Close()

	jbdoc, _ := os.Create("jarbuilder.md")
	jbdoc.WriteString(`# Jar Builder ` + "`" + `v1.1` + "`" + `

	### Config

	` + "```" + `batch
	
	set "JarName=jarname" 
	set "MainClass=jarclass"
	set "KeepSrcOut=false"
	
	set "javaDir="
	` + "```" + `
	When Using This program leave "javaDir" blank if you want to use your default java used by your system or you could put a custom java dir like ` + "`" + `E:\Program Files\Java\jdk-13\bin\` + "`" + ` for example and make sure your path has ` + "`" + `bin/` + "`" + ` at the end or the program will not work. 
	
	Replace "jarname" with the program name and leave out extentions
	and replace "jarclass" with the class that open the program and leave the program to do the rest.
	
	The required objects to make this work are Compiler.bat and the src folder. The program will create and "srcOut" folder with all the compiled files and pack them into a jar while making a bin folder for the jar to stay in. The Jar file will autorun so you can mess with the finished product and check for errors that may occur in the console/terminal.
	
	After the jar is created it will delete the srcOut unless you change the KeepSrcOut to true.
	
	## Experiments

	*new feature added in Jar Builder 1.1*
	
	` + "```" + `batch
	set "experiments=false"
	
	@REM Only works if experiments are enabled
	set "javaBaseDir="
	set "javaJdkVer=jdk-13"
	` + "```" + `
	Get ready and enable the experiments now we can build with any preinstalled jdk. To use this feature turn "experiments" from false to true and add a "javaBaseDir" which is just where your java "JDKs" are installed then what ever folder your jdk is in take the path and paste it in javaBaseDir and type whatever jdks you have installed in that folder. Java Base Example ` + "`" + `E:\Program Files\Java\` + "`" + `.`)

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
