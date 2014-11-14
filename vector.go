package rabble

import (
	_ "fmt"
)

type INode interface {
	GetValue() interface{}
	SetValue(obj interface{})
}

type node struct {
	value interface{}
}

type IVector interface {
	GetNth(uint) interface{}
	SetNth(uint, interface{})
	Cons(interface{})
}

type vector struct {
	nodes []INode
}

func NewNode() node {
	return node{}
}

func NewVector() vector {
	return vector{}
}

func (n *node) GetValue() interface{} {
	return n.value
}

func (n *node) SetValue(obj interface{}) {
	n.value = obj
}

func (vec *vector) GetNth(i int) interface{} {
	var c int = len(vec.nodes)
	if c > i {
		return vec.nodes[i].GetValue()
	}
	return nil
}

func ensureNode(obj interface{}) INode {
	if node, ok := obj.(INode); ok {
		return node
	}
	return &node{obj}
}

func (vec *vector) Cons(obj interface{}) {
	node := ensureNode(obj)
	vec.nodes = append(vec.nodes, node)
}
