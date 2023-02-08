package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

const sessionId = "B4AE798F-1D40-459A-A3B4-46EBB88D2235"

func main() {

}

func getSign(uid string) string {
	t := "575205$" + sessionId + "$1674389396$rVX9ITrrTPrCurUe"
	return mdf(t)
}

func mdf(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}
