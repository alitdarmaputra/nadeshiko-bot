package love

import (
	"fmt"
	"strings"
)

type LoveService struct {
}

func NewLoveService() *LoveService {
	return &LoveService{}
}

func (l *LoveService) LoveCalc(name1 string, name2 string) (percent string) {
	var fullName = name1 + name2
	intFullNames := l.countChar(&fullName)
	intPercent := l.calcMatch(intFullNames)

	result := intPercent[0]*10 + intPercent[1]
	var react string

	if result > 80 {
		react = "You are a good partner"
	} else if result > 40 {
		react = "I think you should give a try"
	} else {
		react = "I think it's better just be a friend"
	}

	return fmt.Sprintf(
		"%s and %s is %d%d%% match, **%s**",
		name1,
		name2,
		intPercent[0],
		intPercent[1],
		react,
	)
}

func (l *LoveService) countChar(fullName *string) (intFullNames []int) {
	var count = map[string]int{}
	var keys []rune

	*fullName = strings.ToLower(*fullName)

	for _, e := range *fullName {
		_, isExist := count[string(e)]
		if isExist {
			count[string(e)]++
		} else {
			keys = append(keys, e)
			count[string(e)] = 1
		}
	}

	for _, key := range keys {
		intFullNames = append(intFullNames, count[string(key)])
	}
	return
}

func (l *LoveService) calcMatch(intFullNames []int) []int {
	if len(intFullNames) == 2 {
		return intFullNames
	} else {
		var subIntFullNames []int
		var left, right = 0, len(intFullNames) - 1

		for left <= right {
			if left == right {
				subIntFullNames = append(subIntFullNames, intFullNames[left]%10)
			} else {
				subIntFullNames = append(subIntFullNames, (intFullNames[left]+intFullNames[right])%10)
			}
			left++
			right--
		}

		return l.calcMatch(subIntFullNames)
	}
}
