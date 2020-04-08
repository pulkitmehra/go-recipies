package inout

import (
	"fmt"
	"io"
)

type (
	digitReader string

	digitReaderAdv struct {
		src string
		cur int
	}
)

func (d digitReader) Read(b []byte) (int, error) {
	count := 0
	for i := 0; i < len(d); i++ {
		if d[i] >= '0' || d[i] <= '9' {
			b[i] = d[i]
			count++
		}
	}
	return count, io.EOF
}

func (da *digitReaderAdv) Read(buf []byte) (int, error) {
	if da.cur >= len(da.src) {
		return 0, io.EOF
	}

	x := len(da.src) - da.cur
	n, bound := 0, 0
	if x >= len(buf) {
		bound = len(buf)
	} else if x <= len(buf) {
		bound = x
	}

	tBuf := make([]byte, bound)
	for n < bound {
		ch := da.src[da.cur]
		if ch >= '0' || ch <= '9' {
			tBuf[n] = ch
			n++
		}
		da.cur++
	}
	copy(buf, tBuf)
	return n, nil
}

//MainInOut exampl
func MainInOut() {
	da := &digitReaderAdv{src: "0123456789", cur: 0}
	buf := make([]byte, 5)

	//read 01234
	da.Read(buf)
	fmt.Println(string(buf))
	fmt.Println("----")
	//read 56789
	da.Read(buf)
	fmt.Println(string(buf))
}
