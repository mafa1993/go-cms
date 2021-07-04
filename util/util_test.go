package util

import (
	"flag"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	// 模拟参数
	args := []string{
		"-conf=truabce",
	}

	flag.CommandLine.Parse(args)
	c := getConfig("F:\\learn\\golang\\src\\sixedu\\data\\config.json")
	fmt.Println(c)
}
