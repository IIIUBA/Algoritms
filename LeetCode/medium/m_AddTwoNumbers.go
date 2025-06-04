package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return &ListNode{Val: 12}
}

func reverseNum(l1 *ListNode) {
	
	// либо перевернуть цепочку
	// либо перебрать всю цепочку
	// пересобрать можн выписав все числа от туда а потом создать новые поля с такими же числами только в другом порядке

	continius := true
	str := ""
	for continius {
		str += fmt.Sprint(l1.Val)
		str += fmt.Sprint(l1.Next.Val)

	}
	fmt.Println(str)

}

func main() {
	// Создаем первое число: 2 -> 4 -> 3 (342)
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	// Создаем второе число: 5 -> 6 -> 4 (465)
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	reverseNum(l1)

	result := addTwoNumbers(l1, l2)

	// Выводим результат
	for result != nil {
		print(result.Val, " ")
		result = result.Next
	}
}
