package main

import (
	"reflect"
	"testing"
)

// func TestPOS(t *testing.T) {
// 	tests := []struct {
// 		word     wordType
// 		n        int
// 		expected wordType
// 	}{
// 		{wordType("hello "), 0, wordType("h")},
// 		{wordType("hello"), 2, wordType("l")},
// 		{wordType("hello"), 4, wordType("o")},
// 		{wordType("hello"), 5, wordType{}},
// 		{wordType("h"), 0, wordType("h")},
// 		{wordType(""), 5, wordType{}},
// 		{wordType(""), 0, wordType{}},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := SUBSTRING(tt.args.word, tt.args.looking); got != tt.want {
// 				t.Errorf("SUBSTRING() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestPOS(t *testing.T) {
	type args struct {
		word wordType
		n    int
	}
	tests := []struct {
		name string
		args args
		want wordType
	}{
		{
			name: "n is positive",
			args: args{
				word: wordType("hello"),
				n:    2,
			},
			want: wordType("l"),
		},
		{
			name: "n is zero",
			args: args{
				word: wordType("hello "),
				n:    0,
			},
			want: wordType("h"),
		},
		{
			name: "n greater then length of word",
			args: args{
				word: wordType("hello"),
				n:    5,
			},
			want: wordType{},
		},
		{
			name: "empty word with position zero",
			args: args{
				word: wordType(""),
				n:    0,
			},
			want: wordType{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := POS(tt.args.word, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("POS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSUBSTRING(t *testing.T) {
	type args struct {
		word    wordType
		looking wordType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				word:    wordType("abcdefg"),
				looking: wordType("cde"),
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				word:    wordType("abcdefg"),
				looking: wordType("xyz"),
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				word:    wordType("abcdefg"),
				looking: wordType("g"),
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				word:    wordType("abcdefg"),
				looking: wordType("ab"),
			},
			want: true,
		},
		{
			name: "Test Case 5",
			args: args{
				word:    wordType("abcdefg"),
				looking: wordType("abcdefgh"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SUBSTRING(tt.args.word, tt.args.looking); got != tt.want {
				t.Errorf("SUBSTRING() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEQWORD(t *testing.T) {
	type args struct {
		word1 wordType
		word2 wordType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test case 1", args: args{word1: wordType("hello"), word2: wordType("hello")}, want: true},
		{name: "Test case 2", args: args{word1: wordType(""), word2: wordType("hello")}, want: false},
		{name: "Test case 3", args: args{word1: wordType("hello"), word2: wordType("world")}, want: false},
		{name: "Test case 4", args: args{word1: wordType(""), word2: wordType("")}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EQWORD(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("EQWORD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMEMBERSHIP(t *testing.T) {
	type args struct {
		language []wordType
		word     wordType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty language and Empty word",
			args: args{
				language: []wordType{},
				word:     wordType{},
			},
			want: false,
		},
		{
			name: "Non-empty language and Empty word",
			args: args{
				language: []wordType{wordType("Hello"), wordType("World")},
				word:     wordType{},
			},
			want: false,
		},
		{
			name: "Empty language and Non-empty word",
			args: args{
				language: []wordType{},
				word:     wordType("Hello"),
			},
			want: false,
		},
		{
			name: "word not in language",
			args: args{
				language: []wordType{wordType("Hello"), wordType("World")},
				word:     wordType("Today"),
			},
			want: false,
		},
		{
			name: "word in language",
			args: args{
				language: []wordType{wordType("Hello"), wordType("World")},
				word:     wordType("Hello"),
			},
			want: true,
		},
		{
			name: "empty word in language",
			args: args{
				language: []wordType{wordType("Hello"), wordType("World"), wordType("")},
				word:     wordType(""),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MEMBERSHIP(tt.args.language, tt.args.word); got != tt.want {
				t.Errorf("MEMBERSHIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func areListsEqual(list1, list2 []wordType) bool {
	if len(list1) != len(list2) {
		return false
	}

	map1 := make(map[string]int)
	map2 := make(map[string]int)

	for _, item := range list1 {
		map1[string(item)]++
	}

	for _, item := range list2 {
		map2[string(item)]++
	}
	return reflect.DeepEqual(map1, map2)
}

func TestREMOVE(t *testing.T) {
	type args struct {
		language []wordType
		word     wordType
	}
	tests := []struct {
		name string
		args args
		want []wordType
	}{
		{
			name: "Remove Word From Language",
			args: args{
				language: []wordType{wordType("hello"), wordType("world"), wordType("hello"), wordType("code")},
				word:     wordType("hello"),
			},
			want: []wordType{wordType("world"), wordType("code")},
		},
		{
			name: "Remove Word Not In Language",
			args: args{
				language: []wordType{wordType("hello"), wordType("world"), wordType("hello"), wordType("code")},
				word:     wordType("python"),
			},
			want: []wordType{wordType("hello"), wordType("world"), wordType("hello"), wordType("code")},
		},
		{
			name: "Remove Word From Empty Language",
			args: args{
				language: []wordType{},
				word:     wordType("hello"),
			},
			want: []wordType{},
		},
		{
			name: "Remove Word From Single Word Language",
			args: args{
				language: []wordType{wordType("hello")},
				word:     wordType("hello"),
			},
			want: []wordType{},
		},
		{
			name: "Remove empty Word From Language with empty word",
			args: args{
				language: []wordType{wordType("hello"), wordType(""), wordType("world"), wordType("hello"), wordType("code")},
				word:     wordType(""),
			},
			want: []wordType{wordType("hello"), wordType("world"), wordType("hello"), wordType("code")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := REMOVE(tt.args.language, tt.args.word); !areListsEqual(got, tt.want) {
				t.Errorf("REMOVE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestREMOVEREPETITIONS(t *testing.T) {
	type args struct {
		language []wordType
	}
	tests := []struct {
		name string
		args args
		want []wordType
	}{
		{
			name: "Remove duplicate words",
			args: args{
				language: []wordType{
					wordType("hello"),
					wordType(""),
					wordType("world"),
					wordType("hello"),
					wordType("code"),
				},
			},
			want: []wordType{
				wordType("hello"),
				wordType(""),
				wordType("world"),
				wordType("code"),
			},
		},
		{
			name: "Remove duplicate fruits",
			args: args{
				language: []wordType{
					wordType("apple"),
					wordType("banana"),
					wordType("apple"),
					wordType("cherry"),
					wordType("banana"),
				},
			},
			want: []wordType{
				wordType("apple"),
				wordType("banana"),
				wordType("cherry"),
			},
		},
		{
			name: "No duplicate words",
			args: args{
				language: []wordType{
					wordType("hello"),
					wordType("world"),
				},
			},
			want: []wordType{
				wordType("hello"),
				wordType("world"),
			},
		},
		{
			name: "Empty language",
			args: args{
				language: []wordType{},
			},
			want: []wordType{},
		},
		{
			name: "Single word",
			args: args{
				language: []wordType{
					wordType("hello"),
					wordType("hello"),
					wordType("hello"),
					wordType("hello"),
				},
			},
			want: []wordType{
				wordType("hello"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := REMOVEREPETITIONS(tt.args.language); !areListsEqual(got, tt.want) {
				t.Errorf("REMOVEREPETITIONS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEQLANG(t *testing.T) {
	type args struct {
		language1 []wordType
		language2 []wordType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Both languages are empty",
			args: args{
				language1: []wordType{},
				language2: []wordType{},
			},
			want: true,
		},
		{
			name: "Both languages have the same words",
			args: args{
				language1: []wordType{
					wordType("Hello"),
					wordType("World"),
				},
				language2: []wordType{
					wordType("World"),
					wordType("Hello"),
				},
			},
			want: true,
		},
		{
			name: "Both languages have different words",
			args: args{
				language1: []wordType{
					wordType("Hello"),
					wordType("World"),
				},
				language2: []wordType{
					wordType("Bonjour"),
					wordType("Monde"),
				},
			},
			want: false,
		},
		{
			name: "One language is empty, the other has words",
			args: args{
				language1: []wordType{},
				language2: []wordType{
					wordType("Hello"),
					wordType("World"),
				},
			},
			want: false,
		},
		{
			name: "One language has words, the other is empty",
			args: args{
				language1: []wordType{
					wordType("Hello"),
					wordType("World"),
				},
				language2: []wordType{},
			},
			want: false,
		},
		{
			name: "Kacper example",
			args: args{
				language1: []wordType{
					wordType("aaa"),
					wordType("a"),
				},
				language2: []wordType{
					wordType("aaa"),
					wordType("a"),
					wordType("c"),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EQLANG(tt.args.language1, tt.args.language2); got != tt.want {
				t.Errorf("EQLANG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCONLANG(t *testing.T) {
	type args struct {
		language1 listType
		language2 listType
	}
	tests := []struct {
		name string
		args args
		want listType
	}{
		{
			name: "Empty + Empty",
			args: args{
				language1: listOfStringToListType([]string{}),
				language2: listOfStringToListType([]string{}),
			},
			want: listOfStringToListType([]string{}),
		},
		{
			name: "Empty + Non-Empty",
			args: args{
				language1: listOfStringToListType([]string{}),
				language2: listType{wordType([]rune("Chinese")), wordType([]rune("Vietnamese"))},
			},
			want: listOfStringToListType([]string{}),
		},
		{
			name: "Non-Empty + Empty",
			args: args{
				language1: listType{wordType([]rune("Chinese")), wordType([]rune("Vietnamese"))},
				language2: listOfStringToListType([]string{}),
			},
			want: listOfStringToListType([]string{}),
		},
		{
			name: "size 1 + size 1",
			args: args{
				language1: listOfStringToListType([]string{"hello"}),
				language2: listOfStringToListType([]string{"world"}),
			},
			want: listOfStringToListType([]string{"helloworld"}),
		},
		{
			name: "size 2 + size 1",
			args: args{
				language1: listOfStringToListType([]string{"hello", "world"}),
				language2: listOfStringToListType([]string{"hello"}),
			},
			want: listOfStringToListType([]string{"hellohello", "worldhello"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CONLANG(tt.args.language1, tt.args.language2); !areListsEqual(got, tt.want) {
				t.Errorf("CONLANG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPOWLANG(t *testing.T) {
	type args struct {
		language listType
		n        int
	}
	tests := []struct {
		name string
		args args
		want listType
	}{
		{
			name: "Empty",
			args: args{
				language: listOfStringToListType([]string{}),
				n:        0,
			},
			want: listOfStringToListType([]string{}),
		},
		{
			name: "Power of 0",
			args: args{
				language: listOfStringToListType([]string{"hello", "world"}),
				n:        0,
			},
			want: listOfStringToListType([]string{}),
		},
		{
			name: "Power of 1",
			args: args{
				language: listOfStringToListType([]string{"hello", "world"}),
				n:        1,
			},
			want: listOfStringToListType([]string{"hello", "world"}),
		},
		{
			name: "Power of 2",
			args: args{
				language: listOfStringToListType([]string{"hello", "world"}),
				n:        2,
			},
			want: listOfStringToListType([]string{"hellohello", "helloworld", "worldhello", "worldworld"}),
		},
		{
			name: "Power of 2 with empty",
			args: args{
				language: listOfStringToListType([]string{"", "world"}),
				n:        2,
			},
			want: listOfStringToListType([]string{"", "world", "worldworld"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := POWLANG(tt.args.language, tt.args.n); !areListsEqual(got, tt.want) {
				t.Errorf("POWLANG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUNION(t *testing.T) {
	type args struct {
		language1 listType
		language2 listType
	}
	tests := []struct {
		name string
		args args
		want listType
	}{
		{
			name: "Empty + Empty",
			args: args{
				language1: listOfStringToListType([]string{}),
				language2: listOfStringToListType([]string{}),
			},
			want: listOfStringToListType([]string{}),
		},
		{
			name: "Empty + Non-Empty",
			args: args{
				language1: listOfStringToListType([]string{}),
				language2: listOfStringToListType([]string{"Chinese", "Vietnamese"}),
			},
			want: listOfStringToListType([]string{"Chinese", "Vietnamese"}),
		},
		{
			name: "Non-Empty + Empty",
			args: args{
				language1: listOfStringToListType([]string{"Chinese", "Vietnamese"}),
				language2: listOfStringToListType([]string{}),
			},
			want: listOfStringToListType([]string{"Chinese", "Vietnamese"}),
		},
		{
			name: "size 1 + size 1",
			args: args{
				language1: listOfStringToListType([]string{"hello"}),
				language2: listOfStringToListType([]string{"world"}),
			},
			want: listOfStringToListType([]string{"hello", "world"}),
		},
		{
			name: "size 2 + size 1",
			args: args{
				language1: listOfStringToListType([]string{"hello", "world"}),
				language2: listOfStringToListType([]string{"hello"}),
			},
			want: listOfStringToListType([]string{"hello", "world"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UNION(tt.args.language1, tt.args.language2); !areListsEqual(got, tt.want) {
				t.Errorf("UNION() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMEET(t *testing.T) {
	type args struct {
		language1 listType
		language2 listType
	}
	tests := []struct {
		name string
		args args
		want listType
	}{
		{
			name: "Empty languages",
			args: args{
				language1: listOfStringToListType([]string{}),
				language2: listOfStringToListType([]string{}),
			},
			want: listOfStringToListType([]string{}),
		},
		{
			name: "Languages with some common elements",
			args: args{
				language1: listOfStringToListType([]string{"English", "Spanish", "French"}),
				language2: listOfStringToListType([]string{"Spanish", "German", "Italian"}),
			},
			want: listOfStringToListType([]string{"Spanish"}),
		},
		{
			name: "Languages with all common elements",
			args: args{
				language1: listOfStringToListType([]string{"Spanish", "German", "Italian"}),
				language2: listOfStringToListType([]string{"Spanish", "German", "Italian"}),
			},
			want: listOfStringToListType([]string{"Spanish", "German", "Italian"}),
		},
		{
			name: "Languages with no common elements",
			args: args{
				language1: listOfStringToListType([]string{"English", "French"}),
				language2: listOfStringToListType([]string{"German", "Italian"}),
			},
			want: listOfStringToListType([]string{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MEET(tt.args.language1, tt.args.language2); !areListsEqual(got, tt.want) {
				t.Errorf("MEET() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSUBWORDS(t *testing.T) {
	type args struct {
		word   wordType
		length int
	}
	tests := []struct {
		name   string
		args   args
		expect listType
	}{
		{
			name: "\"Example\" word with length 3",
			args: args{
				word:   wordType("example"),
				length: 3,
			},
			expect: listOfStringToListType([]string{"exa", "xam", "amp", "mpl", "ple"}),
		},
		{
			name: "\"Hello\" word with length 2",
			args: args{
				word:   wordType("hello"),
				length: 2,
			},
			expect: listOfStringToListType([]string{"he", "el", "ll", "lo"}),
		},
		{
			name: "\"World\" word with length 1",
			args: args{
				word:   wordType("world"),
				length: 1,
			},
			expect: listOfStringToListType([]string{"w", "o", "r", "l", "d"}),
		},
		{
			name: "Empty word with length 0",
			args: args{
				word:   wordType(""),
				length: 0,
			},
			expect: listOfStringToListType([]string{}),
		},
		{
			name: "Empty word with non-zero length",
			args: args{
				word:   wordType(""),
				length: 1,
			},
			expect: listOfStringToListType([]string{}),
		},
		{
			name: "Non-empty word with length 0",
			args: args{
				word:   wordType("World"),
				length: 0,
			},
			expect: listOfStringToListType([]string{}),
		},
		{
			name: "Repetitions",
			args: args{
				word:   wordType("hello"),
				length: 1,
			},
			expect: listOfStringToListType([]string{"h", "e", "l", "o"}),
		},
		{
			name: "length equal to word length",
			args: args{
				word:   wordType("equal"),
				length: 5,
			},
			expect: listOfStringToListType([]string{"equal"}),
		},
		{
			name: "length greater then word length",
			args: args{
				word:   wordType("greater"),
				length: 10,
			},
			expect: listOfStringToListType([]string{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SUBWORDS(tt.args.word, tt.args.length); !areListsEqual(got, tt.expect) {
				t.Errorf("SUBWORDS() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestALLSUBWORDS(t *testing.T) {
	type args struct {
		word wordType
	}
	tests := []struct {
		name string
		args args
		want listType
	}{
		{
			name: "Empty word",
			args: args{
				word: wordType(""),
			},
			want: listOfStringToListType([]string{}),
		},
		{
			name: "Single character word",
			args: args{
				word: wordType("a"),
			},
			want: listOfStringToListType([]string{"a"}),
		},
		{
			name: "Word with multiple non-repeating characters",
			args: args{
				word: wordType("abc"),
			},
			want: listOfStringToListType([]string{"a", "b", "c", "ab", "bc", "abc"}),
		},
		{
			name: "Word with multiple repeating characters",
			args: args{
				word: wordType("abbc"),
			},
			want: listOfStringToListType([]string{"a", "b", "c", "ab", "bb", "bc", "abb", "bbc", "abbc"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ALLSUBWORDS(tt.args.word); !areListsEqual(got, tt.want) {
				t.Errorf("ALLSUBWORDS() = %v, want %v", got, tt.want)
			}
		})
	}
}
