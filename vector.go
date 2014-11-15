package rabble

const (
	bits      = 5
	arraysize = (1 << bits)     // 32
	arraymask = (arraysize - 1) // 31
)

type node struct {
	array [arraysize]interface{}
}

type IVector interface {
	GetNth(int) (*Object, bool)
	SetNth(int, *Object) (IVector, bool)
	Cons(*Object) IVector
	Count() int
	RootArray() [arraysize]interface{}
}

type vector struct {
	root  *node
	tail  [arraysize]interface{}
	shift uint
	count int
}

func newNode() *node {
	n := new(node)
	return n
}

func NewVector() IVector {
	return &vector{root: newNode(), shift: bits, count: 0}
}

func (vec *vector) Count() int {
	return vec.count
}

func (vec *vector) RootArray() [arraysize]interface{} {
	return vec.root.array
}

func (vec *vector) GetNth(i int) (*Object, bool) {
	if vec.count > i {
		a, ok := vec.arrayFor(i)
		if !ok {
			return nil, false
		}
		n := a[i&arraymask]
		if val, ok := n.(*Object); ok {
			return val, true
		}
		return nil, false
	}
	return nil, false
}

func doAssoc(level uint, n *node, i int, obj *Object) (newnode *node) {
	newnode = newNode()
	newnode.array = n.array
	if level == 0 {
		newnode.array[i&arraymask] = obj
	} else {
		subidx := (i >> level) & arraymask
		subobj := n.array[subidx]
		subnode, _ := subobj.(*node)
		newnode.array[subidx] = doAssoc(level-bits, subnode, i, obj)
	}
	return
}

func (vec *vector) SetNth(i int, obj *Object) (IVector, bool) {
	if vec.count == i {
		return vec.Cons(obj), true
	}
	if vec.count > i && i >= 0 {
		if i >= vec.tailoff() {
			newTail := vec.tail
			newTail[i&arraymask] = obj
			return &vector{root: vec.root, shift: vec.shift, tail: newTail, count: vec.count}, true
		}
		newRoot := doAssoc(vec.shift, vec.root, i, obj)
		return &vector{root: newRoot, shift: vec.shift, tail: vec.tail, count: vec.count}, true
	}
	return nil, false
}

func (vec *vector) Cons(obj *Object) IVector {
	i := vec.count
	// room in tail?
	if (i - vec.tailoff()) < arraysize {
		newTail := vec.tail
		newTail[i&arraymask] = obj
		return &vector{root: vec.root, shift: vec.shift, tail: newTail, count: i + 1}
	}

	// full tail, push into tree
	var newroot *node
	tailNode := &node{array: vec.tail}
	newShift := vec.shift

	//overflow root?
	if (vec.count >> bits) > (1 << vec.shift) {
		newroot = newNode()
		newroot.array[0] = vec.root
		newroot.array[1] = newPath(vec.shift, tailNode)
		newShift += bits
	} else {
		newroot = pushTail(vec.count, vec.shift, vec.root, tailNode)
	}

	v := new(vector)
	v.root = newroot
	v.count = vec.count + 1
	v.shift = newShift
	v.tail[0] = obj

	return v
}

func pushTail(cnt int, level uint, parent *node, tail *node) *node {
	//if parent is leaf, insert node,
	// else does it map to an existing child? -> nodeToInsert = pushNode one more level
	// else alloc new path
	//return  nodeToInsert placed in copy of parent
	subidx := ((cnt - 1) >> level) & arraymask
	ret := newNode()
	ret.array = parent.array

	var nodeToInsert *node

	if level == bits {
		nodeToInsert = tail
	} else {
		child := parent.array[subidx]
		val, ok := child.(*node)
		if child == nil || !ok {
			nodeToInsert = newPath(level-bits, tail)
		} else {
			nodeToInsert = pushTail(cnt, level-bits, val, tail)
		}
	}
	ret.array[subidx] = nodeToInsert
	return ret
}

func newPath(level uint, n *node) *node {
	if level == 0 {
		return n
	}
	ret := newNode()
	ret.array[0] = newPath(level-bits, n)
	return ret
}

func (vec *vector) tailoff() int {
	cnt := vec.count
	if cnt < arraysize {
		return 0
	}
	return ((cnt - 1) >> bits) << bits
}

func (vec *vector) arrayFor(i int) ([arraysize]interface{}, bool) {
	if i >= 0 && i < vec.count {
		if i >= vec.tailoff() {
			return vec.tail, true
		}
		n := vec.root
		for level := vec.shift; level > 0; level -= bits {
			obj := n.array[(i>>level)&arraymask]
			val, ok := obj.(*node)
			if !ok {
				return [arraysize]interface{}{}, false
			}
			n = val
		}
		return n.array, true
	}
	return newNode().array, false
}
