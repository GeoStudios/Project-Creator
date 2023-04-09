package gens

import (
	"os"
	"strings"
)

func GeneratePhp(name, ver string) {

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
	pf, _ := os.Create(name + ".php")
	pf.WriteString(`<!DOCTYPE html>
	<html>
	<body>
	
	<?php
	echo "My first PHP script!";
	?>
	
	</body>
	</html>`)
	pf.Close()

}
