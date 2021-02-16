package statemachine

import (
	"github.com/njohansen/internal_dsl/metamodel"
)

//HUEsystemBuilder struct with pointers
type HUEsystemBuilder struct {
	states          map[string]*metamodel.State
	initialState    *metamodel.State
	currentState    *metamodel.State
	currentEvent    string
	brightnessValue int16
}

//GetState is for getting the current state or add it to the slice h.states if it doesn't exist
func (h *HUEsystemBuilder) GetState(name string) *metamodel.State {
	value, ok := h.states[name]
	if ok == false {
		h.states[name] = metamodel.NewState(name)
		return h.states[name]
	}
	return value
}

//State to set a currentState
func (h *HUEsystemBuilder) State(name string) *HUEsystemBuilder {
	h.currentState = h.GetState(name)
	return h
}

//Initial for setting initialState to currentState
func (h *HUEsystemBuilder) Initial() *HUEsystemBuilder {
	h.initialState = h.currentState
	return h
}

//When is for the HUE DSL, so when() "some event happes" and set that event to current event
func (h *HUEsystemBuilder) When(event string) *HUEsystemBuilder {
	h.currentEvent = event
	return h
}

//To is for the HUE DSL, so when("some event").to("Set the state") and do some func().
// I only println() atm. but those could easily be changed.
func (h *HUEsystemBuilder) To(state string, effect func()) *HUEsystemBuilder {
	transition := metamodel.NewTransition(h.currentEvent, h.GetState(state), effect)
	h.currentState.AddTransition(transition)
	return h
}

//IsLessThan is just a dummy method, that should check a brightnessvalue
func (h *HUEsystemBuilder) IsLessThan(value int16) *HUEsystemBuilder {
	h.brightnessValue = value
	return h
}

//IsMoreThan is just a dummy method, that should check a brightnessvalue
func (h *HUEsystemBuilder) IsMoreThan(value int16) *HUEsystemBuilder {
	h.brightnessValue = value
	return h
}

//Equals is just a dummy method, that should check a brightnessvalue
func (h *HUEsystemBuilder) Equals(value int16) *HUEsystemBuilder {
	h.brightnessValue = value
	return h
}

//mapToSlice is a method for getting values in states from the map and add values to a slice
func (h *HUEsystemBuilder) mapToSlice(dict map[string]*metamodel.State) []*metamodel.State {
	m := make([]*metamodel.State, 0, len(h.states))
	for _, val := range h.states {
		m = append(m, val)
	}
	return m
}

//Build is from the builder pattern and for building the DSL
func (h *HUEsystemBuilder) Build() *metamodel.HUEsystem {
	return metamodel.NewHUEsystem(h.mapToSlice(h.states), *h.initialState)
}

//HUEsystem init. the HUEsystemBuilder
func HUEsystem() *HUEsystemBuilder {
	return &HUEsystemBuilder{
		//Initialize so map isn't nil
		states: make(map[string]*metamodel.State),
	}
}
