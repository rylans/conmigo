package main

import "testing"

func TestTaskSerializeDeserialize(t *testing.T){
  t1 := NewTask("Clean up.", 1)
  t1string := t1.String()

  assertEqual(t, "1 - Clean up.", t1string)

  deserialized := ValueOf(t1string)

  if t1 != deserialized {
    t.Fail()
  }
}

func TestTaskSerializeDeserializeClosedTask(t *testing.T){
  t1 := NewTask("buy milk", 213)
  t1.open = false
  t1string := t1.String()

  assertEqual(t, "213 - CLOSED - buy milk", t1string)

  deserialized := ValueOf(t1string)

  if t1 != deserialized {
    t.Fail()
  }
}

func assertEqual(t *testing.T, expected, actual string) {
  if expected != actual {
    t.Fail()
  }
}
