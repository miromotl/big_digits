// bigdigits.go

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

const usage = "usage: %s [option] <whole number>\n\n" +
              "options:\n" +
              "  -b, --bars: print bars of stars\n" +
              "  -h, --help: display this message\n\n"

func main() {
    bars := false
    stringOfDigits := ""
    
    if len(os.Args) == 1 {
        // print usage information
        fmt.Printf(usage, filepath.Base(os.Args[0]))
        os.Exit(1)
    } else {
        if os.Args[1] == "-h" || os.Args[1] == "--help" {
            fmt.Printf(usage, filepath.Base(os.Args[0]))
            os.Exit(1)
        } else if os.Args[1] == "-b" || os.Args[1] == "--bars" {            
            bars = true
            stringOfDigits = os.Args[2]
        } else {
            stringOfDigits = os.Args[1]
        }
    }

    // Create the bar string
    bar := ""
    for column := range stringOfDigits {
        digit := stringOfDigits[column] - '0'
        if 0 <= digit && digit <= 9 {
            bar += strings.Repeat("*", len(bigDigits[digit][0])+1)
        } else {
            log.Fatal("invalid whole number")
        }
    }
    bar = bar[:len(bar)-1]

    if bars {
        fmt.Println(bar)
    }
    
    for row := range bigDigits[0] {
        line := ""
        for column := range stringOfDigits {
            digit := stringOfDigits[column] - '0'
            if 0 <= digit && digit <= 9 {
                line += bigDigits[digit][row] + " "
            } else {
                log.Fatal("invalid whole number")
            }
        }
        fmt.Println(line)
    }
    
    if bars {
        fmt.Println(bar)
    }
}

var bigDigits = [][]string {
    {
        "  000  ",
        " 0   0 ",
        "0     0",
        "0     0",
        "0     0",
        " 0   0 ",
        "  000  ",
    },
    {
        " 1 ", 
        "11 ", 
        " 1 ", 
        " 1 ", 
        " 1 ", 
        " 1 ", 
        "111",
    },
    {
        " 222 ", 
        "2   2", 
        "   2 ", 
        "  2  ", 
        " 2   ", 
        "2    ", 
        "22222",
    },
    {
        " 333 ", 
        "3   3", 
        "    3", 
        "  33 ", 
        "    3", 
        "3   3", 
        " 333 ",
    },
    {
        "   4  ", 
        "  44  ", 
        " 4 4  ", 
        "4  4  ", 
        "444444", 
        "   4  ",
        "   4  ",
    },
    {
        "55555", 
        "5    ", 
        "5    ", 
        " 555 ", 
        "    5", 
        "5   5", 
        " 555 ",
    },
    {
        " 666 ", 
        "6    ", 
        "6    ", 
        "6666 ", 
        "6   6", 
        "6   6", 
        " 666 ",
    },
    {
        "77777", 
        "    7", 
        "   7 ", 
        "  7  ", 
        " 7   ", 
        "7    ", 
        "7    ",
    },
    {
        " 888 ", 
        "8   8", 
        "8   8", 
        " 888 ", 
        "8   8", 
        "8   8", 
        " 888 ",
    },
    {
        " 9999", 
        "9   9", 
        "9   9", 
        " 9999", 
        "    9", 
        "    9", 
        "    9",
    },
}
