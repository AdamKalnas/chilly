package main

import "fmt"

func main() {
  var therm = HallwayThermometer{}
  var t = therm.GetTemp()
  fmt.Println("Current tempature: ", t)

  var tempController = ResidentialTempController{}
  switch {
  case t > 70:
    tempController.Cool()
  case t < 50:
    tempController.Heat()
  }
}

/* Thermometer  */
type Thermometer interface {
  GetTemp() int
}

type HallwayThermometer struct {
}

func (h HallwayThermometer) GetTemp() int {
  return 40
}

/*   Temp controller */
type TempContoller interface {
  Cool()
  Heat()
}

type ResidentialTempController struct {
}

func (t ResidentialTempController) Cool() {
  fmt.Println("Turning on air conditioning")
}

func (t ResidentialTempController) Heat() {
  fmt.Println("Turning on heat")
}
