package chimps_log

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/chimpss/go-utils/chimps-math"
)

const (
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
)

// PrintlnColorful PrintlnColorful
func PrintlnColorful(color string, vals ...interface{}) {
	fmt.Printf(color)
	fmt.Println(vals...)
	fmt.Printf("\033[0m")
}

// PrintlnDefault PrintlnDefault
func PrintlnDefault(vals ...interface{}) {
	fmt.Printf(Green)
	fmt.Println(vals...)
	fmt.Printf("\033[0m")
}

// PrintfColorful PrintfColorful
func PrintfColorful(color string, format string, vals ...interface{}) {
	fmt.Printf(color)
	fmt.Printf(format, vals...)
	fmt.Printf("\033[0m")
}

// P P
func P(vals ...interface{}) {
	PrintlnColorful(Green, vals...)
}

// PrintfDefault PrintfDefault
func PrintfDefault(format string, vals ...interface{}) {
	fmt.Printf(Yellow)
	fmt.Printf(format, vals...)
	fmt.Printf("\033[0m")
}

func Info(vals ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	names := strings.Split(file, "/")
	fmt.Printf("[%s:%d]  ", names[chimps_math.Max(0, len(names)-1)], line)
	PrintlnColorful(Yellow, vals...)
}

func Error(vals ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("[%s:%d]  ", file, line)
	PrintlnColorful(Red, vals...)
}

func Check(label string, err error) {
	if err != nil {
		Error(label, " failed,", err)
		os.Exit(1)
	}
	Info(label, " ok")
}
