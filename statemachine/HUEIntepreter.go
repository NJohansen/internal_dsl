package statemachine

import (
	"fmt"

	"github.com/njohansen/internal_dsl/metamodel"
)

//HUEIntepreter struct with pointer to HUEsystem and State
type HUEIntepreter struct {
	hueSystem    *metamodel.HUEsystem
	currentState *metamodel.State
}

//NewHUEIntepreter constructor where i set currentState to initstate
func NewHUEIntepreter(hueSystem *metamodel.HUEsystem) HUEIntepreter {
	return HUEIntepreter{
		hueSystem:    hueSystem,
		currentState: hueSystem.GetInitState(),
	}
}

//Reset initState if needed
func (i *HUEIntepreter) Reset() {
	i.currentState = i.hueSystem.GetInitState()
}

//ProcessEvent for processing all events and Println if i have an unhandled event
func (i *HUEIntepreter) ProcessEvent(event string) {
	for _, transition := range i.currentState.GetTransitions() {
		if transition.GetEvent() == event {
			transition.Effect()
			i.currentState = transition.GetState()
			return
		}
	}
	fmt.Println("Unhandled event" + event)
}
