package chimps_io

import (
	"bufio"
	"io"
)

func ReadLine(f io.Reader, readLine func(content string)) {
	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		readLine(string(a))
	}
}
