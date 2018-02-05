package main

import (
  "fmt"
  "flag"
  "os"
  "io/ioutil"
  "strings"
)

const VERSION = "0.0.1"

type TaskStore struct {
  conmigoDir string
  tasksFile string
}

func NewTaskStore() *TaskStore {
  dir := conmigoDirPath()
  file := conmigoTasksPath()

  ensureDirExists()
  ensureFileExists(file)
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

func ensureFileExists(filepath string) {
  if _, err := os.Stat(filepath); os.IsNotExist(err) {
    err2 := ioutil.WriteFile(filepath, []byte(""), 0744)
    if err2 != nil {
      panic(err2)
    }
  }
}

func (ts *TaskStore) AddTask(summary string) {
  newtid := maxTaskId(ts.readTasks()) + 1

  task := NewTask(summary, newtid)
  ts.writeTask(task)
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

func maxTaskId(tasks []Task) int {
  maxtid := 0
  for _, t := range (tasks){
    if t.tid > maxtid {
      maxtid = t.tid
    }
  }
  return maxtid
}

func cmdVersion(){
  fmt.Println("conmigo version", VERSION)

  os.Exit(0)
}

func cmdStartTask(ts *TaskStore, summary string){
  ts.AddTask(summary)
  fmt.Println("Task started:", summary)

  os.Exit(0)
}

func cmdListTasks(ts *TaskStore) {
  tasks := ts.readTasks()

  for _, task := range tasks {
    fmt.Println(task.String())
  }

  os.Exit(0)
}

func main(){
  versionPtr := flag.Bool("v", false, "Print version information and exit")
  startPtr := flag.Bool("start", false, "Start a new task")
  flag.Parse()

  ts := NewTaskStore()

  if *versionPtr {
    cmdVersion()
  } 
  if *startPtr {
    args := flag.Args()
    if len(args) != 1 {
      fmt.Println("expected one argument after -start. try using quotes")
      os.Exit(1)
    }
    cmdStartTask(ts, args[0])
  }

  // otherwise: list all open tasks
  cmdListTasks(ts)
}

