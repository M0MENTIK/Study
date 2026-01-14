package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Inpun command: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("error input")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text) //Input string, output slice[strung]

		if len(fields) == 0 {
			println("You nothing write")
			return
		}

		//fmt.Println("text:", text)

		pp.Println("Words:", fields)

		cmd := fields[0]

		if cmd == "out" {
			println("You exit functions. Buy!")
			return
		}

		if cmd == "add" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i != len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("You want to add:", str)

		} else if cmd == "delete" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i != len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("YOu want to delete:", str)
		} else if cmd == "help" {
			println("-- Command: help (List know commands)")
			fmt.Println("-- Command: add (What need add)")
			println("-- Command: delete (what need delete)")
		} else {
			println("You write unknown command")
		}

	}

}
