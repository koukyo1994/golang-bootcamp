package word

import (
	"math/rand"
	"testing"
	"time"
)

// randomPalindromeは、疑似乱数生成器rngから長さと内容が計算された
// 回文を返します
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // 24までのランダムな長さ
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // '\u0999'までのランダムなルーン
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

// randomNotPalindromeは、疑似乱数生成器を用いて回文ではない文章を返します
func randomNotPalindrome(rng *rand.Rand) string {
	var n int
	for {
		n = rng.Intn(25)
		if n >= 2 {
			break
		}
	}
	runes := make([]rune, n)
	for {
		for i := 0; i < n; i++ {
			r := rune(rng.Intn(0x1000)) // '\u0999'までのランダムなルーン
			runes[i] = r
		}

		palindrome := true
		for i := 0; i < (n+1)/2; i++ {
			if runes[i] != runes[n-1-i] {
				palindrome = false
				break
			}
		}
		if !palindrome {
			break
		}
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// 擬似乱数生成器を初期化する
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}

		p = randomNotPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true, runes: %v", p, []rune(p))
		}
	}
}
