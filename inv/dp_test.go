package inv

import (
	"testing"
	"fmt"
)

func maxprofit(price []int) int {
	dp := make([]int, len(price))
	dp[0] = price[0]
	for i:=1;i<len(price);i++ {
		switch {
		case price[i] < dp[i-1]:
			dp[i] = price[i]
		default:
			dp[i] = dp[i-1]
		}
	}
	max := 0
	profit := 0
	for i:=0;i<len(price);i++ {
		 profit = price[i] - dp[i]
		 if profit > max {
			 max = profit
		 }
	 }
	 return max
}

func TestProfit(t *testing.T) {
	p := []int {7,1,5,3,6,4}
	res := maxprofit(p)
	if res != 5 {
		t.Errorf("expected: %d ,get %d", 5, res)
	}
}

func TestPack(t *testing.T) {
	v := []int{2,4,6}
	w := []int{1,1,5}
	V := 6
	res := solvePack(v,w,V)
	fmt.Println(res)
}

func solvePack(v, w []int, V int) int {
	var dp [][]int
	N := len(v)
	dp = make([][]int, N+1)
	for i:=0; i<N+1; i++ {
		dp[i] = make([]int, V+1)
	}
	for i:=1; i<=N; i++ {
		val := v[i-1]
		wei := w[i-1]
		for j:=1; j<=V; j++ {
			if j>= wei {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wei]+ val)
			} else {
				dp[i][j] = dp[i-1][j]
				
			}
		}
	}
	fmt.Printf("%v", dp)
	return dp[N][V]
}
func max(a, b int) int {
	if a>b {
		return a
	} else {
		return b
	}
}

//variant of pack puzzle 
func main01() {
    var total, num int
    fmt.Scanf("%d %d", &total, &num)
    total/=10
    v := make([]int, num)
    w := make([]int, num)
    b := make([]int, num)
    for i:=0;i<num;i++ {
        fmt.Scanf("%d %d %d", &v[i], &w[i], &b[i])
        v[i] /= 10
        
    }
    fmt.Println(solve(v,w,b,total))
}

func solve(v, w, b []int, total int) int {
    dp := make([][]int, len(v)+1)
    for i:=0; i<len(dp); i++ {
        dp[i] = make([]int, total+1)
    }
    m := map[string]int{}
    
    
    for i:=1; i<=len(v); i++ {
            vv := v[i-1]
            ww := w[i-1]
            bb := b[i-1]
        for j:=1; j<=total; j++ {
            if j < vv {
                dp[i][j] = dp[i-1][j]
                m[key(i,j)] |= m[key(i-1,j)] 
                continue
            } else {
                if bb == 0 || ( bb>0 && ( (1 << bb) & m[key(i-1,j-vv)]  ) > 0 ) {
                    if bb ==0  && (m[key(i-1,j-vv)] & (1 << i) > 0) {
                        dp[i][j] = dp[i-1][j]
                        m[key(i,j)] |= m[key(i-1,j)]  
                        continue
                    }
                    if dp[i-1][j] < dp[i-1][j-vv] + vv*ww {
                        dp[i][j] = dp[i-1][j-vv] + vv*ww
                        m[key(i,j)] |=  m[key(i-1,j-vv)]
                        m[key(i,j)] |= (1 << i)
                        continue 
                    } else {
                        dp[i][j] = dp[i-1][j]
                        m[key(i,j)] |= m[key(i-1,j)]  
                        continue
                    }
                    
                    //dp[i][j].m = append(dp[i][j].m , i)
                } else {
                         vv2 := v[bb-1]
                         ww2 := w[bb-1]
                         if j >=  vv + vv2 {
                             if dp[i-1][j] < dp[i-1][j-vv-vv2] + vv*ww + vv2*ww2 {
                                 dp[i][j] = dp[i-1][j-vv-vv2] + vv*ww + vv2*ww2
                                 m[key(i,j)] |= m[key(i-1,j-vv-vv2)]
                                 m[key(i,j)] |= (1 << bb)
                                 continue
                             } else {
                                 dp[i][j] = dp[i-1][j]
                                 m[key(i,j)] |= m[key(i-1,j)] 
                                 continue
                             }
                             
                            //dp[i][j].m = append(dp[i][j].m, i)
                        } else {
                            dp[i][j] = dp[i-1][j]
                             m[key(i,j)] |=  m[key(i-1,j)] 
                        }                       
                        
                    }
                }
            }
            
    }

    for i:=len(dp)-1; i>0; i-- {
        if dp[i][total] > 0 {
            return dp[i][total] * 10
        
      } 
    }
    return 0
}

func key(i, j int) string {
    return fmt.Sprintf("%d:%d", i, j)
}

// 300
func lst(a []int) int {
	dp := make([]int, len(a))
	for idx  :=range dp {
		dp[idx] = 1 
	}
	for i:=0;i<len(dp);i++ {
		for j:=0;j<i;j++ {
			if a[i] > a[j] {
				dp[i] =max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0 
	for i:=0;i<len(dp);i++ {
		res = max(res, dp[i])
	}
	return res
}


func hot(a [][2]int, m int) {
	max := 0
	min := 1000
	di := map[int]int{}
	for _, e :=range a {
		if e[0] + e[1] > max {
			max = e[0] + e[1]
		}  
		if e[0] + e[1] < min {
			min = e[0] + e[1]
		}
		di[e[0]+e[1]]++
	}
	da := make([]int, max+1)
	for key, val :=range di {
		da[key] = val
	}
	fmt.Println(da)
	dp := make([]int, max+1)
	for i:=min; i<min+m; i++ {
		dp[i] = da[i]
	}
	for i:=min;i<len(dp);i++ {
		dp[i] = maxc(dp[i-m]+da[i], dp[i-1])
	}
	fmt.Println(dp[max])
}
func maxc(a, b int) int {
	if a > b {
	return a
	}
	return b
}

func TestHot(t *testing.T) {
	a := [][2]int{ {1,2},{1,3},{2,3}}
	m := 1
	hot(a, m)
}

