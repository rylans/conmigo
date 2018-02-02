package main

import (
  "fmt"
  "os"
  "io/ioutil"
)

func conmigoDirPath() string {
  return os.Getenv("HOME") + "/.conmigo"
}

func conmigoTasksPath() string {
  return conmigoDirPath() + "/tasks.dat"
}

func writeTask(bytes []byte) error {
  dirpath := conmigoDirPath()
  if _, err := os.Stat(dirpath); os.IsNotExist(err) {
    os.Mkdir(dirpath, 0744)
  }
  return ioutil.WriteFile(conmigoTasksPath(), bytes, 0744)
}

func main(){
  homeDir := os.Getenv("HOME")
  fmt.Println(homeDir)

  fmt.Println(writeTask([]byte("conmigo tasks")))
}

