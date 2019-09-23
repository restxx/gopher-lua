package lua


func OpenBit32(L *LState) int {
	mod := L.RegisterModule(Bit32LibName, bit32Funcs)
	L.Push(mod)
	return 1
}

var bit32Funcs = map[string]LGFunction{
	"band":   bandFn,
	"bnot":   bnotFn,
	"bor":    borFn,
	"bxor":   bxorFn,
	"lshift": lshiftFn,
	"rshift": rshiftFn,
}

func bandFn(L *LState) int {
	a := L.CheckInt(1)
	b := L.CheckInt(2)
	result := a & b
	L.Push(LNumber(result))
	return 1
}

func bnotFn(L *LState) int {
	a := L.CheckInt(1)
	result := ^a
	L.Push(LNumber(result))
	return 1
}

func borFn(L *LState) int {
	a := L.CheckInt(1)
	b := L.CheckInt(2)
	result := a | b
	L.Push(LNumber(result))
	return 1
}

func bxorFn(L *LState) int {
	a := L.CheckInt(1)
	b := L.CheckInt(2)
	result := a ^ b
	L.Push(LNumber(result))
	return 1
}

func lshiftFn(L *LState) int {
	a := L.CheckInt(1)
	b := L.CheckInt(2)
	result := a << uint(b)
	L.Push(LNumber(result))
	return 1
}

func rshiftFn(L *LState) int {
	n := L.CheckInt64(1)
	n2 := L.CheckInt(2)
	result := n >> uint8(n2)
	L.Push(LNumber(result))
	return 1
}
