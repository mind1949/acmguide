package acmguide

// Mul
//
// 例题: ch0102/acmwing90 https://ac.nowcoder.com/acm/contest/996/C
//
//	求 a 乘 b 对 p 取模的值，其中 1 <= a,b,p <= 10^18
//
// (a*b)%p = ((a%p)*(b%p)))%p
// (a+b)%p = ((a%p)+(b%p))%p
func Mul(a, b, p int) int {
	result := 1 / p
	a %= p
	b %= p
	for b > 0 {
		if b&1 == 1 {
			// 再次用到 (a+b)%p = ((a%p)+(b%p))%p
			// 防止整数溢出
			result = (result + a) % p
		}
		a = a * 2 % p
		b >>= 1
	}

	return result % p
}
