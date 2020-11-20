package perfect

import "errors"

type Classification int

const (
	ClassificationDeficient = iota

	ClassificationPerfect

	ClassificationAbundant
)

var ErrOnlyPositive error = errors.New("non positive number")

func Classify(number int64) (Classification, error) {
	if number <= 0 {
		return ClassificationDeficient, ErrOnlyPositive
	}

	var i, sum int64 = 1, 0
	for ; i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}

	if sum < number {
		return ClassificationDeficient, nil
	} else if sum == number {
		return ClassificationPerfect, nil
	} else {
		return ClassificationAbundant, nil
	}
}
