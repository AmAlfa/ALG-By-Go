package _removeDuplicatesFromSortedArray

import (
	"math"
	"testing"
)
//统计所有小于非负整数 n 的质数的数量。
//
//示例:
//
//输入: 10
//输出: 4
//解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。


func TestHello(t *testing.T) {
	n := 10
	t.Log(countPrimes(n))
}

//时间最优
//国外1
func countPrimes1(n int) int {
	if n < 3 {
		return 0
	}
	if n == 3 {
		return 1
	}
	res := n / 2
	prime := make([]bool, n)
	for i := 3; i*i < n; i += 2 {
		if prime[i] {
			continue
		}
		for j := i * i; j < n; j += 2 * i {
			if !prime[j] {
				res--
				prime[j] = true
			}
		}
	}
	return res
}
//国外2
func countPrimes2(n int) int {
	isNotPrime := make([]bool, n)

	for i := 2; i*i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * i; j < n; j = j + i {
			isNotPrime[j] = true
		}
	}
	var count int
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			count++
		}
	}
	return count
}
//国内3
func countPrimes3(n int) int {
	if n < 3 {
		return 0
	}
	cnt := n - 2
	nums := make([]bool, n)
	m := int(math.Sqrt(float64(n)))
	for i := 2; i <= m; i++ {
		if !nums[i] {
			for j := i; i*j < n; j++ {
				if !nums[i*j] {
					nums[i*j] = true
					cnt--
				}
			}
		}
	}
	return cnt
}
//国内4
func countPrimes4(n int) int {
	if n < 3{
		return 0
	}

	var m int = n / 2

	a := make([]bool, m)
	b := m

	var k int = (int(math.Sqrt(float64(n))) + 1) / 2
	for i:=1;i<k;i++{
		if !a[i] {
			for j:= 2*i*i+2*i;j<m;j+=2*i+1{
				if !a[j] {
					a[j] = true
					b --
				}
			}
		}
	}

	return b
}
//内存最优
//国外5
func countPrimes5(n int) int {
	count := 0
	for i := 1; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

func isPrime(num int) bool {
	if num == 1 { return false }
	if num == 2 { return true }
	bound := int(math.Ceil(math.Sqrt(float64(num))))
	for i := 2; i <= bound; i++ {
		if num % i == 0 {
			return false
		}
	}

	return true
}
//国外6
func isPrime(n int) bool {
	if n % 2 == 0 {
		return false
	}

	for i:= 3; i*i <= n; i=i+2 {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func countPrimes6(n int) int {
	count := 0
	if n<= 2 {
		return count
	}

	count = 1
	for i:=3; i< n; i++ {
		if isPrime(i) == true {
			count++
		}
	}
	return count
}
//国内7
//国内8
