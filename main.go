package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

const (
	Cols = 100
	Lines = 42 // исправить на 40 и поправить в коде, костыль в 42 сделан для непостоянной генерации
)

var console_lines string = "lines=" + strconv.Itoa(Lines)
var console_cols string = "cols=" + strconv.Itoa(Cols)

func main() {
	//генерация поля в консоли в заданных параметрах 
	cmd := exec.Command("cmd", "/c", "mode", "con", console_cols, console_lines)
	cmd.Run()
	gamemap := make([]string, 0, Cols*(Lines - 2))
	for x := 0; x < Cols*(Lines-2); x++{
		switch {
		case x < Cols:
			gamemap = append(gamemap, "1")
		case x >= Cols*(Lines-3):
			gamemap = append(gamemap, "2")
		case x%Cols == 0:
			gamemap = append(gamemap, "3")
		case x%Cols == 99:
			gamemap = append(gamemap, "4")
		default:
			gamemap = append(gamemap, " ")
		}
	}
		

	for _, x := range gamemap{
		fmt.Printf(x)
	}
}