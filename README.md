# Consistent Hashing

Consistent Hashing is a distributed hash table algorithm that allows for more efficient remapping of keys to nodes in a distributed system when nodes are added or removed. This implementation is written in Go and it provides a basic implementation of consistent hashing.

### Installation

To use this package, you can simply go get it:

```sh
go get github.com/sandeepkv93/consistent-hashing
```

### Usage

Here is an example of how to use the `ConsistentHash` struct:

```go
package main

import (
    "fmt"
    "github.com/yourusername/consistent-hashing"
)

func main() {
    consistentHash := consistent_hashing.NewConsistentHash()
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
```

In this example, the main method creates a new ConsistentHash object, adds three nodes to it, then adds some keys and tests which node they map to. Then it removes a node and tests again, and adds a node and tests again. This way you can see how the keys are distributed among the nodes and how the distribution changes when nodes are added or removed.

### Methods

`NewConsistentHash()` : This function creates and returns a new instance of the ConsistentHash struct.

`AddNode(node string)` : This method is used to add a new node to the consistent hash.

`RemoveNode(node string)` : This method is used to remove an existing node from the consistent hash.

`GetNode(key string)` : This method returns the appropriate node for a given key.

### Properties

`nodes` : A map that stores the nodes.

`circle` : A map that stores the circle, where the keys are the generated hash points and the values are the corresponding nodes.

### Note

You can adjust the number of replicas in the `generateCircle()` method to increase the number of points in the circle, which can help to distribute keys more evenly among the nodes.

### References

-   [Consistent Hashing - Wikipedia](https://en.wikipedia.org/wiki/Consistent_hashing)

### Additional Information

Consistent Hashing is a technique that is used to distribute keys across multiple nodes in a distributed system. It is particularly useful when the number of nodes in the system is subject to change, as it allows for the remapping of keys to nodes with minimal disruption to the overall system. This implementation is a basic implementation of consistent hashing and can be used as a starting point for more complex and sophisticated distributed systems.
