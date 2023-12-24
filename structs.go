package main

type myObject struct {
	str string
	objs map[string]int
	val float64
}

func CreateObj(name string) myObject{
	obj := myObject{
		str: name,
		objs: map[string]int{},
		val: 69,
	}

	return obj
}