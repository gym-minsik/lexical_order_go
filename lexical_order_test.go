package lexical_order

import (
	"sort"
	"testing"
)

var service = NewAlphabetOrderService()

func TestBetween(t *testing.T) {

	prev := "ABCdE"
	next := "ABChi"
	mid := service.Between(prev, next)

	if mid != "ABCf" {
		t.Errorf("Expected ABCf but %s", mid)
	}

	// Check lexicographical order: prev < result < next
	strs := []string{prev, mid, next}
	sorted := append([]string{}, strs...) // copy
	sort.Strings(sorted)
	for i := range strs {
		if strs[i] != sorted[i] {
			t.Errorf("Expected order: %v, got: %v", sorted, strs)
			break
		}
	}
}

func TestNotPositiveKeyCount(t *testing.T) {
	for i := 0; i >= -10; i-- {
		result := service.Generate(i)
		expected := []string{}

		if len(result) != len(expected) {
			t.Errorf("For keyCount = %d, expected empty list, got %v", i, result)
		}
	}
}

func TestGeneratedKeysShouldBeSorted(t *testing.T) {
	// Test with smaller keyCounts
	for keyCount := 1; keyCount <= 2600; keyCount++ {
		result := service.Generate(keyCount)
		copyAndSorted := append([]string{}, result...)
		sort.Strings(copyAndSorted)

		if len(result) != keyCount {
			t.Errorf("For keyCount = %d, expected %d keys, got %d", keyCount, keyCount, len(result))
		}

		// Ensure all generated keys are non-empty
		for _, key := range result {
			if key == "" {
				t.Errorf("Generated an empty key for keyCount = %d", keyCount)
			}
		}

		// Ensure the result is sorted lexicographically
		for i := 1; i < len(result); i++ {
			if result[i-1] > result[i] {
				t.Errorf("Generated keys are not sorted at keyCount = %d", keyCount)
				break
			}
		}
	}

	// Test with larger keyCounts
	bigKeyCounts := []int{
		10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000,
	}

	for _, keyCount := range bigKeyCounts {
		result :=
			service.Generate(keyCount)
		copyAndSorted := append([]string{}, result...)
		sort.Strings(copyAndSorted)

		if len(result) != keyCount {
			t.Errorf("For keyCount = %d, expected %d keys, got %d", keyCount, keyCount, len(result))
		}

		// Ensure all generated keys are non-empty
		for _, key := range result {
			if key == "" {
				t.Errorf("Generated an empty key for keyCount = %d", keyCount)
			}
		}

		// Ensure the result is sorted lexicographically
		for i := 1; i < len(result); i++ {
			if result[i-1] > result[i] {
				t.Errorf("Generated keys are not sorted at keyCount = %d", keyCount)
				break
			}
		}
	}
}
