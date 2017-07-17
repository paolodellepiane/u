package term

import (
	"fmt"
	"math"
	"strings"
)

func ClearLine(width int) {
	fmt.Printf("\r" + strings.Repeat(" ", width) + "\n")
}

func OverwriteWithProgress(s string, progress float64, width int) {
	var progLength float64 = 20
	prog := math.Min(progress*progLength+1, progLength)
	p := fmt.Sprintf("%s%.2f%s ", "["+strings.Repeat("=", int(prog))+strings.Repeat(" ", int(progLength-prog))+"] ", progress*100, "%%")
	end := ""
	var fillerLength = int(width) - len(s) - len(p)
	if fillerLength > 0 {
		end = strings.Repeat(" ", fillerLength)
	}
	str := fmt.Sprintf("\r%s%s%s", p, s, end)
	if len(str) > int(width) {
		str = str[:width]
	}
	fmt.Printf(str)
}
