package main

import (
  "strconv"
  "strings"
)

const ID_NAME_SEP = " - "

type Task struct {
  tid int 
  name string
  open bool
}

func NewTask(name string, tid int) Task {
  return Task{name: name, tid: tid, open: true}
}

func (t Task) String() string {
  tidstr := strconv.Itoa(t.tid)
  closedstr := ""
  if !t.open  {
    closedstr = ID_NAME_SEP + "CLOSED"
  }
  return tidstr + closedstr + ID_NAME_SEP + t.name
}

func ValueOf(str string) Task {
  parts := strings.Split(str, ID_NAME_SEP)
  tid, err := strconv.Atoi(parts[0])
  if err != nil { panic(err) }
  if len(parts) > 2 {
    t := NewTask(parts[2], tid)
    if parts[1] == "CLOSED" {
      t.open = false
    }
    return t
  }
  return NewTask(parts[1], tid)
}
