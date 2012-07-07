package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
  CODE = `
yeq
ejp mysljylc kd kxveddknmc re jsicpdrysi
rbcpc ypc rtcsra dkh wyfrepkym veddknkmkrkcd
de kr kd eoya kw aej tysr re ujdr lkgc jv
z`

  KEY  = `
aoz
our language is impossible to understand
there are twenty six factorial possibilities
so it is okay if you want to just give up
q`
)


func main() {

  // Open the input, start Reader.
  datafile := os.Args[1]
  f, _ := os.Open(datafile)
  defer f.Close()
	r := bufio.NewReader(f)

  T := getInt(getline(r))

  // Loop through file, T lines.
  for X:=int64(0); X<T; X++ {
    G := getline(r)

    translator := func(r rune) rune {
      charPosition := strings.IndexRune(CODE, r)
      return rune(KEY[charPosition])
    }
    fmt.Printf("Case #%v: %v\n", X+1, strings.Map(translator, G))
  }
}

//------------------------------------------------------------------------------
// Helper functions.

func getInt(str string) (i int64) {
  i, _ = strconv.ParseInt(str, 10, 64)
  return
}

func getline(r *bufio.Reader) string {
	buffer := bytes.NewBuffer(make([]byte, 0))
  isPrefix := true;
  var part []byte
	for ;isPrefix; {
    part, isPrefix, _ = r.ReadLine()
		buffer.Write(part)
	}
	return buffer.String()
}

