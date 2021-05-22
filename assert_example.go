package main

import (
   "fmt"
   "github.com/jimlawless/assertjl"
)

func main()  {
   assertjl.Assert(1==1)
   fmt.Printf("Made it to the first output line.\n")
   assertjl.Assert(1==2)
   fmt.Printf("Made it to the second output line.\n")
}
