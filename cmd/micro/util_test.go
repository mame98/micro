package main

import (
	"reflect"
	"testing"
)

func TestNumOccurences(t *testing.T) {
	var tests = []struct {
		inputStr  string
		inputChar byte
		want      int
	}{
		{"aaaa", 'a', 4},
		{"\trfd\ta", '\t', 2},
		{"∆ƒ\tø ® \t\t", '\t', 3},
	}
	for _, test := range tests {
		if got := NumOccurrences(test.inputStr, test.inputChar); got != test.want {
			t.Errorf("NumOccurences(%s, %c) = %d", test.inputStr, test.inputChar, got)
		}
	}
}

func TestSpaces(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{4, "    "},
		{0, ""},
	}
	for _, test := range tests {
		if got := Spaces(test.input); got != test.want {
			t.Errorf("Spaces(%d) = \"%s\"", test.input, got)
		}
	}
}

func TestIsWordChar(t *testing.T) {
	if IsWordChar("t") == false {
		t.Errorf("IsWordChar(t) = false")
	}
	if IsWordChar("T") == false {
		t.Errorf("IsWordChar(T) = false")
	}
	if IsWordChar("5") == false {
		t.Errorf("IsWordChar(5) = false")
	}
	if IsWordChar("_") == false {
		t.Errorf("IsWordChar(_) = false")
	}
	if IsWordChar("ß") == false {
		t.Errorf("IsWordChar(ß) = false")
	}
	if IsWordChar("~") == true {
		t.Errorf("IsWordChar(~) = true")
	}
	if IsWordChar(" ") == true {
		t.Errorf("IsWordChar( ) = true")
	}
	if IsWordChar(")") == true {
		t.Errorf("IsWordChar()) = true")
	}
	if IsWordChar("\n") == true {
		t.Errorf("IsWordChar(\n)) = true")
	}
}

func TestJoinAndSplitCommandArgs(t *testing.T) {
	tests := []struct {
		Query  []string
		Wanted string
	}{
		{[]string{`test case`}, `"test case"`},
		{[]string{`quote "test"`}, `"quote \"test\""`},
		{[]string{`slash\\\ test`}, `"slash\\\\\\ test"`},
		{[]string{`path 1`, `path\" 2`}, `"path 1" "path\\\" 2"`},
		{[]string{`foo`}, `foo`},
		{[]string{`foo\"bar`}, `"foo\\\"bar"`},
		{[]string{``}, ``},
		{[]string{`"`}, `"\""`},
		{[]string{`a`, ``}, `a `},
		{[]string{``, ``, ``, ``}, `   `},
		{[]string{"\n"}, `"\n"`},
		{[]string{"foo\tbar"}, `"foo\tbar"`},
	}

	for i, test := range tests {
		if result := JoinCommandArgs(test.Query...); test.Wanted != result {
			t.Errorf("JoinCommandArgs failed at Test %d\nGot: %q", i, result)
		}

		if result := SplitCommandArgs(test.Wanted); !reflect.DeepEqual(test.Query, result) {
			t.Errorf("SplitCommandArgs failed at Test %d\nGot: `%q`", i, result)
		}
	}

	splitTests := []struct {
		Query  string
		Wanted []string
	}{
		{`"hallo""Welt"`, []string{`halloWelt`}},
		{`"hallo" "Welt"`, []string{`hallo`, `Welt`}},
		{`\"`, []string{`\"`}},
		{`"foo`, []string{`"foo`}},
		{`"foo"`, []string{`foo`}},
		{`"\"`, []string{`"\"`}},
		{`"C:\\"foo.txt`, []string{`C:\foo.txt`}},
		{`"\n"new"\n"line`, []string{"\nnew\nline"}},
	}

	for i, test := range splitTests {
		if result := SplitCommandArgs(test.Query); !reflect.DeepEqual(test.Wanted, result) {
			t.Errorf("SplitCommandArgs failed at Split-Test %d\nGot: `%q`", i, result)
		}
	}
}

func TestStringWidth(t *testing.T) {
	tabsize := 4
	if w := StringWidth("1\t2", tabsize); w != 5 {
		t.Error("StringWidth 1 Failed. Got", w)
	}
	if w := StringWidth("\t", tabsize); w != 4 {
		t.Error("StringWidth 2 Failed. Got", w)
	}
	if w := StringWidth("1\t", tabsize); w != 4 {
		t.Error("StringWidth 3 Failed. Got", w)
	}
	if w := StringWidth("\t\t", tabsize); w != 8 {
		t.Error("StringWidth 4 Failed. Got", w)
	}
	if w := StringWidth("12\t2\t", tabsize); w != 8 {
		t.Error("StringWidth 5 Failed. Got", w)
	}
}

func TestWidthOfLargeRunes(t *testing.T) {
	tabsize := 4
	if w := WidthOfLargeRunes("1\t2", tabsize); w != 2 {
		t.Error("WidthOfLargeRunes 1 Failed. Got", w)
	}
	if w := WidthOfLargeRunes("\t", tabsize); w != 3 {
		t.Error("WidthOfLargeRunes 2 Failed. Got", w)
	}
	if w := WidthOfLargeRunes("1\t", tabsize); w != 2 {
		t.Error("WidthOfLargeRunes 3 Failed. Got", w)
	}
	if w := WidthOfLargeRunes("\t\t", tabsize); w != 6 {
		t.Error("WidthOfLargeRunes 4 Failed. Got", w)
	}
	if w := WidthOfLargeRunes("12\t2\t", tabsize); w != 3 {
		t.Error("WidthOfLargeRunes 5 Failed. Got", w)
	}
}
