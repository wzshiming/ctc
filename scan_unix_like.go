package ctc

import (
	"bytes"
	"io"
)

func ScanUnixLike(r io.Reader) (Color, []byte, bool) {
	buf := [20]byte{}
	off := 0
	c := Color(0)
	for _, v := range pre {
		r.Read(buf[off : off+1])
		if buf[off] != v {
			return c, buf[:off+1], false
		}
		off++
	}
	code := byte(0)

	for {
		if off >= 16 {
			return c, buf[:off], false
		}
		r.Read(buf[off : off+1])
		switch buf[off] {
		default:
			return c, buf[:off+1], false
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			code *= 10
			code += buf[off] - '0'
		case 'm', ';':
			switch code {
			case 30, 31, 32, 33, 34, 35, 36, 37:
				c |= Color(code - 30)
			case 40, 41, 42, 43, 44, 45, 46, 47:
				c |= Color(code-40) << 4
			case 90, 91, 92, 93, 94, 95, 96, 97:
				c |= Color(code-90) | ForegroundBright
			case 100, 101, 102, 103, 104, 105, 106, 107:
				c |= Color(code-100)<<4 | BackgroundBright
			}
			if buf[off] == 'm' {
				return c, buf[:off+1], true
			}
		}
		off++
	}
}

func PipeUnixLike(w io.Writer) io.Writer {
	buf := newSyncBuffer()
	go func() {
		for {
			c, b, ok := ScanUnixLike(buf)
			if ok {
				c.Apply()
			} else {
				w.Write(b)
			}
		}
	}()
	return buf
}

type syncBuffer struct {
	buf bytes.Buffer
	s   chan struct{}
}

func newSyncBuffer() *syncBuffer {
	return &syncBuffer{
		s: make(chan struct{}),
	}
}

func (b *syncBuffer) Read(p []byte) (n int, err error) {
	n, err = b.buf.Read(p)
	if b.buf.Len() == 0 {
		<-b.s
	}
	return
}

func (b *syncBuffer) Write(p []byte) (n int, err error) {
	n, err = b.buf.Write(p)
	if b.buf.Len() != 0 {
		b.s <- struct{}{}
	}
	return
}
