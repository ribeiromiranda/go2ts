package tests

type CustomInt int

type InterfaceName interface {
	Method1(arg int) CustomInt
	Method2() CustomInt
	Method3()
}
