package main

import "log"

// the states of the state machine using constants (or an enum)
type State int

const (
	StateIdle State = iota
	StateRunning
	StatePaused
	StateStopped
)

// a struct that represents the state machine
type StateMachine struct {
	State State
	Value int
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		State: StateIdle, // the state machine is initialised in the idle state
		Value: 0,
	}
}

// methods on the state machine struct that handle state transitions and perform actions based on the existing state
func (sm *StateMachine) Start() {
	log.Printf("START | before state | %s", printState(sm.State))
	switch sm.State {
	case StateIdle, StateStopped:
		sm.State = StateRunning
		sm.Value = 0 // reset the value, revert to initial state, to maintain a consistent starting point
		log.Printf("state machine started | value initialized to: %d", sm.Value)
	case StateRunning:
		log.Printf("state machine is already running")
	case StatePaused:
		sm.State = StateRunning
		log.Printf("state machine resumed from paused state | value: %d", sm.Value)
	}
	log.Printf("START | after state | %s", printState(sm.State))
	log.Print("--------------------")
}

func (sm *StateMachine) Pause() {
	log.Printf("PAUSE | before state | %s", printState(sm.State))
	switch sm.State {
	case StateIdle:
		log.Printf("❌ cannot pause from the idle state")
	case StateRunning:
		sm.State = StatePaused
		log.Printf("state machine paused | value: %d", sm.Value)
	case StatePaused:
		log.Printf("state machine is already paused")
	case StateStopped:
		log.Printf("❌ cannot pause from the stopped state")
	}
	log.Printf("PAUSE | before state | %s", printState(sm.State))
	log.Print("--------------------")
}

func (sm *StateMachine) Stop() {
	log.Printf("STOP | before state | %s", printState(sm.State))
	switch sm.State {
	case StateIdle:
		log.Printf("state machine is already stopped.")
	case StateRunning, StatePaused:
		sm.State = StateStopped
		log.Printf("state machine stopped | value: %d", sm.Value)
	case StateStopped:
		log.Printf("state machine is already stopped")
	}
	log.Printf("STOP | after state | %s", printState(sm.State))
	log.Print("--------------------")
}

func (sm *StateMachine) ProcessValue() {
	log.Printf("PROCESS VALUE | attempting...")
	switch sm.State {
	case StateRunning:
		sm.Value++
		log.Printf("processed value | new value: %d", sm.Value)
	default:
		log.Printf("❌ cannot process value in the %s state", printState(sm.State))
	}
	log.Print("--------------------")
}

func main() {
	// create a new state machine using the constructor
	sm := NewStateMachine()

	// start the state machine (from an idle state)
	sm.Start()

	// process value while the state machine is running
	for i := 0; i < 3; i++ {
		sm.ProcessValue()
	}

	// pause the state machine
	sm.Pause()

	// cannot process the value while paused
	sm.ProcessValue()

	// restart the state machine from a paused state
	sm.Start()

	// process more values
	for i := 0; i < 6; i++ {
		sm.ProcessValue()
	}

	// stop the state machine
	sm.Stop()

	// cannot process value while stopped
	sm.ProcessValue()

	// restart the state machine from a stopped state
	sm.Start()

	// process more values
	for i := 0; i < 3; i++ {
		sm.ProcessValue()
	}
}
