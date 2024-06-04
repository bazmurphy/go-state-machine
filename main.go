package main

import "log"

// define the states of your state machine using an enum or constants

type State int

const (
	StateIdle State = iota
	StateRunning
	StatePaused
	StateStopped
)

// the iota keyword provides a convenient way to assign incrementing values
// to a set of related constants without explicitly specifying the values for each constant.
// it starts at 0 and increments by 1 for each subsequent constant declaration within the same const block.

// StateIdle     = 0
// StateRunning  = 1
// StatePaused   = 2
// StateStopped  = 3

// ----------

// define a struct that represents your state machine
// it should include the current state and any other relevant data

type StateMachine struct {
	currentState State
	// add other fields as needed
}

// ----------

// implement methods on the state machine struct
// to handle state transitions and perform actions based on the current state

func (sm *StateMachine) Start() {
	log.Println("Start() | BEFORE | sm.currentState:", sm.currentState)

	switch sm.currentState {
	case StateIdle:
		sm.currentState = StateRunning
		// perform actions for starting the state machine
	case StateRunning:
		// already running, do nothing
	case StatePaused:
		sm.currentState = StateRunning
		// resume from the paused state
	case StateStopped:
		// cannot start from the stopped state
	}

	log.Println("Start() | AFTER | sm.currentState:", sm.currentState)
}

func (sm *StateMachine) Pause() {
	log.Println("Pause() | BEFORE | sm.currentState:", sm.currentState)

	switch sm.currentState {
	case StateIdle:
		// cannot pause from the idle state
	case StateRunning:
		sm.currentState = StatePaused
		// perform actions for pausing the state machine
	case StatePaused:
		// already paused, do nothing
	case StateStopped:
		// cannot pause from the stopped state
	}

	log.Println("Pause() | AFTER | sm.currentState:", sm.currentState)
}

func (sm *StateMachine) Stop() {
	log.Println("Stop()  | BEFORE | sm.currentState:", sm.currentState)

	switch sm.currentState {
	case StateIdle:
		// already stopped, do nothing
	case StateRunning, StatePaused:
		sm.currentState = StateStopped
		// perform actions for stopping the state machine
	case StateStopped:
		// already stopped, do nothing
	}

	log.Println("Stop()  | AFTER | sm.currentState:", sm.currentState)
}

// ----------

// create an instance of the state machine
// and use its methods to control the state transitions

func main() {
	log.Println("Possible States:")
	log.Println("Idle:", StateIdle)
	log.Println("Running:", StateRunning)
	log.Println("Paused:", StatePaused)
	log.Println("Stopped:", StateStopped)

	sm := &StateMachine{currentState: StateIdle}

	sm.Start()
	// state machine is now running

	sm.Pause()
	// state machine is now paused

	sm.Start()
	// state machine is running again

	sm.Stop()
	// state machine is now stopped

	sm.Start()
	// but cannot start from the stopped state
}
