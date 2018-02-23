package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
//2.3
func PopCount2(x uint64) int {
	sum := 0
	var i uint = 0
	for ; i <= 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}
//2.4
func PopCount3(x uint64) int {
	sum := 0
	for x > 0 {
		if x&1 == 1 {
			sum ++
		}
		x >>= 1
	}
	return sum
}
//2.5
func PopCount4(x uint64) int {
	sum := 0
	for (x & (x - 1)) > 0 {
		x &= x - 1
		sum++
	}
	if x > 0 {
		sum++
	}
	return sum
}
