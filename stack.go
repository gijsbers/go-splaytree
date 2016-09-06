package splaytree

// A stack to remember parent nodes when
// traversing a tree.
type stack struct {
	nodes []*node
}

func newStack() *stack {
	return &stack{nodes: make([]*node, 0, 10)}
}

func (stack *stack) push(node *node) {
	stack.nodes = append(stack.nodes, node)
}

func (stack *stack) length() int {
	return len(stack.nodes)
}

func (stack *stack) empty() bool {
	return len(stack.nodes) == 0
}

func (stack *stack) tos() *node {
	return stack.nodes[len(stack.nodes)-1]
}

func (stack *stack) pop() *node {
	node := stack.nodes[len(stack.nodes)-1]
	stack.nodes = stack.nodes[:len(stack.nodes)-1]
	return node
}
