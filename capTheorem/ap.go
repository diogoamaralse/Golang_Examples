package capTheorem

import (
	"fmt"
	"time"
)

//AP (Availability + Partition Tolerance)
//Availability and Partition Tolerance are guaranteed, but Consistency might be compromised during network partitions.
//This means that even if some nodes can't communicate, the system will still accept reads and writes, but they may not reflect the latest data.
//Example: AP System
//In this example, we simulate a scenario where the system continues to function during a partition, but we might get stale data due to the sacrifice of consistency.

func SimulateNetworkPartitionAP() {
	// Simulate network partition
	isPartition = true

	// Write operation (allowed even during partition)
	writeDataAP("SimulateNetworkPartitionAP: Data during Partition")

	// Read operation (could return stale data)
	fmt.Println("SimulateNetworkPartitionAP: Read during partition:", readDataCP())

	// Simulate removing the partition
	time.Sleep(2 * time.Second)
	isPartition = false

	// After partition removed, system operates normally
	writeDataAP("SimulateNetworkPartitionAP: Data after Partition")
	fmt.Println("SimulateNetworkPartitionAP: Read after partition:", readDataAP())
}

func writeDataAP(newData string) {
	if isPartition {
		// During partition, we still allow writes, but might lead to stale data
		data = newData
		fmt.Println("SimulateNetworkPartitionAP: Data written:", data)
		return
	}
	data = newData
	fmt.Println("SimulateNetworkPartitionAP: Data written:", data)
}

func readDataAP() string {
	return data // Data could be stale during partition
}
