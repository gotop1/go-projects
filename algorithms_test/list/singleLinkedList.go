package main

import "algorithms/list"

func main() {
	var head *list.Node = list.Create(1)
	head.Append(2)
	head.Append(3)
	head.Append(4)

	head.Insert(3, 8)
	head.Insert(9, 63)

	head.Iterate()
}
