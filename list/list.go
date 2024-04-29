package list

type MyList[T any] interface {
	AddFront(data T)
	AddLast(data T)
	Add(index int, data T) error
	RemoveFront() error
	RemoveLast() error
	Remove(index int) error
	Get(index int) (T, error)
	Length() int
	ToSlice() []T
	Iterator() Iterator[T]
}
