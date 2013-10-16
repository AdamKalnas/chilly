package main

import "testing"

func TestTurnOnHeating(t *testing.T) {
  var therm = FakeThermometer{}
  var tempController = FakeTempController{} 
  therm.mode = "its_cold"
  ControlTemp(therm, &tempController)
  if (tempController.HeatIsOn != true) {
	t.Errorf("The heat isnt on and should be")
  }
}

func TestTurnOnCooling(t *testing.T) {
  var therm = FakeThermometer{}
  var tempController = FakeTempController{}
  therm.mode = "its_hot"
  ControlTemp(therm, &tempController)
  if (tempController.AcIsOn != true) {
	t.Errorf("The AC should be on")
  }
}



/* Fake Thermometor */
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

/* Fake Temp Controller */

type FakeTempController struct {
	HeatIsOn bool
	AcIsOn bool
}

func (t* FakeTempController) Cool() {
  t.AcIsOn = true
}

func (t* FakeTempController) Heat() {
  t.HeatIsOn = true
}
