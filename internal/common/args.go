package common

type Args struct {
	name string
	age  int
}

func NewArgs(name string, age int) *Args {
	arg := &Args{
		name: name,
		age:  age,
	}
	return arg
}
