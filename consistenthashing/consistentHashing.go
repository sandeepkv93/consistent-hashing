package consistenthashing

import (
	"crypto/md5"
	"sort"
	"strconv"
)

type ConsistentHash struct {
	// map that stores the nodes in the system
	nodes map[string]string
	// map that stores the circle, where the keys are the generated hash points
	// and the values are the corresponding nodes
	circle map[uint32]string
}

// NewConsistentHash creates and returns a new instance of the ConsistentHash struct.
func NewConsistentHash() *ConsistentHash {
	return &ConsistentHash{
		nodes:  make(map[string]string),
		circle: make(map[uint32]string),
	}
}

// AddNode is used to add a new node to the consistent hash.
func (c *ConsistentHash) AddNode(node string) {
	c.nodes[node] = node
	c.generateCircle()
}

// RemoveNode is used to remove an existing node from the consistent hash.
func (c *ConsistentHash) RemoveNode(node string) {
	delete(c.nodes, node)
	c.generateCircle()
}

// GetNode returns the appropriate node for a given key.
func (c *ConsistentHash) GetNode(key string) string {
	point := c.getPoint(key)
	var node string
	for k, v := range c.circle {
		if k >= point {
			node = v
			break
		}
	}
	if node == "" {
		for _, v := range c.circle {
			node = v
			break
		}
	}
	return node
}

// getPoint uses the md5 package to generate a md5 hash of the key and
// return a 32-bit integer representation of the hash.
func (c *ConsistentHash) getPoint(key string) uint32 {
	h := md5.New()
	h.Write([]byte(key))
	return uint32(h.Sum(nil)[3])<<24 + uint32(h.Sum(nil)[2])<<16 + uint32(h.Sum(nil)[1])<<8 + uint32(h.Sum(nil)[0])
}

// generateCircle is called when nodes are added or removed to regenerate the circle.
func (c *ConsistentHash) generateCircle() {
	c.circle = make(map[uint32]string)
	for node := range c.nodes {
		for i := 0; i < 160; i++ {
			// generate 160 replicas for each node
			point := c.getPoint(node + ":" + strconv.Itoa(i))
			c.circle[point] = node
		}
	}
	// sort the circle keys in ascending order
	var keys []int
	for k := range c.circle {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
}
