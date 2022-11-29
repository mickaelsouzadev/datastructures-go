package datastructures

type Stack struct {
	count int
	items map[int]interface{}
}

func (stack *Stack) Push(element interface{}) {
	stack.items[stack.count] = element
	stack.count++
}

func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty() {
		stack.count--
		element := stack.items[stack.count]
		delete(stack.items, stack.count)

		return element
	}

	return nil
}

func (stack *Stack) Size() int {
	return stack.count
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Peek() interface{} {
	if !stack.IsEmpty() {
		i := stack.count - 1
		element := stack.items[i]
		return element
	}

	return nil
}

func InitStack() Stack {
	emptyitems := make(map[int]interface{})
	stack := Stack{count: 0, items: emptyitems}

	return stack
}
