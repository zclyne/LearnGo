package main

import "fmt"



func main() {
	// 创建treenode
	var root tree.TreeNode // value = 0, left = right = nil
	fmt.Println(root)
	root1 := TreeNode{ // 创建value为3、left指向root的treenode
		value: 3,
		left:  &root,
		right: nil,
	}
	root1.right = &TreeNode{} // root1的right同样是一个value为0、left和right都是nil的treeNode，注意要加上&符号
	root.left = &TreeNode{5, nil, nil} // 按照顺序初始化
	root.left.right = new(TreeNode) // 使用内建函数new()，注意虽然root.left是一个指针，但是不像C++那样用->，而是用.
	root.left.left = createNode(2)

	// treeNode数组
	nodes := []TreeNode {
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	// 调用struct的方法
	root.Print() // 0
	root.SetValue_wrong(4)
	root.Print() // 0，因为是值传递
	root.SetValue(4)
	root.Print() // 4

	pRoot := &root
	pRoot.Print() // 4，虽然pRoot是指针，但是编译器会自动pRoot指向的root，然后执行相应的print
	pRoot.SetValue(5)
	pRoot.Print() // 5
	root.Print()  // 5
	// 在调用struct的方法时，编译器会自动进行指针和指针所指向的值的相应转换
	// nil指针也可以调用方法

	var pRoot1 *TreeNode
	pRoot1.SetValue(2)

	root.Traverse()

	// 值接收者vs指针接收者：
	// 1. 要改变内容必须使用指针接收者
	// 2. 结构过大也要考虑使用指针接收者
	// 3. 一致性（建议，非必须）：如有指针接收者，最好都是指针接收者
}
