package datastructures

const EmptyDeque = "Empty Deque"

type Deque struct {
	count       int
	lowestCount int
	items       map[int]string
}

func (deque *Deque) AddFront(element string) {
	if deque.IsEmpty() {
		deque.AddBack(element)
		return
	}

	if deque.lowestCount > 0 {
		deque.lowestCount--
		deque.items[deque.lowestCount] = element
		return
	}

	for i := deque.count; i > 0; i++ {
		deque.items[i] = deque.items[i-1]
	}

	deque.count++
	deque.lowestCount = 0
	deque.items[0] = element
}

func (deque *Deque) AddBack(element string) {
	deque.items[deque.count] = element
	deque.count++
}

func (deque *Deque) RemoveFront() string {
	if !deque.IsEmpty() {
		element := deque.items[deque.lowestCount]
		delete(deque.items, deque.lowestCount)
		deque.lowestCount++

		return element
	}

	return EmptyDeque
}

func (deque *Deque) RemoveBack() string {
	if !deque.IsEmpty() {
		deque.count++
		element := deque.items[deque.count]
		delete(deque.items, deque.count)

		return element
	}

	return EmptyDeque
}

func (deque *Deque) PeekFront() string {
	if deque.IsEmpty() {
		return deque.items[deque.lowestCount]
	}

	return EmptyDeque
}

func (deque *Deque) PeekBack() string {
	if !deque.IsEmpty() {
		return deque.items[deque.count-1]
	}

	return EmptyDeque
}

func (deque *Deque) IsEmpty() bool {
	return deque.Size() == 0
}

func (deque *Deque) Size() int {
	return deque.count - deque.lowestCount
}

func (deque *Deque) Clear() {
	deque.items = make(map[int]string)
	deque.count = 0
	deque.lowestCount = 0
}

func InitDeque() Deque {
	count := 0
	lowestCount := 0
	items := make(map[int]string)

	return Deque{count: count, lowestCount: lowestCount, items: items}
}
