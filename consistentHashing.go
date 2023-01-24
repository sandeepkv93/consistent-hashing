package main

import (
	"crypto/md5"
	"sort"
	"strconv"
)

type ConsistentHash struct {
	nodes  map[string]string
	circle map[uint32]string
}

func NewConsistentHash() *ConsistentHash {
	return &ConsistentHash{
		nodes:  make(map[string]string),
		circle: make(map[uint32]string),
	}
}

func (c *ConsistentHash) AddNode(node string) {
	c.nodes[node] = node
	c.generateCircle()
}

func (c *ConsistentHash) RemoveNode(node string) {
	delete(c.nodes, node)
	c.generateCircle()
}

func (c *ConsistentHash) GetNode(key string) string {
	point := c.getPoint(key)
	var node string
	for k, v := range c.circle {
		if k > point {
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

func (c *ConsistentHash) getPoint(key string) uint32 {
	h := md5.New()
	h.Write([]byte(key))
	return uint32(h.Sum(nil)[3])<<24 + uint32(h.Sum(nil)[2])<<16 + uint32(h.Sum(nil)[1])<<8 + uint32(h.Sum(nil)[0])
}

func (c *ConsistentHash) generateCircle() {
	c.circle = make(map[uint32]string)
	for node := range c.nodes {
		for i := 0; i < 160; i++ {
			point := c.getPoint(node + ":" + strconv.Itoa(i))
			c.circle[point] = node
		}
	}
	var keys []int
	for k := range c.circle {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
}
