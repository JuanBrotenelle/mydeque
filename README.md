
# MyDeque

My double-ended queue. The capacity and length of the array are always equal. There is also a possibility of synchronous addition and deletion of elements.



## Where the information comes from

 - [Yandex Article. Deque](https://education.yandex.ru/handbook/algorithms/article/dek-(veque-double-ended-queue))
 - [Habr](https://habr.com/ru/articles/833444/)



## Methods

- `PushFront(value)`, `PopFront()` - add/remove element to/from front

- `PushBack(value)`, `PopBack()` - add/remove element to/from back

- `PushSync(value 1, value 2)`, `PopSync()` - add/remove elements to/from front or back at the same time

- `Front()`, `Back()` - return front or back elements

- `Len()`, `Cap()`, `isEmpty()`, `Clear()` - speaks for itself

- `Print()` - fmt.Println("Deque", deque)
