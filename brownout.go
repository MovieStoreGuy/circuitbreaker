package circuitbreaker

type BrownOut interface {
  Switch

  CPULoadIsBelow(percentage float32) Switch
}
