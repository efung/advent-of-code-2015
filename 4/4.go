package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}
	secret := os.Args[1]
	prefix_len, err := strconv.Atoi(os.Args[2])
	if err != nil {
		os.Exit(2)
	}

	MD5 := md5.New()
	for i := 0; ; i++ {
		MD5.Reset()
		input := secret + strconv.Itoa(i)
		MD5.Write([]byte(input))
		digest := MD5.Sum(nil)

		if strings.HasPrefix(hex.EncodeToString(digest), strings.Repeat("0", prefix_len)) {
			fmt.Printf("\n%d\n", i)
			break
		}
		if i%100000 == 0 {
			fmt.Print(".")
		}
	}
}
