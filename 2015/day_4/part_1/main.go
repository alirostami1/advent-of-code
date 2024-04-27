package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	input := scanner.Text()

	i := 1
	for {
		data := fmt.Sprintf("%s%d", input, i)

		hash := md5.Sum([]byte(data))
		hex := hex.EncodeToString(hash[:])

		for j := 0; j < 5; j++ {
			if hex[j] != '0' {
				break
			}
			if j == 4 {
				fmt.Print(data)
				return
			}
		}
		i++
	}
}
