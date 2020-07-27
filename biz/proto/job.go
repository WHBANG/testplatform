package proto

type JobPhase int32

const (
	PhaseUnset       JobPhase = 0
	PhaseCreated     JobPhase = 1
	PhasePending     JobPhase = 2
	PhaseRunning     JobPhase = 3
	PhaseFailed      JobPhase = 4
	PhaseTerminating JobPhase = 5
	PhaseDone        JobPhase = 6
)

var JobPhase_name = map[int32]string{
	0: "PhaseUnset",
	1: "PhaseCreated",
	2: "PhasePending",
	3: "PhaseRunning",
	4: "PhaseFailed",
	5: "PhaseTerminating",
	6: "PhaseDone",
}
