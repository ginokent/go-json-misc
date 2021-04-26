# go-json-misc

Convert string to a escaped JSON string.

```go
b := []byte(`{"key":"`)

b = jsonmisc.AppendQuote(b, `"\"The string to be escaped.`+"\x00"+"\b"+"\f"+"\n"+"\r"+"\t"+"\x1f"+"\x20"+`\""`)

b = append(b, []byte(`"}`)...)

os.Stdout.Write(b)
// -> {"key":"\"\\\"The string to be escaped.\u0000\b\f\n\r\t\u001f \\\"\""}
```
