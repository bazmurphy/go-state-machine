package main

func printState(currentState State) string {
	switch currentState {
	case StateIdle:
		return "0️⃣  Idle"
	case StateRunning:
		return "1️⃣  Running"
	case StatePaused:
		return "2️⃣  Paused"
	case StateStopped:
		return "3️⃣  Stopped"
	default:
		return ""
	}
}
