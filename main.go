package main

import (
	"ProjectGen/gens"
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Supported Languages")
	fmt.Println("  { 1 } Python")
	fmt.Println("  { 2 } Java")
	fmt.Println("  { 3 } Go")
	fmt.Println("  { 4 } C++")
	fmt.Println("  { 5 } Website")
	fmt.Println("")
	fmt.Println("Choose A Language:")

	scanner.Scan()
	lang := scanner.Text()
	fmt.Println("Project Name:")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Project Version:")
	scanner.Scan()
	ver := scanner.Text()

	if lang == "1" {
		gens.GeneratePython(name, ver)
	}
	if lang == "2" {
		gens.GenerateJava(name, ver)
	}
	if lang == "3" {
		gens.GenerateGolang(name, ver)
	}
	if lang == "4" {
		gens.GenerateCpp(name, ver)
	}
	if lang == "5" {
		gens.GenerateWebsite(name, ver)
	}

}
