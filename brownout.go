package circuitbreaker

// BrownOut is an extention to the Switch interface that looks at simplifying
// blocking routes based on runtime metrics
type BrownOut interface {
	Switch

	CPULoadIsBelow(percentage float32) Switch

	MemoryUsageIsBelow(count int) Switch
}
