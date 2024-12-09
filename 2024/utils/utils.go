package utils

import (
	"os"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func FileLines(fName string) []string {
	data, _ := os.ReadFile(fName)
	split := strings.Split(string(data), "\n")
	return split[:len(split)-1]
}

func AllFn[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func CountFn[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, item := range slice {
		if predicate(item) {
			count++
		}
	}
	return count
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func UniqueSlice[T comparable](slice []T) []T {
	seen := make(map[T]struct{})
	unique := make([]T, 0)

	for _, item := range slice {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			unique = append(unique, item)
		}
	}

	return unique
}
