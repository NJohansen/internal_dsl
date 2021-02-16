package metamodel

//Transition transitions
type Transition struct {
	name     string
	newState *State
	effect   func()
}

//NewTransition is constructor - Mangler Runnable effect!
func NewTransition(name string, newState *State, effect func()) *Transition {
	return &Transition{
		name:     name,
		newState: newState,
		effect:   effect,
	}
}

//GetEvent for getting transition event names
func (t *Transition) GetEvent() string {
	return t.name
}

//GetState for getting new transition states
func (t *Transition) GetState() *State {
	return t.newState
}

//Effect is for executing lambda function in an effect
func (t *Transition) Effect() {
	t.effect()
}
