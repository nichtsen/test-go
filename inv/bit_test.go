package inv
import "testing"

func hammingDistance(x, y int) int {
	diff := x ^ y
	cnt := 0
	for diff > 0 {
		if diff & 1 == 1{
			cnt++
		}
		diff >>= 1
	}
	return cnt
}


func TestHamming(t *testing.T) {
	x, y := 1,4
	res := hammingDistance(x, y)
	if res != 2 {
		t.Errorf("expected to be %d, but got %d", 2, res)
	}
}

func reverseBit(n uint32) uint32 {
	var ans uint32
	for i:=0;i<32;i++ {
		ans <<= 1
		ans += n & 1
		n >>= 1
	}
	return ans
}

//thanks to eXclusive OR operator that is associative and commutative
func onlyonce(a []int) (ans int) { 
	for _, v := range a {
		ans ^= v
	}
	return ans
}

func maxProduct(s []string) int {
	res := 0
	hash := make(map[uint32]int)
	for _, str := range s {
		var bits uint32
		for _, byt := range str {
			bits |= 1 << (byt - 'a')
		}
		if len(str) > hash[bits] {
			hash[bits] = len(str)
		}
		for bit, size := range hash {
			// non replicated character 
			if bit & bits == 0 && size * len(str) > res  {
				res = size * len(str)
			}
		}
	}
	return res
}

func countBit(n int) []int {
	dp := make([]int, n+1)
	for i:=1;i<n+1;i++ {
		if i & 1 == 0 {
			dp[i] = dp[i>>1]
		} else {
			dp[i] = dp[i-1]+1
		}
	}
	return dp
}
				
	
