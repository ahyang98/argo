package dp

import "fmt"

type RegularExpressionMatch struct {
}

func NewRegularExpressionMatch() *RegularExpressionMatch {
	return &RegularExpressionMatch{}
}

func (r *RegularExpressionMatch) Match() {

	if r.isMatch("", "a*") {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func (r *RegularExpressionMatch) isMatch(s string, p string) bool {
	var t string
	//连续多个*的情况，缩成一个*
	for _, str := range p {
		if str == '*' && t[len(t)-1] == '*' {
			continue
		}
		t = t + string(str)
	}
	p = t
	n := len(s)
	m := len(p)
	const N = 310
	var f [N][N]bool
	s = " " + s
	p = " " + p
	//都是空
	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			//下一个字符是*的先跳过，再-*一起判断
			if j+1 <= m && p[j+1] == '*' {
				continue
			}
			if i > 0 && p[j] != '*' {
				f[i][j] = f[i-1][j-1] && (s[i] == p[j] || p[j] == '.')
			} else if p[j] == '*' {
				f[i][j] = f[i][j-2] || i > 0 && f[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return f[n][m]
}
