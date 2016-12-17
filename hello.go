package main

import (

        "os"
        "fmt"
 	"github.com/jonathb/gutil"
)

func doALL(a [] string) {
   fmt.Print("Got huge batch of ") ;fmt.Print(len(a)) ; fmt.Print(" args: ")
   fmt.Println(a)
   if len(a) >= 10 {panic("Dont grok so many args\n")}
}

func doX(as ... string) {
   fmt.Print("Got long list of ") ;fmt.Print(len(as)) ; fmt.Print(" args: ")
   stringmaker := ""
   for _, s := range as {
      stringmaker += s + " "
   }  // not sposed to be efficient alg, sposed to demo clever goisms 
   fmt.Println(stringmaker)
}

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

  wholeline := os.Args ;justargs := os.Args[1:] ;numargs := len(os.Args) -1
  fmt.Println(wholeline)
  fmt.Print(numargs) ; fmt.Print(" : ") ; fmt.Println(justargs)
  switch numargs {
  case 1:
  case 2: do1(os.Args[1])
  case 3: do2(os.Args[1],os.Args[2])
  case 4: do3(os.Args[1],os.Args[2],os.Args[3])
  case 5: doX(os.Args[1],os.Args[2],os.Args[3],os.Args[4])
  case 6: doX(os.Args[1],os.Args[2],os.Args[3],os.Args[4],os.Args[5])
  case 7: doX(os.Args[1],os.Args[2],os.Args[3],os.Args[4],os.Args[5],os.Args[6])
  default: doALL(os.Args[1:]) // isnt go slice brill?
  }
  fmt.Printf(gutil.Reverse("dlrow ,olleh") + "\n")
}


