package volume

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type MultipartWriter struct {
	uploadPart func(r io.Reader, off int64, id int) error
	complete   func() error
	cancel     func() error
	size       int64
	partSize   int

	off    int64 // for write only, writeAt does not change it
	blocks []bool
}

func (mw *MultipartWriter) init() {
	cnt := int(mw.size / int64(mw.partSize))
	if int64(cnt*mw.partSize) < mw.size {
		cnt++
	}
	mw.blocks = make([]bool, cnt)
}

func (mw *MultipartWriter) Write(p []byte) (n int, e error) {
	n, e = mw.WriteAt(p, mw.off)
	mw.off += int64(n)
	return
}

// WriteAt remote object
// for now, the offset must be multiple of part size,
// as well as the length of input bytes p, unless it's the last one
func (mw *MultipartWriter) WriteAt(p []byte, off int64) (n int, e error) {
	if off%int64(mw.partSize) != 0 {
		return 0, fmt.Errorf("%w: write at invalid offset, and it's not buffered, yet", io.ErrShortBuffer)
	}

	// may write more than one blocks at once
	for len(p[n:]) > 0 {
		var section []byte
		if len(p[n:]) > mw.partSize {
			section = p[n : n+mw.partSize]
		} else if off+int64(len(p)) == mw.size {
			// last block' size doesn't matter
			section = p[n:]
		} else {
			return n, errors.New("invalid part size to write at")
		}

		id := (int(off) + n) / mw.partSize
		e = mw.uploadPart(bytes.NewBuffer(section), off+int64(n), id)
		if e != nil {
			return
		}
		mw.blocks[id] = true
		n += len(section)
	}

	return
}
