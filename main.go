package main

import (
	"fmt"

	"github.com/njohansen/internal_dsl/statemachine"
)

//The main function that executes my example
func main() {
	stateMachineExample()
}

//Function that shows the DSL
func stateMachineExample() {
	hue := statemachine.HUEsystem().State("OFF").Initial().
		When("turn on").
		To("ON", func() { fmt.Println("Lights is now ON") }).
		State("ON").
		When("+brightness").IsLessThan(100).
		To("ON", func() { fmt.Println("Brightness incremented with 10!") }).
		When("-brightness").IsMoreThan(0).
		To("ON", func() { fmt.Println("Brightness decreased with 10!") }).
		When("color cold").
		To("ON", func() { fmt.Println("Color is now in cold mode!") }).
		When("color blue").
		To("ON", func() { fmt.Println("Color is now in blue mode!") }).
		When("ZERO brightness").Equals(0).
		To("OFF", func() { fmt.Println("Light is now OFF") }).
		Build()

	hi := statemachine.NewHUEIntepreter(hue)
	hi.ProcessEvent("turn on")
	hi.ProcessEvent("+brightness")
	hi.ProcessEvent("-brightness")
	hi.ProcessEvent("color cold")
	hi.ProcessEvent("color blue")
	hi.ProcessEvent("ZERO brightness")
}
