package intset

import (
	"bytes"
	"fmt"
)

const intSize = 32 << (^uint(0) >> 63)

// IntSetは負でない小さな整数のセットです。
// そのゼロ値は空セットを表しています。
type IntSet struct {
	words []uint
}

// Hasは負でない値xをセットが含んでいるか否かを報告します。
func (s *IntSet) Has(x int) bool {
	word, bit := x/intSize, uint(x%intSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Addはセットに負でない値xを追加します。
func (s *IntSet) Add(x int) {
	word, bit := x/intSize, uint(x%intSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWithは、sとtの和集合をsに設定します。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Stringは"{1 2 3}"の形式の文字列としてセットを返します。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", intSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Lenはセットの要素数を返す。
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// Removeはセットからxを取り除く。
func (s *IntSet) Remove(x int) {
	word, bit := x/intSize, uint(x%intSize)
	s.words[word] &= ^(1 << bit)
}

// Clearはセットから全ての要素を取り除く。
func (s *IntSet) Clear() {
	s.words = nil
}

// Copyはセットのコピーを返す。
func (s *IntSet) Copy() *IntSet {
	var copiedSet IntSet
	copiedSet.words = make([]uint, 0)
	copiedSet.words = append(copiedSet.words, s.words...)
	return &copiedSet
}

func (s *IntSet) AddAll(values ...int) {
	for _, v := range values {
		s.Add(v)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if s.words == nil {
		return
	} else if t.words == nil {
		s.words = nil
		return
	}
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i := 0; i < len(s.words); i++ {
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = (s.words[i] &^ tword) | (tword &^ s.words[i])
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Elems() []int {
	elems := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, intSize*i+j)
			}
		}
	}
	return elems
}
