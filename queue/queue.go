package datastructures

const EmptyQueue = "Empty Queue"

type Queue struct {
	count       int
	lowestCount int
	items       map[int]string
}

func (queue *Queue) Enqueue(element string) {
	queue.items[queue.count] = element
	queue.count++
}

func (queue *Queue) Dequeue() string {
	if !queue.IsEmpty() {
		element := queue.items[queue.lowestCount]
		delete(queue.items, queue.lowestCount)
		queue.lowestCount++

		return element
	}

	return EmptyQueue
}

func (queue *Queue) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *Queue) Size() int {
	return queue.count - queue.lowestCount
}

func (queue *Queue) Clear() {
	queue.count = 0
	queue.lowestCount = 0
	queue.items = make(map[int]string)
}

func (queue *Queue) Peek() string {
	if !queue.IsEmpty() {
		return queue.items[queue.lowestCount]
	}

	return EmptyQueue
}

func InitQueue() Queue {
	count := 0
	lowestCount := 0
	items := make(map[int]string)

	return Queue{count: count, lowestCount: lowestCount, items: items}
}
