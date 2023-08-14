package main

import (
  "fmt"
  "os"
)

var errorNoArgumentProvided = fmt.Errorf("Need to have at least 1 argument to run this program!")

// TODO: 0 - run program successfully!
// TODO: 1 - get argument from program!
// TODO: 2 - take 1st argument to read file stat!
// TODO: 3 - show last 10 character only!
// TODO: 4 - parameterize character count!
// TODO: 5 - follow stream if file is continuously writing!

// TODO: 0 - run program successfully!
func main() {
  fmt.Printf("%s is running!\n", os.Args[0])
  // TODO: 1 - get argument from program!
  argsWithoutProg := os.Args[1:]

  if len(argsWithoutProg) < 1 {
    exitWithError(errorNoArgumentProvided)
  }

  // TODO: 2 - take 1st argument to read file stat!
  file, err := os.Open(argsWithoutProg[0])
  if err != nil {
    exitWithError(err)
  }

  fileStat, err := file.Stat()
  if err != nil {
    exitWithError(err)
  }

  fmt.Printf("Name: %s \n", fileStat.Name())
}

func exitWithError(err error) {
    fmt.Fprintf(os.Stderr, "error: %v\n", err)
    os.Exit(1)
}
