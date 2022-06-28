package inv 

import (
	"testing"
	"fmt"
)

// 167
func TestSum(t *testing.T) {
	solve := func(a []int, target int) []int {
		r := len(a) -1
		l := 0
		for ;l!=r; {
			sum := a[l] +a[r]
			switch {
				case sum == target:
					return []int{r, l}
				case sum > target:
						r--
				default:
						l++
			}
		}
		return []int{-1}
	}
	a := []int{1,2,7,11,13}
	res := solve(a, 9)
	fmt.Println(res)
}
// 88 
func TestMerge(t *testing.T) {
	merge := func(nums1, nums2 []int, m, n int) {
		pos := m+n-1
		n--; m--;
		for ;pos>0; {
			switch {
			case n<0:
				nums1[pos] = nums1[m]
				m--
			case nums1[m] > nums2[n]:
				nums1[pos] = nums1[m]
				m--
			case nums1[m] <= nums2[n]:
				nums1[pos] = nums2[n]
				n--
			}
			pos--
		}
	}
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5,6}
	merge(nums1, nums2, len(nums1)- len(nums2), len(nums2))
	fmt.Println(nums1)
}

//76
func TestWindow(t *testing.T) {
	window := func(s, t string) string {
		sr := []rune(s)
		tr := []rune(t)
		var flag [256]bool
		var chars [256]int
		for _, r := range tr {
			flag[r] = true
			chars[r]++
		}
		cnt := 0; l := 0; minL := 0; minWidth := len(sr)+1;
		for r:=0;r<len(sr);r++ {
			if flag[sr[r]] {
				chars[sr[r]]--
				if chars[sr[r]] >= 0 {
					cnt++
				}
			}
			// if window has contained all chars in t
			for ;cnt==len(tr);l++ {
				width := r-l+1
				if width < minWidth { minWidth = width; minL = l;}
				if flag[sr[l]] {
					chars[sr[l]]++
					if chars[sr[l]] > 0 {
						cnt--
					}
				}
			}
		}
		switch minWidth {
			case len(sr)+1:
				return ""
			default:   
				return s[minL:minL + minWidth]
		}
	}
	s := "adbefcooobanc"
	tt := "abc"
	res := window(s, tt)
	fmt.Println(res)
} 
