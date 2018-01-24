package algorithms

//Given two binary strings, return their sum (also a binary string).
//
//For example,
//a = "11"
//b = "1"
//Return "100".

func addBinary(a string, b string) string {
	endA, endB := len(a)-1, len(b)-1
	if endA < endB {
		a, b = b, a
		endA, endB = endB, endA
	}
	r := make([]byte, len(a))
	var carry uint8 = 0
	f := func(foo uint8) {
		switch foo {
		case 0:
			r[endA] = '0'
		case 1:
			r[endA] = '1'
			carry = 0
		case 2:
			r[endA] = '0'
			carry = 1
		case 3:
			r[endA] = '1'
			carry = 1
		}
	}
	for endB >= 0 {
		f(a[endA]+b[endB]-48*2+carry)
		endA--
		endB--
	}
	for ; endA >= 0; endA-- {
		f(a[endA]-48+carry)
	}
	if carry == 1 {
		return "1" + string(r)
	}
	return string(r)
}
