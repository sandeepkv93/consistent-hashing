package main

import "fmt"

func main() {
	consistentHash := NewConsistentHash()
	consistentHash.AddNode("node1")
	consistentHash.AddNode("node2")
	consistentHash.AddNode("node3")

	// Add some keys and test which node they map to
	keys := []string{"key1", "key2", "key3", "key4", "key5"}
	for _, key := range keys {
		fmt.Printf("Key '%s' maps to node: %s\n", key, consistentHash.GetNode(key))
	}

	// Remove a node and test again
	fmt.Println("\nRemoving node 'node2'...")
	consistentHash.RemoveNode("node2")
	for _, key := range keys {
		fmt.Printf("Key '%s' maps to node: %s\n", key, consistentHash.GetNode(key))
	}

	// Add a node and test again
	fmt.Println("\nAdding node 'node4'...")
	consistentHash.AddNode("node4")
	for _, key := range keys {
		fmt.Printf("Key '%s' maps to node: %s\n", key, consistentHash.GetNode(key))
	}
}
