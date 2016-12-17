
package main



import (

        "os"
        "fmt"



 	"github.com/jonathb/stringutil"

)





func main() {

    wholecomlin := os.Args
    fmt.Println(wholecomlin)
    justargs := os.Args[1:]
    fmt.Println(justargs)
    numargs := len(os.Args)
    fmt.Println(numargs)

    arg := os.Args[3]
    fmt.Println(arg)

    fmt.Printf(stringutil.Reverse("hello, world\n") + "\n")
}


