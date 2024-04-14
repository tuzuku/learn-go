package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type MyFloat float64
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	return math.Sqrt(x), nil
}

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (reader *MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	default:
		return b
	}
}

func (rr *rot13Reader) Read(p []byte) (int, error) {
	n, err := rr.r.Read(p)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, nil

}

func main() {
	if res, err := Sqrt(4); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	if res1, err2 := Sqrt(-2); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(res1)
	}

	r := strings.NewReader("Hello, Reader!")

	buffer := make([]byte, 8)

	for {
		n, err := r.Read(buffer)

		fmt.Printf("n = %v err = %v b = %v\n", n, err, buffer)
		fmt.Printf("b[:n] = %q\n", buffer[:n])
		if err == io.EOF {
			break
		}
	}

	nb := MyReader{}
	n, err := nb.Read(buffer)

	fmt.Printf("n = %v err = %v b = %v\n", n, err, buffer)
	fmt.Printf("b[:n] = %q\n", buffer[:n])
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rr := rot13Reader{s}
	io.Copy(os.Stdout, &rr)
}
