package main

import (
  "os"
  "io/ioutil"
)

const filemode = 0744

func ensureDirExists(dirpath string) {
  if _, err := os.Stat(dirpath); os.IsNotExist(err) {
    os.Mkdir(dirpath, filemode)
  }
}

func ensureFileExists(filepath string) {
  if _, err := os.Stat(filepath); os.IsNotExist(err) {
    err2 := ioutil.WriteFile(filepath, []byte(""), filemode)
    if err2 != nil {
      panic(err2)
    }
  }
}

func appendToFile(path string, line string) {
  f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, filemode)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  if _, err := f.WriteString(line + "\n"); err != nil {
    panic(err)
  }
}

func readFile(path string) (raw string) {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    panic (err)
  }

  dat, err := ioutil.ReadFile(path)
  if err != nil {
    panic(err) 
  }

  raw = string(dat)
  return
}

func replaceFileContents(path string, content string){
  err := ioutil.WriteFile(path, []byte(content), filemode)
  if err != nil {
    panic(err)
  }
}
