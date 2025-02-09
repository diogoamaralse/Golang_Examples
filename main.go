package main

import (
	"GoBestPratices/resilience"
)

// Structuring Large Go Projects
// Follow Standard Go Project Layout (pkg/, cmd/, internal/).
// Separate concerns using repository, service, handler patterns.
//├── cmd/        # Main entry points
//├── internal/   # Private application code
//├── pkg/        # Reusable packages
//├── api/        # API definitions
//├── config/     # Configuration files
//├── docs/       # Documentation

func main() {
	/*pratices.GoRoutine()
	pratices.RunChannel()
	pratices.ChannelGoRoutine()
	pratices.Context()
	pratices.ErrorHandling()
	pratices.Optimizing()
	pratices.Mutex()
	capTheorem.SimulateNetworkPartitionCP()
	capTheorem.SimulateNetworkPartitionAP()
	capTheorem.SimulateNetworkPartitionCA()
	consistency.TwoPhaseCommit()
	consistency.ProcessWithRedis()*/

	resilience.GoCircuitBreake()
}
