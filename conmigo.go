package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "strings"
)

type TaskStore struct {
  conmigoDir string
  tasksFile string
}

func NewTaskStore() *TaskStore {
  dir := conmigoDirPath()
  file := conmigoTasksPath()

  ensureDirExists()
  return &TaskStore{conmigoDir: dir, tasksFile: file}
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

func (ts *TaskStore) writeTask(task Task) {
  path := ts.tasksFile
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

func (ts *TaskStore) readTasks() []Task {
  path := ts.tasksFile
  if _, err := os.Stat(path); os.IsNotExist(err) {
    panic (err)
  }

  dat, err := ioutil.ReadFile(path)
  if err != nil {
    panic(err) 
  }

  tasks := make([]Task, 0)
  for _, line := range strings.Split(string(dat), "\n") {
    if line == "" { 
      continue 
    }
    tasks = append(tasks, ValueOf(line))
  }

  return tasks
}

func main(){
  ts := NewTaskStore()

  t := NewTask("clean your room", 12)
  fmt.Println(t)

  ts.writeTask(t)

  fmt.Println(ts.readTasks())

}

