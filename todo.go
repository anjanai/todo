package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func main() {
	data, _ := ioutil.ReadFile("todo.txt")
	s := strings.Fields(string(data))

	rand.Seed(time.Now().Unix())

	do := s[rand.Intn(len(s))]
	fmt.Print(do)
}
