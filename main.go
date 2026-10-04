package main

import (
	"double_linked_list/types"
	"fmt"
)

func main(){

	dll := types.Dll{}

	appendDll(&dll, 1)
	appendDll(&dll, 2)
	appendDll(&dll, 3)
	appendDll(&dll, 4)
	appendDll(&dll, 5)
	appendDll(&dll, 6)
	appendDll(&dll, 7)
	appendDlls(&dll, 8, 9, 10)
	prependDll(&dll, 0)

	currentNode := dll.Head
	fmt.Println("Head node Previous <nil>")
	fmt.Println("currentNode:", currentNode)
	fmt.Println("next:", currentNode.Next.Data)
	fmt.Println("-----")
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		fmt.Println("Previous:", currentNode.Previous.Data)
		fmt.Println("currentNode:", currentNode)
		if  currentNode.Next != nil {
			fmt.Println("next:", currentNode.Next.Data)
		}
		fmt.Println("----- ")
	}
}

func appendDll(dll *types.Dll, data int) {

	if dll.Head == nil {
		head_node := types.Node{nil, data, nil}
		dll.Head = &head_node; return
	}

	currentNode := dll.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	next_node := types.Node{currentNode, data, nil}
	currentNode.Next = &next_node

}

func appendDlls(dll *types.Dll, data ...int) {

	if dll.Head == nil {
		head_node := types.Node{nil, data[0], nil}
		dll.Head = &head_node; return
	}

	currentNode := dll.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	for _, d := range data {
		next_node := &types.Node{currentNode, d, nil}
		currentNode.Next = next_node
		currentNode = next_node
	}

}

func prependDll(dll *types.Dll, data int) {

	if dll.Head == nil {
		head_node := &types.Node{nil, data, nil}
		dll.Head = head_node
	}

	previos_node := dll.Head
	newNode := &types.Node{nil, data, previos_node}
	previos_node.Previous = newNode
	dll.Head = newNode
}
