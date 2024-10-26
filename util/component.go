package util

import (
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

func ViewTitle(title string) {
	gola.ClearScreen()
	fmt.Printf("\n%s\n\n", gola.Tf(gola.Bold, fmt.Sprintf("%s %s %s", "===", title, "==="), gola.LightBlue))
}
