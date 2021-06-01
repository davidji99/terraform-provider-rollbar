package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// getTeamID extracts the team ID attribute generically from a Rollbar resource.
func getTeamID(d *schema.ResourceData) int {
	var teamID int
	if v, ok := d.GetOk("team_id"); ok {
		vs := v.(int)
		log.Printf("[DEBUG] team_id: %d", vs)
		teamID = vs
	}

	return teamID
}

// getProjectID extracts the project ID attribute generically from a Rollbar resource.
func getProjectID(d *schema.ResourceData) int {
	var projectID int
	if v, ok := d.GetOk("project_id"); ok {
		vs := v.(int)
		log.Printf("[DEBUG] Project id: %d", vs)
		projectID = vs
	}

	return projectID
}

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

// ParseCompositeImportID takes a composite id separated by a colon and returns two string values.
func ParseCompositeImportID(id string) (p1 string, p2 string, err error) {
	parts := strings.Split(id, ":")
	if len(parts) == 2 {
		p1 = parts[0]
		p2 = parts[1]
	} else {
		err = fmt.Errorf("error: Import composite ID requires two parts separated by colon, eg x:y")
	}
	return
}

// ParseCompositeID splits a given string based on a specified number of
func ParseCompositeID(id string, numOfSplits int) ([]string, error) {
	parts := strings.SplitN(id, ":", numOfSplits)

	if len(parts) != numOfSplits {
		return nil, fmt.Errorf("error: composite ID requires %d parts separated by a colon (eg x:y)", numOfSplits)
	}
	return parts, nil
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
