package lexical_order

import (
	"math"
	"strings"
)

func removeTrailings(str, trailingCharacter string) string {
	return strings.TrimRight(str, trailingCharacter)
}

func generateExtraSteps(keys []string, count int) []string {
	count %= len(keys)
	result := []string{}
	for i := 1; i <= count; i++ {
		index := int((float64(len(keys)) / float64(count+1)) * float64(i))
		result = append(result, keys[RoundToNearestEven(float64(index))])
	}
	return result
}

func generateFoundation(keys []string, length int) []string {
	var result []string
	var dfs func(n int, str string)
	dfs = func(n int, str string) {
		if n > 0 {
			for _, k := range keys {
				dfs(n-1, str+k)
			}
		} else {
			result = append(result, str)
		}
	}
	dfs(length, "")
	return result
}

func generate(keys []string, count int) []string {
	if count <= 0 {
		return nil
	}
	minLength := int(math.Floor(LogBase(float64(count), float64(len(keys)))))
	averageExtraSteps := float64(count+1)/math.Pow(float64(len(keys)), float64(minLength)) - 1.0
	stepWeight := averageExtraSteps - math.Floor(averageExtraSteps)
	stepsSelector := 0.5 + stepWeight

	a := int(averageExtraSteps)
	extraStepsList := [][]string{
		generateExtraSteps(keys, a),
		generateExtraSteps(keys, a+1),
	}

	var result []string
	for _, f := range generateFoundation(keys, minLength) {
		result = append(result, removeTrailings(f, keys[0]))
		for _, step := range extraStepsList[int(stepsSelector)] {
			result = append(result, f+step)
		}
		stepsSelector = math.Mod(stepsSelector, 1)
		stepsSelector += stepWeight
		// stepsSelector = math.Mod(stepsSelector+stepWeight, 1)
	}
	return result
}

func GenerateOrderKeys(keys []string, count int) []string {
	if count <= 0 {
		return nil
	}
	if count < len(keys) {
		return generateExtraSteps(keys, count)
	}
	return generate(keys, count)[1:]
}
