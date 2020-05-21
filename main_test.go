package main

import (
	"testing"
)

func TestConvert(t *testing.T) {
	type (
		testin struct {
			val string
		}
		testout struct {
			result string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in:  testin{val: "0b11111111"},
			out: testout{result: "0xFF"},
		},
		{
			in:  testin{val: "0b1000"},
			out: testout{result: "0x8"},
		},
		{
			in:  testin{val: "1000"},
			out: testout{result: "0x8"},
		},
		{
			in:  testin{val: "11111"},
			out: testout{result: "0x1F"},
		},
	}

	for _, c := range cases {
		b, err := Convert(c.in.val)
		if err != nil {
			t.Error(err)
		}

		if c.out.result != b {
			t.Errorf("[want] %v\t[got] %v\n", c.out.result, b)
		}
	}
}
