package main

import (

        "os"
        "fmt"
 	"github.com/jonathb/stringutil"
)

func do3(a1 string, a2 string, a3 string) {
   fmt.Println("Got 3 args")
   fmt.Println(a1,a2,a3)
}

func do2(a1 string, a2 string) {
   fmt.Println("Got 2 args")
   fmt.Println(a1,a2)
}

func do1(a1 string) {
   fmt.Println("Got 1 arg")
   fmt.Println(a1)
}


func main() {

    wholecomlin := os.Args
    fmt.Println(wholecomlin)
    justargs := os.Args[1:]
    fmt.Println(justargs)
    numargs := len(os.Args)
    fmt.Println(numargs)
    switch numargs {
    case 1:
    case 2: do1(os.Args[1])
    case 3: do2(os.Args[1],os.Args[2])
    case 4: do3(os.Args[1],os.Args[2],os.Args[3])
    default: panic("Dont grok more than 3 args\n")
    }
    fmt.Printf(stringutil.Reverse("dlrow ,olleh") + "\n")
}


