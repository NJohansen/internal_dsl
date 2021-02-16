package metamodel

//State struct
type State struct {
	name        string
	transitions []*Transition
}

//NewState constructor
func NewState(name string) *State {
	return &State{
		name:        name,
		transitions: []*Transition{},
	}
}

//GetName for getting a statename
func (s *State) GetName() string {
	return s.name
}

//AddTransition for adding transitions to []Transition
func (s *State) AddTransition(transition *Transition) {
	s.transitions = append(s.transitions, transition)
}

//GetTransitions returns trans (s State) for binding the method on the struct
func (s *State) GetTransitions() []*Transition {
	return s.transitions
}
