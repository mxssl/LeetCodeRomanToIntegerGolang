package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "III"

	expected := 3

	output := romanToInt(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}

func TestCase2(t *testing.T) {
	input := "IV"

	expected := 4

	output := romanToInt(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}

func TestCase3(t *testing.T) {
	input := "IX"

	expected := 9

	output := romanToInt(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}

func TestCase4(t *testing.T) {
	input := "LVIII"

	expected := 58

	output := romanToInt(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}
