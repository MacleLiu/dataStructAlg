package main

import (
	linkList "dataStructAlg/linkList/singleLinkedList"
	"fmt"
)

func main() {
	head := linkList.NewLinkList()
	linkList.Insert(head, 5)
	fmt.Println(head.Next.Data)
}
