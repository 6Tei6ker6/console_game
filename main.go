package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

const (
	Cols = 200
	Lines = 62 // исправить на ровное число
)

var console_lines string = "lines=" + strconv.Itoa(Lines)
var console_cols string = "cols=" + strconv.Itoa(Cols)

func main() {
	//генерация поля в консоли в заданных параметрах 
	cmd := exec.Command("cmd", "/c", "mode", "con", console_cols, console_lines)
	cmd.Run()
	gamemap := make([]string, 0, Cols*(Lines - 2))
	for x := 0; x < Cols*(Lines-2); x++{ // исправить на Lines - 1
		switch {
		case x < Cols || x%Cols == (Cols-1) || x%Cols == 0 || x >= Cols*(Lines-3): // исправить на Lines - 1
			gamemap = append(gamemap, "@")
		case x == (Lines*Cols/2) + (Cols/2):
			gamemap = append(gamemap, "1")
		default:
			gamemap = append(gamemap, " ")
		}
	}
		

	for _, x := range gamemap{
		fmt.Printf(x)
	}
}