package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)

	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(counts)

	file, err := os.Open("default.log")
	if err != nil {
		fmt.Println(err)
	}

	counts, err = countLetters(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(counts)

	r, closer, err := buildGZipReader("default.log.gz")
	if err != nil {
		fmt.Println(err)
	}
	defer closer()
	counts, err = countLetters(r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(counts)

	b := NopCloser(sr)
	some(b)

	[...data..]<read> ===> {buffer} ===> [...data...]<write>
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func NopCloser(r io.Reader) io.ReadCloser {
	return nopCloser{r}
}

func some(a io.ReadCloser) {

}
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		err := gr.Close()
		if err != nil {
			return
		}
		err = r.Close()
		if err != nil {
			return
		}
	}, nil
}
