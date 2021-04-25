package jsonmisc_test

import (
	"os"
	"testing"

	jsonmisc "github.com/djeeno/go-json-escape"
)

// go test -cover -v

const (
	LineFeed       byte = 0x0A
	CarriageReturn byte = 0x0D
)

var (
	LF   = []byte{LineFeed}
	CR   = []byte{CarriageReturn}
	CRLF = []byte{LineFeed, CarriageReturn}
	eol  = LF
)

func EOL() []byte {
	return eol
}

func SetEOL(newEOL []byte) {
	// TODO: sync
	eol = newEOL
}

type fixture struct {
	before string
	expect string
}

var testcases = map[string]fixture{
	"JSON_strings":       {`{"a":"https://github.com/"}`, `{\"a\":\"https://github.com/\"}`},
	"control_characters": {"START" + "\x00" + "\b" + "\f" + "\n" + "\r" + "\t" + "\x1f" + "\x20" + "END", `START\u0000\b\f\n\r\t\u001f END`},
	"japanese":           {"ぁ震天裂空斬光旋風滅砕神罰割殺撃", "ぁ震天裂空斬光旋風滅砕神罰割殺撃"},
	"emoji":              {"😭🙇🏻‍♂️", "😭🙇🏻‍♂️"},
}

func TestAppendQuote(t *testing.T) {
	SetEOL(LF)

	var key string

	key = "JSON_strings"
	t.Run(key, func(t *testing.T) {
		var byteSlice []byte

		byteSlice = jsonmisc.AppendQuote(byteSlice, testcases[key].before)

		expect := testcases[key].expect
		actual := string(byteSlice)
		if expect != actual {
			t.Fail()
		}

		byteSlice = append(byteSlice, EOL()...)

		os.Stdout.Write(byteSlice)
	})

	key = "control_characters"
	t.Run(key, func(t *testing.T) {
		var byteSlice []byte

		byteSlice = jsonmisc.AppendQuote(byteSlice, testcases[key].before)

		expect := testcases[key].expect
		actual := string(byteSlice)
		if expect != actual {
			t.Fail()
		}

		byteSlice = append(byteSlice, EOL()...)

		os.Stdout.Write(byteSlice)
	})

	key = "japanese"
	t.Run(key, func(t *testing.T) {
		var byteSlice []byte

		byteSlice = jsonmisc.AppendQuote(byteSlice, testcases[key].before)

		expect := testcases[key].expect
		actual := string(byteSlice)
		if expect != actual {
			t.Fail()
		}

		byteSlice = append(byteSlice, EOL()...)

		os.Stdout.Write(byteSlice)
	})

	key = "emoji"
	t.Run(key, func(t *testing.T) {
		var byteSlice []byte

		byteSlice = jsonmisc.AppendQuote(byteSlice, testcases[key].before)

		expect := testcases[key].expect
		actual := string(byteSlice)
		if expect != actual {
			t.Fail()
		}

		byteSlice = append(byteSlice, EOL()...)

		os.Stdout.Write(byteSlice)
	})
}

// go test -bench . -benchmem -test.run=none -test.benchtime=1000ms

func Benchmark(b *testing.B) {
	b.Run("", func(b *testing.B) {
		var byteSlice []byte
		for i := 0; i < b.N; i++ {
			byteSlice = jsonmisc.AppendQuote(byteSlice, testcases["json"].before)
			byteSlice = append(byteSlice, EOL()...)
		}

		os.Stdout.Write(byteSlice)
	})
}
