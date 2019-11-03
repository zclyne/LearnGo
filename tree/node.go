package main

import "fmt"

// go中通过大小写来区分public或private
// 首字母大写则是public，小写则是private，这些是相对包而言的，若要访问另一个包中的某个方法，则这个方法必须是public的
// struct的方法不一定要定义在同一个文件中，但必须在同一个包中
// 导入其他的包要使用import语句，要使用相对导入，必须把要导入的包放在GOPATH/src下
// 使用其他包中的类型时，需要用包名.方法名或包名.struct名字的方法

// go中只有struct，没有class，只支持封装，而没有继承和多台
// go采用面向接口编程
type TreeNode struct {
	value int
	left, right *TreeNode
}

// 创建treeNode的工厂函数
func createNode(value int) *TreeNode {
	// 下面这种写法中，返回的地址指向一个局部变量，这在C++中会导致程序出错
	// 但是go语言中，这种写法是可行的，而且分配在堆上还是栈上是不确定的
	// 如果在函数内部创建的变量并没有返回指针到外部，则编译器可能把他分配到栈上，函数退出时自动销毁
	// 如果返回了指针给外部，则可能分配到堆上，并会参与垃圾回收
	return &TreeNode{value: value}
}

// 为结构体定义方法，并不写在结构体内部，而是写在外部，并且需要定义一个接收者
// 在函数名之前加入的内容表示这个print()方法是给node来接收的
// 传递方法同样是传值
// 这种写法其实和func Print(node treeNode)没有区别，只不过调用时可以用.来调用，而不用作为参数传入
func (node TreeNode) Print() {
	fmt.Println(node.value)
}

// 错误的改变结构体中的值的方式，因为是值传递，所以无效
func (node TreeNode) SetValue_wrong(value int) {
	node.value = value
}

// 正确的方式。注意内部仍然用node.value，go中只有.
func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	// 虽然在nil上可以调用方法，但是在执行下面这句设置value时还是会报错
	node.value = value
}

// 中序遍历一棵树
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	// 因为允许在nil上调用函数，所以这里不需要判断left和right是否为nil
	node.left.Traverse()
	node.Print()
	node.right.Traverse()
}

// 由于go中没有继承，若要扩展系统类型或其他包中的类型，需要使用别名或组合的方法
type myTreeNode struct {
	node *TreeNode
}

// 后续遍历
func (myNode *myTreeNode) postOrder() {
	// 这里还要判断一下包装的node是否为nil
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.left}
	left.postOrder()
	right := myTreeNode{myNode.node.right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	// 创建treenode
	var root TreeNode // value = 0, left = right = nil
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

	// 使用组合的方式扩展TreeNode
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
