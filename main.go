package main

import (
	"fmt"

	"askar.khaimuldin/example/list"
	"askar.khaimuldin/example/list/linkedlist"
)

func main() {
	myList := linkedlist.New[string]()

	fillList(myList)
	err := myList.Add(3, "Paris")
	if err != nil {
		fmt.Printf("error while insertion at index: %v", err)

		return
	}

	err = myList.Remove(4)
	if err != nil {
		fmt.Printf("error while deletion at index: %v", err)

		return
	}

	err = myList.RemoveFront()
	if err != nil {
		fmt.Printf("error while deletion at front: %v", err)

		return
	}

	err = myList.RemoveLast()
	if err != nil {
		fmt.Printf("error while deletion at tail: %v", err)

		return
	}

	err = printList(myList)
	if err != nil {
		fmt.Printf("error while printing the list: %v", err)
	}
}

func fillList(myList list.MyList[string]) {
	myList.AddLast("Astana")
	myList.AddLast("Madrid")
	myList.AddLast("London")
	myList.AddFront("Almaty")
	myList.AddFront("New York")
}

func printList(myList list.MyList[string]) error {
	for it := myList.Iterator(); it.HasNext(); {
		data, err := it.Next()
		if err != nil {
			return err
		}

		fmt.Println(data)
	}

	return nil
}
