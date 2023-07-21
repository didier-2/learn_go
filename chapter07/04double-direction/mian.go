package main

import "fmt"

type LinkNode struct { //节点
	data     int
	next     *LinkNode
	previous *LinkNode
}

func buildDLink() *LinkNode { //构造一个双链表
	n1 := &LinkNode{data: 1}
	n2 := &LinkNode{data: 5}
	n3 := &LinkNode{data: 10}

	n1.next = n2
	n2.previous = n1

	n2.next = n3
	n3.previous = n2

	return n1
}

func insertNode(root *LinkNode, newNode *LinkNode) *LinkNode {
	tmpNode := root

	//整个链表是空的，新增节点
	if root == nil {
		return newNode
	}

	//在链表头部新增节点
	if root.data >= newNode.data {
		newNode.next = tmpNode
		tmpNode.previous = newNode
		return newNode
	}
	for {
		if tmpNode.next == nil {
			//已经到尾部 追加节点
			tmpNode.next = newNode
			newNode.previous = tmpNode
			return root
		} else {
			if tmpNode.next.data >= newNode.data {
				//找到位置，在此插入节点
				newNode.previous = tmpNode
				newNode.next = tmpNode.next

				tmpNode.next = newNode
				newNode.next.previous = newNode

				return root
			}
			tmpNode = tmpNode.next
		}
	}
}
func deleteNode(root *LinkNode, v int) *LinkNode {
	if root == nil {
		return nil
	}

	if root.data == v {
		//要删除的数据在第一个节点
		left := root
		root = root.next

		left.next = nil
		root.previous = nil

		//todo 需要解决只有一个节点的情况

		return root
	}

	tmpNode := root
	for {
		if tmpNode.next == nil {
			//链表走到尾部，仍然没有找到要删除的数据，直接返回root
			return root
		} else {
			if tmpNode.next.data == v {
				//找到节点，开始删除，删除后返回原root
				right := tmpNode.next
				tmpNode.next = right.next
				right.next.previous = tmpNode

				//清理右手上的link链接，保证gc正常回收
				right.next = nil
				right.previous = nil
				return root
			}
		}
		tmpNode = tmpNode.next
	}
}
func rangeLink(root *LinkNode) {
	fmt.Println("从头到尾")
	tmpNode := root
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.next == nil {
			//f(tmpNode.data)
			break
		}
		tmpNode = tmpNode.next
	}
	fmt.Println("从尾到头")
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.previous == nil {
			break
		}
		tmpNode = tmpNode.previous
	}

}
func main() {
	root := buildDLink()
	root = insertNode(root, &LinkNode{data: 3}) //需要重新对root进行赋值，刷新节点
	root = insertNode(root, &LinkNode{data: 7})
	root = insertNode(root, &LinkNode{data: 11})
	root = insertNode(root, &LinkNode{data: 0})
	root = insertNode(root, &LinkNode{data: -1})
	rangeLink(root)
	fmt.Println("删除节点")
	root = deleteNode(root, -1)
	root = deleteNode(root, 3)
	rangeLink(root)

}
