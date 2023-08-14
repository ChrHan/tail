package main

import (
  "fmt"
  "os"
  "io"
  "strings"
)

var errorNoArgumentProvided = fmt.Errorf("Need to have at least 1 argument to run this program!")

// TODO: 0 - run program successfully!
// TODO: 1 - get argument from program!
// TODO: 2 - take 1st argument to read file stat!
// TODO: 3 - show last 10 character only!
// TODO: 4 - parameterize character count!
// TODO: 5 - follow stream if file is continuously writing!

// a way to learn golang by implementing `tail`
// general idea:
// 1. get file size
// 2. seek file from fixed cursor number (e.g. 10)
// 3. a. if file size is less than fixed cursor number, return whole fileStat
//    b. else, return only from cursor number to end of file
// TODO: 0 - run program successfully!
func main() {
  fmt.Printf("%s is running!\n", os.Args[0])
  // TODO: 1 - get argument from program!
  argsWithoutProg := os.Args[1:]

  if len(argsWithoutProg) < 1 {
    checkAndExitIfError(errorNoArgumentProvided)
  }

  // TODO: 2 - take 1st argument to read file stat!
  file, err := os.Open(argsWithoutProg[0])
  checkAndExitIfError(err)

  fileStat, err := file.Stat()
  checkAndExitIfError(err)
  fileSize := fileStat.Size()

  fmt.Printf("Name: %s \n", fileStat.Name())
  fmt.Printf("File Size: %d \n", fileSize)

  // TODO: 3 - show last 10 character only!
  // cursor for navigating characters
  var cursor int64 = -10
  // endless loop until all char is read
  for {
    cursor += 1
    // try reading the whole file to buffer
    _, err = file.Seek(cursor, io.SeekEnd)
    checkAndExitIfError(err)

    currentChar := make([]byte, 1)
    _, err := file.Read(currentChar)
    checkAndExitIfError(err)

    // printing every char found, not ideal - might be more efficient to stream it using other method 
    _, err = io.Copy(os.Stdout, strings.NewReader(string(currentChar)))
    checkAndExitIfError(err)

    if cursor == -1 {
      // stop if we reach the end of file!
      break
    }
  }
}

func checkAndExitIfError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "error: %v\n", err)
    os.Exit(1)
  }
}

