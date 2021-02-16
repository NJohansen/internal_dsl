package metamodel

//HUEsystem struct
type HUEsystem struct {
	initState *State
	states    []*State
}

//NewHUEsystem constructor
func NewHUEsystem(states []*State, initState State) *HUEsystem {
	return &HUEsystem{
		initState: &initState,
		states:    append(states, states...),
	}
}

//GetStates for getting all states for the HUEsystem
func (h HUEsystem) GetStates() []*State {
	return h.states
}

//GetInitState for initState for HUEsystem
func (h HUEsystem) GetInitState() *State {
	return h.initState
}
