package deque

import "fmt"

type Deque struct {
	data []interface{}
}

// Create a deque. Here u can add any values, because the data type is empty interface
func NewDeque() *Deque {
	return &Deque{
		data: make([]interface{}, 0),
	}
}

// Add new value to the front
func (q *Deque) PushFront(v interface{}) {
	var newData = make([]interface{}, len(q.data)+1)

	newData[0] = v

	copy(newData[1:], q.data)

	q.data = newData
}

// Add new value to the back
func (q *Deque) PushBack(v interface{}) {
	var newData = make([]interface{}, len(q.data)+1)

	copy(newData, q.data)

	newData[len(q.data)] = v

	q.data = newData
}

// Remove the front value
func (q *Deque) PopFront() interface{} {
	var newData = make([]interface{}, len(q.data)-1)
	value := q.data[0]

	copy(newData, q.data[1:])

	q.data = newData

	return value
}

// Remove the back value
func (q *Deque) PopBack() interface{} {
	var newData = make([]interface{}, len(q.data)-1)
	value := q.data[len(q.data)-1]

	copy(newData, q.data[:len(q.data)-1])

	q.data = newData

	return value
}

// Push to front and back pos of the deque at the same time
func (q *Deque) PushSync(v1, v2 interface{}) {
	var newData = make([]interface{}, len(q.data)+2)

	newData[0] = v1

	if len(q.data) != 0 {
		copy(newData[1:], q.data)
	}

	newData[len(q.data)+1] = v2

	q.data = newData
}

// Remove and return front and back values at the same time
func (q *Deque) PopSync() (interface{}, interface{}) {
	var newData = make([]interface{}, len(q.data)-2)

	copy(newData, q.data[1:len(q.data)])

	v1, v2 := q.data[0], q.data[len(q.data)-1]

	q.data = newData

	return v1, v2
}

// Return front value
func (q *Deque) Front() interface{} {
	result := q.data[0]

	return result
}

// Return back value
func (q *Deque) Back() interface{} {
	result := q.data[len(q.data)-1]

	return result
}

// Return length of the deque
func (q *Deque) Len() int {
	return len(q.data)
}

// Return capacity of the deque
func (q *Deque) Cap() int {
	return cap(q.data)
}

// Return the state of emptiness
func (q *Deque) IsEmpty() bool {
	return len(q.data) == 0
}

// To clear the current deque
func (q *Deque) Clear() {
	q.data = make([]interface{}, 0)
}

// fmt.Print("Deque", deque). Print the values of the deque
func (q *Deque) Print() {
	fmt.Println("Deque:", q.data)
}
