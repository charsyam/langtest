package avl

type Node struct {
    left, right, parent *Node
    value, level int
}

func (o *Node) Init(v int) *Node {
    o.left = o.right = o.parent = nil
    o.value = v
    o.level = 0

    return o
}

func (o *Node) GetLeft() *Node {
    return o.left
}

func (o *Node) GetRight() *Node {
    return o.right
}

func (o *Node) GetParent() *Node {
    return o.parent
}

func (o *Node) GetValue() int {
    return o.value
}

func (o *Node) GetLevel() int {
    return o.level
}

type AVLTree struct {
    root *Node

}

func New() *AVLTree {
    return new(AVLTree).Init()
}

func (o *AVLTree) Init() *AVLTree {
    o.root = nil
    return o
}

func (o *AVLTree) Find(v int) *Node {
    return o._FindNode(o.root, v)
}

func (o *AVLTree) _FindNode(node *Node, v int) *Node {
    if node == nil {
        return nil
    }

    if node.GetValue() == v {
        return node
    }

    if node.GetValue() < v {
        return o._FindNode(node.GetRight(), v)
    } else {
        return o._FindNode(node.GetLeft(), v)
    }
}

func (o *AVLTree) Add(v int) {
    newNode := new(Node).Init(v)
    if o.root == nil {
        o.root = newNode
    } else {
        o.SimpleAddNode(o.root, newNode)
    }
}

func (o *AVLTree) SimpleAddNode(parent *Node, newNode *Node) {
    if parent.GetValue() <= newNode.GetValue() {
        if parent.GetRight() == nil {
            parent.SetRight(newNode) 
                newNode.SetParent(parent)
        } else {
            o.AddNode(parent.GetRight(), newNode)
        } 
    } else {
        if parent.GetLeft() == nil {
            parent.SetLeft(newNode) 
            newNode.SetParent(parent)
        } else {
            o.AddNode(parent.GetLeft(), newNode)
        } 
    }
}
