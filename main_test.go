package main

import (
	"reflect"
	"testing"
)

func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Error("Expected IsEven(2) to be true")
	}
	if IsEven(3) {
		t.Error("Expected IsEven(3) to be false")
	}
}

func TestSumOfDigits(t *testing.T) {
	if SumOfDigits(123) != 6 {
		t.Errorf("Expected SumOfDigits(123) to be 6")
	}
	if SumOfDigits(456) != 15 {
		t.Errorf("Expected SumOfDigits(456) to be 15")
	}
}

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("racecar") {
		t.Error("Expected IsPalindrome(\"racecar\") to be true")
	}
	if IsPalindrome("hello") {
		t.Error("Expected IsPalindrome(\"hello\") to be false")
	}
}

func TestFactorial(t *testing.T) {
	if Factorial(5) != 120 {
		t.Errorf("Expected Factorial(5) to be 120")
	}
	if Factorial(0) != 1 {
		t.Errorf("Expected Factorial(0) to be 1")
	}
}

func TestFizzBuzz(t *testing.T) {
	if FizzBuzz(3) != "Fizz" {
		t.Errorf("Expected FizzBuzz(3) to be \"Fizz\"")
	}
	if FizzBuzz(5) != "Buzz" {
		t.Errorf("Expected FizzBuzz(5) to be \"Buzz\"")
	}
	if FizzBuzz(15) != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz(15) to be \"FizzBuzz\"")
	}
	if FizzBuzz(7) != "7" {
		t.Errorf("Expected FizzBuzz(7) to be \"7\"")
	}
}

func TestReverseString(t *testing.T) {
	if ReverseString("hello") != "olleh" {
		t.Errorf("Expected ReverseString(\"hello\") to be \"olleh\"")
	}
	if ReverseString("Go") != "oG" {
		t.Errorf("Expected ReverseString(\"Go\") to be \"oG\"")
	}
}

func TestCountVowels(t *testing.T) {
	if CountVowels("hello") != 2 {
		t.Errorf("Expected CountVowels(\"hello\") to be 2")
	}
	if CountVowels("AEIOU") != 5 {
		t.Errorf("Expected CountVowels(\"AEIOU\") to be 5")
	}
}

func TestIsPrime(t *testing.T) {
	if !IsPrime(17) {
		t.Error("Expected IsPrime(17) to be true")
	}
	if IsPrime(4) {
		t.Error("Expected IsPrime(4) to be false")
	}
}

func TestFindMax(t *testing.T) {
	numbers := []int{1, 5, 3, 9, 2}
	if FindMax(numbers) != 9 {
		t.Errorf("Expected FindMax(%v) to be 9", numbers)
	}
	if FindMax([]int{}) != 0 {
		t.Error("Expected FindMax of empty slice to be 0")
	}
}

func TestGenerateFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	result := GenerateFibonacci(10)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected GenerateFibonacci(10) to be %v, got %v", expected, result)
	}
}
