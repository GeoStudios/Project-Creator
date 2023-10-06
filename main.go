// Copyright (c) 2023 Geo-Studios - All Rights Reserved.

package main

import (
	"ProjectGen/gens"
	"bufio"
	"fmt"
	"os"
)

func main() {
	languages := map[string]func(string, string){
		"1":  gens.GenerateWebsite,
		"2":  gens.GeneratePython,
		"3":  gens.GenerateJava,
		"4":  gens.GenerateGolang,
		"5":  gens.GenerateC,
		"6":  gens.GenerateCpp,
		"7":  gens.GenerateCarbon,
		"8":  gens.GenerateRust,
		"9":  gens.GeneratePhp,
		"10": gens.GenerateNodejs,
		"11": gens.GenerateObjC,
		"12": gens.GenerateAsm,
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Supported Languages")
	for lang, _ := range languages {
		fmt.Printf("  { %s } %s\n", lang, gens.GetLanguageName(lang))
	}

	fmt.Println("\nChoose A Language:")
	scanner.Scan()
	lang := scanner.Text()

	if generator, exists := languages[lang]; exists {
		fmt.Println("Project Name:")
		scanner.Scan()
		name := scanner.Text()
		fmt.Println("Project Version:")
		scanner.Scan()
		ver := scanner.Text()
		generator(name, ver)
	} else {
		fmt.Println("Invalid language selection.")
	}
}
