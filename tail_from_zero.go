package main

import (
  "fmt"
  "os"
)

// TODO: 0 - run program successfully!
// TODO: 1 - get argument from program!
// TODO: 2 - take 1st argument to read file stat!
// TODO: 3 - show last 10 character only!
// TODO: 4 - parameterize character count!
// TODO: 5 - follow stream if file is continuously writing!

func main() {
  fmt.Printf("%s is running!\n", os.Args[0])
  argsWithoutProg := os.Args[1:]

  fmt.Println(argsWithoutProg)
}


