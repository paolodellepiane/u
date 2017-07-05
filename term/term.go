package term

import (
	"fmt"
	"math"
	"strings"
)

func OverwriteWithProgress(s string, progress float64, width int) {
	var progLength float64 = 20
	prog := math.Min(progress*progLength+1, progLength)
	p := fmt.Sprintf("%s%.2f%s ", "["+strings.Repeat("=", int(prog))+strings.Repeat(" ", int(progLength-prog))+"] ", progress*100, "%%")
	var fillerLength = int(width) - len(s) - len(p)
	end := ""
	if fillerLength == 0 {
		end = strings.Repeat(" ", fillerLength)
	}
	var str = "\r" + p + s + end
	if len(str) > int(width) {
		str = str[:width]
	}
	fmt.Printf(str)
}
