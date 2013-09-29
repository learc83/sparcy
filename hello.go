package main

import (
	"fmt"
	"strings"
	"strconv"
)

var PC int;
var NPC int;
var jump int;
var lines []string;
var labels map[string]int;
var regs map[string]int = map[string]int{"%l0":0, "%l1":0, "%l2":0, "%l3":0,
										 "%o0":0, "%o1":0}

func main() {
	readProgram()

	PC = labels["main"]
	NPC = PC + 1

	running := true;

	for(running == true) {
		running = execute()

		PC = NPC
		if jump == 0 {
			NPC++	
		} else {
			NPC = jump
			jump = 0
		}
		
	}
}

func execute() (running bool) {
	running = true
	//line := lines[PC]
	//lineArray := strings.Split(line, " ")

	fmt.Println(lines[PC])
	fmt.Println(regs)

	command,params := parseInstruction(lines[PC])

	switch command {
	case "ta": 
		running = false
	case "mov":
		mov(params)
	case "add":
		add(params)
	case "sub":
		sub(params)
	case "ba":
		ba(params)
	case "nop":
		
	}

	return
}

func ba(params []string) {
	jump = labels[params[0]]
}

//TODO add check that command is formatted properly
func mov(params []string) {
	value := getValue(params[0])

	reg := strings.Trim(params[1], " ")

	regs[reg] = value 
}

func add(params []string) {
	val1 := getValue(params[0])

	val2 := getValue(params[1])

	reg := strings.Trim(params[2], " ")//TODO trim whitespace during parsing
	
	regs[reg] = val1 + val2
}

func sub(params []string) {
	val1 := getValue(params[0])

	val2 := getValue(params[1])

	reg := strings.Trim(params[2], " ")//TODO trim whitespace during parsing
	
	regs[reg] = val1 - val2
}

//returns an int, either directly or from the corresponding register
func getValue(param string) (val int) {
	if strings.Contains(param, "%") {
		reg := strings.Trim(param, " ")
		val  = regs[reg]
	} else {
		val2Str := strings.Trim(param, " ")
		val,_ = strconv.Atoi(val2Str)	
	}

	return	
}

func parseInstruction(instr string) (command string, params []string) {
	//split into 2 substrings. The whitespace seperated command in the first
	//and the rest of the instructions in the second
	instrArray := strings.SplitN(instr, " ", 2)
	command = instrArray[0]	

	//split second array into params with comma delimeter
	if len(instrArray) > 1 {
		params = strings.Split(instrArray[1], ",")
	}
	return
}

func readProgram() {
	lines = []string{"mov 4, %l0",
					 "add %l0, 1, %l0",
					 "add %l0, 2, %l0",
					 "sub %l0, 3, %l0", 
					 "ba loop", 
					 "nop", 
					 "ta 0"}
	labels = map[string]int{"main": 0, "loop": 1}
}

