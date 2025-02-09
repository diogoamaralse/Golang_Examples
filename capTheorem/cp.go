package capTheorem

import (
	"fmt"
	"time"
)

//CP (Consistency + Partition Tolerance)
//In CP systems, Consistency and Partition Tolerance are guaranteed, but Availability is sacrificed.
//If a partition occurs, the system might reject reads and writes to maintain consistency.
//Example: CP System
//In this example, we simulate a scenario where partition tolerance is maintained, but we are willing to sacrifice availability in favor of consistency.

func SimulateNetworkPartitionCP() {
	// Simulate network partition
	isPartition = true

	// Write operation (will be rejected due to partition)
	writeDataCP("SimulateNetworkPartitionCP: New Data")

	// Read operation (will also be rejected due to partition)
	readDataCP()

	// Simulate removing the partition
	time.Sleep(2 * time.Second)
	isPartition = false

	// After partition removed, we can write and read again
	writeDataCP("SimulateNetworkPartitionCP: New Data After Partition - ")
	fmt.Println("SimulateNetworkPartitionCP: Read after partition:", readDataCP())
}

func writeDataCP(newData string) {
	if isPartition {
		fmt.Println("SimulateNetworkPartitionCP: Network partition detected. Cannot write data.")
		return
	}
	data = newData
	fmt.Println("SimulateNetworkPartitionCP: Data written:", data)
}

func readDataCP() string {
	if isPartition {
		fmt.Println("SimulateNetworkPartitionCP: Network partition detected. Cannot read data.")
		return ""
	}
	return data
}
