package main

import "testing"

func TestHeating(t *testing.T) {
  var therm = FakeThermometer{}
  therm.mode = "its_cold"
  ControlTemp(therm)
}

type FakeThermometer struct {
  mode string
}

func (f FakeThermometer) GetTemp() int {
  switch f.mode {
  case "its_hot": return 80
  case "its_cold": return 10
  default: return 60
  }
}
