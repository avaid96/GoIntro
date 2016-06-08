package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m:=make(map[string]int)
	for _,v := range strings.Fields(s) {
		//fmt.Println(v)
		_, exists:=m[v]
		//fmt.Println(word)
		if exists {
			m[v]=m[v]+1
		} else {
			m[v]=1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
