package test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

const (
	TestConfigAccountAccessToken TestConfigKey = iota
	TestConfigProjectAccessToken
	TestConfigAcceptanceTestKey
	TestConfigPagerDutyAPIKey
	TestConfigUserEmailKey
	TestConfigTeamID
	TestConfigEmailAddress
)

var testConfigKeyToEnvName = map[TestConfigKey]string{
	TestConfigAccountAccessToken: "ROLLBAR_ACCOUNT_ACCESS_TOKEN",
	TestConfigProjectAccessToken: "ROLLBAR_PROJECT_ACCESS_TOKEN",
	TestConfigPagerDutyAPIKey:    "ROLLBAR_PD_API_KEY",
	TestConfigUserEmailKey:       "ROLLBAR_USER_EMAIL",
	TestConfigTeamID:             "ROLLBAR_TEAM_ID",
	TestConfigEmailAddress:       "ROLLBAR_EMAIL_ADDRESS",
	TestConfigAcceptanceTestKey:  resource.TestEnvVar,
}

type TestConfigKey int

type TestConfig struct{}

func NewTestConfig() *TestConfig {
	return &TestConfig{}
}

func (k TestConfigKey) String() (name string) {
	if val, ok := testConfigKeyToEnvName[k]; ok {
		name = val
	}
	return
}

func (t *TestConfig) Get(keys ...TestConfigKey) (val string) {
	for _, key := range keys {
		val = os.Getenv(key.String())
		if val != "" {
			break
		}
	}
	return
}

func (t *TestConfig) GetOrSkip(testing *testing.T, keys ...TestConfigKey) (val string) {
	t.SkipUnlessAccTest(testing)
	val = t.Get(keys...)
	if val == "" {
		testing.Skip(fmt.Sprintf("skipping test: config %v not set", keys))
	}
	return
}

func (t *TestConfig) GetOrAbort(testing *testing.T, keys ...TestConfigKey) (val string) {
	t.SkipUnlessAccTest(testing)
	val = t.Get(keys...)
	if val == "" {
		testing.Fatal(fmt.Sprintf("stopping test: config %v must be set", keys))
	}
	return
}

func (t *TestConfig) SkipUnlessAccTest(testing *testing.T) {
	val := t.Get(TestConfigAcceptanceTestKey)
	if val == "" {
		testing.Skip(fmt.Sprintf("Acceptance tests skipped unless env '%s' set", TestConfigAcceptanceTestKey.String()))
	}
}

func (t *TestConfig) GetUserEmailOrAbort(testing *testing.T) (val string) {
	return t.GetOrAbort(testing, TestConfigUserEmailKey)
}

func (t *TestConfig) GetPagerDutyAPIKeyorAbort(testing *testing.T) (val string) {
	return t.GetOrAbort(testing, TestConfigPagerDutyAPIKey)
}

func (t *TestConfig) GetTeamIDorAbort(testing *testing.T) (val string) {
	return t.GetOrAbort(testing, TestConfigTeamID)
}

func (t *TestConfig) GetTeamEmailAddress(testing *testing.T) (val string) {
	return t.GetOrAbort(testing, TestConfigEmailAddress)
}
