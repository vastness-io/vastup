package util

import (
	"math"
	"time"

	"github.com/PI-Victor/vastup/pkg/up"
)

const (
	errorMaxCap = 10
)

type execFunc func(up.BuildContext) error

// RetryOnFailure implements an exponential backoff algorithm that will
// retry a function till hitting errorMaxCap.
func RetryOnFailure(ctx up.BuildContext, f execFunc) (bool, []error) {

	var errors []error

	for i := 0; i <= errorMaxCap; i++ {
		var (
			t    = time.Duration(math.Pow(2, float64(i)))
			wait = time.Millisecond * t
		)

		if err := f(ctx); err != nil {
			errors = aggregateErrors(errors, err)
			time.Sleep(wait)
			continue
		}
		return true, errors
	}
	return false, errors
}

func aggregateErrors(errors []error, err error) []error {
	if errors == nil {
		return append([]error{}, err)
	}

	return append(
		errors,
		err,
	)
}
