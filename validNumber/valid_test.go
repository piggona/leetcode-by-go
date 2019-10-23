package main

import (
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	t.Run("0", func(t *testing.T) { fmt.Println(isNumber("0")) })
	t.Run(" 0.1", func(t *testing.T) { fmt.Println(isNumber(" 0.1")) })
	t.Run("abc", func(t *testing.T) { fmt.Println(isNumber("abc")) })
	t.Run("1 a", func(t *testing.T) { fmt.Println(isNumber("1 a")) })
	t.Run("2e10", func(t *testing.T) { fmt.Println(isNumber("2e10")) })
	t.Run(" -90e3   ", func(t *testing.T) { fmt.Println(isNumber(" -90e3   ")) })
	t.Run(" 1e", func(t *testing.T) { fmt.Println(isNumber(" 1e")) })
	t.Run("e3", func(t *testing.T) { fmt.Println(isNumber("e3")) })
	t.Run("99e2.5", func(t *testing.T) { fmt.Println(isNumber("99e2.5")) })
}
