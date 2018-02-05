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

func (ts *TaskStore) AddTask(summary string) *Task {
  newtid := maxTaskId(ts.readTasks()) + 1

  task := NewTask(summary, newtid)
  ts.writeTask(task)
  return &task
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

func (ts *TaskStore) closeTask(taskid int) *Task {
  dat, err := ioutil.ReadFile(ts.tasksFile)
  if err != nil {
    panic(err) 
  }

  linesToWrite := make([]string, 0)
  lines := strings.Split(string(dat), "\n")

  foundTask := false
  var task Task
  for _, line := range lines {
    if line == "" {
      continue
    }
    thisTask := ValueOf(line)
    if thisTask.tid != taskid {
      linesToWrite = append(linesToWrite, line)
    } else {
      task = thisTask
      foundTask = true
    }
  }

  if !foundTask {
    fmt.Println("No task found with ID:", taskid)
    os.Exit(1)
  }

  writeThis := strings.Join(linesToWrite, "\n") + "\n"
  err = ioutil.WriteFile(ts.tasksFile, []byte(writeThis), 0744)
  if err != nil {
    panic(err)
  }

  closedtask := NewTask(task.name, task.tid)
  closedtask.open = false
  ts.writeTask(closedtask)

  return &closedtask
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
  task := ts.AddTask(summary)
  fmt.Println("Task started:")
  fmt.Println("\t", *task)

  os.Exit(0)
}

func cmdListTasks(ts *TaskStore) {
  tasks := ts.readTasks()

  for _, task := range tasks {
    if task.open {
      fmt.Println(task.String())
    }
  }

  os.Exit(0)
}

func cmdCloseTask(ts *TaskStore, taskid int) {
  task := ts.closeTask(taskid)
  fmt.Println("Task closed:")
  fmt.Println("\t", *task)

  os.Exit(0)
}

func main(){
  versionPtr := flag.Bool("v", false, "Print version information and exit")
  startPtr := flag.Bool("start", false, "Start a new task")
  endPtr := flag.Int("end", -1, "Close the task with the given ID")

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
  if *endPtr > 0 {
    cmdCloseTask(ts, *endPtr)
  }

  // otherwise: list all open tasks
  cmdListTasks(ts)
}

