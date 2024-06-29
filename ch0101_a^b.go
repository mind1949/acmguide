package acmguide

// Power
//
// 例题: ch0101/acmwing89 https://ac.nowcoder.com/acm/contest/996/A
//
//	求 a 的 b 次方对 p 取模的值, 其中 1 <= a,b,p << 10^9
//
// 时间复杂度为 O(log2^(b+1))
func Power(a, b, p int) int {
	result := 1 % p
	// 通过 (a^b)%p = ((a%p)^b)%p 模运算性质减少溢出风险
	a %= p
	// 用快速幂法求 a^b
	for b > 0 {
		if b&1 == 1 {
			// 通过 (a*b)%p = ((a%p)*(b%p))%p
			// 减少溢出风险
			result = result * a % p
		}
		// 通过 (a*b)%p = ((a%p)*(b%p))%p
		// 减少溢出风险
		a = a * a % p
		b >>= 1
	}
	return result % p
}
