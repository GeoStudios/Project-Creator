package gens

import (
	"os"
	"strings"
)

func GenerateWebsite(name, ver string) {

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
	pf, _ := os.Create("index.html")
	pf.WriteString(strings.ReplaceAll(`<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<script src="script.js" defer></script>
		<link rel="stylesheet" href="style.css">
		
		<title>$Doc</title>
	</head>
	<body>
		
		<h1>Hello World</h1>

	</body>
	</html>`, "$Doc", name))
	pf.Close()
	ps, _ := os.Create("script.js")
	ps.WriteString(`function main(){
	console.log("Hello World")
}
main()`)
	ps.Close()
	pc, _ := os.Create("style.css")
	pc.WriteString(`body{background-color: black; color: white; display: flex; justify-content: center;}`)
	pc.Close()

}
