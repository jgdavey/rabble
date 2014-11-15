package rabble

type Object struct {
	value interface{}
}

func (o *Object) Value() interface{} {
	return o.value
}

func EnsureObject(obj interface{}) Object {
	if node, ok := obj.(Object); ok {
		return node
	}
	return Object{obj}
}
