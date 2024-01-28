package empty_interface

import "testing"

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	println("Integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	println("string", s)
	// 	return
	// }
	// println("Unknow Type")
	switch v := p.(type) {
	case int:
		println("Integer", v)
	case string:
		println("string", v)
	default:
		println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
