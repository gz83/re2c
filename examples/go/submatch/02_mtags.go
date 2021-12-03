// Code generated by re2c, DO NOT EDIT.
//line "go/submatch/02_mtags.re":1
//go:generate re2go $INPUT -o $OUTPUT
package main

import (
	"reflect"
	"testing"
)

const (
	mtagRoot int = -1
	mtagNil int = -2
)

type mtagElem struct {
	val  int
	pred int
}

type mtagTrie = []mtagElem

func createTrie(capacity int) mtagTrie {
	return make([]mtagElem, 0, capacity)
}

func mtag(trie *mtagTrie, tag int, val int) int {
	*trie = append(*trie, mtagElem{val, tag})
	return len(*trie) - 1
}

// Recursively unwind both tag histories and consruct submatches.
func unwind(trie mtagTrie, x int, y int, str string) []string {
	if x == mtagRoot && y == mtagRoot {
		return []string{}
	} else if x == mtagRoot || y == mtagRoot {
		panic("tag histories have different length")
	} else {
		xval := trie[x].val
		yval := trie[y].val
		ss := unwind(trie, trie[x].pred, trie[y].pred, str)

		// Either both tags should be nil, or none of them.
		if xval == mtagNil && yval == mtagNil {
			return ss
		} else if xval == mtagNil || yval == mtagNil {
			panic("tag histories positive/negative tag mismatch")
		} else {
			s := str[xval:yval]
			return append(ss, s)
		}
	}
}

func lex(str string) []string {
	var cursor, marker int
	trie := createTrie(256)
	x := mtagRoot
	y := mtagRoot
	
//line "go/submatch/02_mtags.go":62
	yytm1 := mtagRoot
	yytm2 := mtagRoot
//line "go/submatch/02_mtags.re":58


	
//line "go/submatch/02_mtags.go":69
{
	var yych byte
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		yytm2 = mtag(&trie, yytm2, mtagNil)
		yytm1 = mtag(&trie, yytm1, mtagNil)
		goto yy2
	case 'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z':
		yytm1 = mtag(&trie, yytm1, cursor)
		goto yy6
	default:
		goto yy4
	}
yy2:
	cursor += 1
	x = yytm1
	y = yytm2
//line "go/submatch/02_mtags.re":73
	{ return unwind(trie, x, y, str) }
//line "go/submatch/02_mtags.go":90
yy4:
	cursor += 1
yy5:
//line "go/submatch/02_mtags.re":74
	{ return nil }
//line "go/submatch/02_mtags.go":96
yy6:
	cursor += 1
	marker = cursor
	yych = str[cursor]
	switch (yych) {
	case ';':
		yytm2 = mtag(&trie, yytm2, cursor)
		goto yy7
	case 'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z':
		goto yy9
	default:
		goto yy5
	}
yy7:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		goto yy2
	case 'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z':
		yytm1 = mtag(&trie, yytm1, cursor)
		goto yy9
	default:
		goto yy8
	}
yy8:
	cursor = marker
	goto yy5
yy9:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case ';':
		yytm2 = mtag(&trie, yytm2, cursor)
		goto yy7
	case 'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z':
		goto yy9
	default:
		goto yy8
	}
}
//line "go/submatch/02_mtags.re":75

}

func TestLex(t *testing.T) {
	var tests = []struct {
		str string
		res []string
	}{
		{"\000", []string{}},
		{"one;two;three;\000", []string{"one", "two", "three"}},
		{"one;two\000", nil},
	}

	for _, x := range tests {
		t.Run(x.str, func(t *testing.T) {
			res := lex(x.str)
			if !reflect.DeepEqual(res, x.res) {
				t.Errorf("got %v, want %v", res, x.res)
			}
		})
	}
}
