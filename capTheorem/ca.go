package capTheorem

import (
	"fmt"
	"time"
)

//CA (Consistency + Availability)
//CA systems, Consistency and Availability are guaranteed, but Partition Tolerance is sacrificed.
//This means that the system will always be available for reads and writes and will maintain consistency unless a network partition occurs, at which point the system becomes unavailable.
//Example: CA System (Unavailable on Partition)
//Here we simulate a CA system where, during a partition, the system becomes unavailable.

func SimulateNetworkPartitionCA() {
	// Normal operation: no partition
	writeData("SimulateNetworkPartitionCA: Initial Data")
	fmt.Println("SimulateNetworkPartitionCA: Read data:", readData())

	// Simulate network partition (System becomes unavailable)
	isPartition = true
	writeData("SimulateNetworkPartitionCA: New Data during Partition") // Write fails
	readData()                                                         // Read fails

	// Simulate removing the partition
	time.Sleep(2 * time.Second)
	isPartition = false

	// After partition removed, we can write and read again
	writeData("SimulateNetworkPartitionCA: Data after Partition")
	fmt.Println("SimulateNetworkPartitionCA: Read after partition:", readData())
}

func writeData(newData string) {
	if isPartition {
		// During partition, we cannot perform writes
		fmt.Println("SimulateNetworkPartitionCA: Network partition detected. Cannot write data.")
		return
	}
	data = newData
	fmt.Println("SimulateNetworkPartitionCA : Data written:", data)
}

func readData() string {
	if isPartition {
		// During partition, we cannot read data
		fmt.Println("SimulateNetworkPartitionCA: Network partition detected. Cannot read data.")
		return ""
	}
	return data
}
