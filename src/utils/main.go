// Package slices implements various slice algorithms.
package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// Map turns a []TInput to a []TOutput using a mapping function.
// This works with slices of any type.
func Map[TInput, TOutput any](s []TInput, f func(element TInput, index int32) TOutput) []TOutput {
	var index int32 = 0
	mappedData := make([]TOutput, len(s))

	for i, element := range s {
		mappedData[i] = f(element, index)
		index++
	}

	return mappedData
}

func ForEach[TInput any](data []TInput, f func(element TInput) error) error {
	var err error

	for _, element := range data {
		if err == nil {
			err = f(element)
		}
	}

	return err
}

// Reduces a []TInput to a single value of type TOutput, by applying a reduction function.
func Reduce[TInput, TOutput any](s []TInput, f func(accumulator TOutput, element TInput, index int32) TOutput, initialValue TOutput) TOutput {
	var index int32 = 0
	reducedValue := initialValue

	for _, element := range s {
		reducedValue = f(reducedValue, element, index)
		index++
	}

	return reducedValue
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
func Filter[T any](s []T, f func(element T, index int32) bool) []T {
	var index int32 = 0
	var filteredData []T

	for _, element := range s {
		if f(element, index) {
			filteredData = append(filteredData, element)
		}

		index++
	}

	return filteredData
}

func First[T any](results []T, err error) (*T, error) {
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return &results[0], nil
}

func Find[T any](s []T, f func(element T, index int32) bool) *T {
	filtered := Filter[T](s, f)
	firstElement, _ := First[T](filtered, nil)

	return firstElement
}

func Singlify[T any, U any](cb func(inputArray []T) ([]U, error)) func(singleInput T) (*U, error) {
	return func(singleInput T) (*U, error) {
		results, err := cb([]T{singleInput})

		if err != nil {
			return nil, err
		}

		if len(results) == 0 {
			return nil, nil
		}

		return &results[0], nil
	}
}

// NOTE: this isn't multi-Unicode-codepoint aware, like specifying skintone or
// gender of an emoji: https://unicode.org/emoji/charts/full-emoji-modifiers.html
//
// Taken from: https://stackoverflow.com/questions/12311033/extracting-substrings-in-go/56129336#56129336
func Substring(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func EncodeBase64() {

}

func DecodeBase64(base64Text string) (string, error) {
	output := make([]byte, base64.StdEncoding.DecodedLen(len(base64Text)))

	byteCount, err := base64.StdEncoding.Decode(output, []byte(base64Text))
	if err != nil {
		return "", err
	}

	return string(output[:byteCount]), nil
}

// UTF-8 Safe
func CountOccurrences(text string, substring string) int {
	return strings.Count(text, substring)
}

func CreateUuid() pgtype.UUID {
	newUuid := uuid.New()
	pgUUID := pgtype.UUID{}
	pgUUID.Bytes = newUuid

	return pgUUID
}

func CreateUuidV2() string {
	return uuid.New().String()
}

func As[T any](data any) T {
	var castedValue T = data.(T)

	return castedValue
}

func CreateRandomString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
