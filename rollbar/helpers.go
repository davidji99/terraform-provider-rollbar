package rollbar

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// StringToInt converts a string parameter to an integer.
func StringToInt(s string) int {
	intValue, _ := strconv.Atoi(s)
	return intValue
}

// Int64ToString converts an int64 parameter to a string.
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// GenerateRandomResourceID returns a random number based on the epoch time in nanoseconds.
func GenerateRandomResourceID() string {
	rand.Seed(int64(time.Now().Nanosecond()))
	return strconv.Itoa(rand.Int())
}

// ParseCompositeID takes a composite id separated by a colon and returns two string values.
func ParseCompositeID(id string) (p1 string, p2 string, err error) {
	parts := strings.Split(id, ":")
	if len(parts) == 2 {
		p1 = parts[0]
		p2 = parts[1]
	} else {
		err = fmt.Errorf("error: Import composite ID requires two parts separated by colon, eg x:y")
	}
	return
}

// Contains takes a string slice and checks if another string value is in it.
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// DoesNotContain does the exact opposite of Contains.
func DoesNotContain(s []string, e string) bool {
	return !Contains(s, e)
}
