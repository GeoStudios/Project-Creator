// Copyright (c) 2023 Geo-Studios - All Rights Reserved.

package gens

import (
	"os"
	"strings"
)

func GenerateAsm(name, ver string) {

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
	pf, _ := os.Create(name + ".asm")
	pf.WriteString(`SECTION .data         

Msg: db "hello world", 10           
MsgLen: equ $ - Msg                


SECTION .text          

global start
start:

    push dword MsgLen  
    push dword Msg      
    push dword 1        
    sub esp, 4         
    mov eax, 4          
    int 80H             
    add esp, 16        

    push dword 0
    sub esp, 4
    mov eax, 1
    int 80H`)
	pf.Close()

}
