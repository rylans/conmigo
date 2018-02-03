package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "strconv"
  "time"
)

type Task struct {
  name string
  createdAt time.Time
}

func NewTask(name string) Task {
  return Task{name: name, createdAt: time.Now()}
}

func (t Task) String() string {
  ts := strconv.FormatInt(t.createdAt.Unix(), 10)
  return ts + " " + t.name
}

func conmigoDirPath() string {
  return os.Getenv("HOME") + "/.conmigo"
}

func conmigoTasksPath() string {
  return conmigoDirPath() + "/.tasks.dat"
}

func ensureDirExists() {
  dirpath := conmigoDirPath()
  if _, err := os.Stat(dirpath); os.IsNotExist(err) {
    os.Mkdir(dirpath, 0744)
  }
}

func writeTask(task Task) {
  path := conmigoTasksPath()
  if _, err := os.Stat(path); os.IsNotExist(err) {
    err2 := ioutil.WriteFile(path, []byte(""), 0744)
    if err2 != nil {
      panic(err2)
    }
  }

  f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  if _, err := f.WriteString(task.String() + "\n"); err != nil {
    panic(err)
  }
}

func main(){
  ensureDirExists()

  t := NewTask("my task")
  fmt.Println(t)

  writeTask(t)

}

