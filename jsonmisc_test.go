package jsonmisc_test

import (
	"io"
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

const (
	json  = "JSON_strings"
	ctrl  = "control_characters"
	jp    = "japanese"
	emoji = "emoji"
)

var testcases = map[string]fixture{
	json:  {`{"a":"https://github.com/"}`, `{\"a\":\"https://github.com/\"}`},
	ctrl:  {"START" + "\x00" + "\b" + "\f" + "\n" + "\r" + "\t" + "\x1f" + "\x20" + "END", `START\u0000\b\f\n\r\t\u001f END`},
	jp:    {"狂人の真似とて大路を走らば、即ち狂人なり。", "狂人の真似とて大路を走らば、即ち狂人なり。"},
	emoji: {"👍😭🙇‍♂️🙇🏻‍♂️🙇🏼‍♂️🙇🏽‍♂️🙇🏾‍♂️🙇🏿‍♂️👫👫🏻👫🏼👫🏽👫🏾👫🏿", "👍😭🙇‍♂️🙇🏻‍♂️🙇🏼‍♂️🙇🏽‍♂️🙇🏾‍♂️🙇🏿‍♂️👫👫🏻👫🏼👫🏽👫🏾👫🏿"},
}

func TestAppendQuote(t *testing.T) {
	SetEOL(LF)

	var key string

	key = json
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

	key = ctrl
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

	key = jp
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

	key = emoji
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
	var byteSlice []byte
	b.Run("", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			byteSlice = jsonmisc.AppendQuote(byteSlice, testcases[json].before)
			byteSlice = append(byteSlice, EOL()...)
		}
	})
	// nolint: errcheck
	io.Discard.Write(byteSlice)
}
