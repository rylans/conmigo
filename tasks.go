package main

import (
  "strconv"
  "strings"
)

const ID_NAME_SEP = " - "

type Task struct {
  tid int 
  name string
}

func NewTask(name string, tid int) Task {
  return Task{name: name, tid: tid}
}

func (t Task) String() string {
  tidstr := strconv.Itoa(t.tid)
  return tidstr + ID_NAME_SEP + t.name
}

func ValueOf(str string) Task {
  parts := strings.Split(str, ID_NAME_SEP)
  tid, err := strconv.Atoi(parts[0])
  if err != nil { panic(err) }
  return NewTask(parts[1], tid)
}
