# Testing

## Provider Tests
In order to test the provider, you can simply run `make test`.

```bash
$ make test
```

## Acceptance Tests

You can run the complete suite of Rollbar acceptance tests by doing the following:

```bash
$ make testacc TEST="./rollbar/" 2>&1 | tee test.log
```

To run a single acceptance test in isolation replace the last line above with:

```bash
$ make testacc TEST="./rollbar/" TESTARGS='-run=TestAccRollbarProjects_Basic'
```

A set of tests can be selected by passing `TESTARGS` a substring. For example, to run all Rollbar project tests:

```bash
$ make testacc TEST="./rollbar/" TESTARGS='-run=RollbarProject'
```

### Test Parameters

The following parameters are available for running the test. The absence of some of the non-required parameters will cause certain tests to be skipped.

* **TF_ACC** (`integer`) **Required** - must be set to `1`.
* **ROLLBAR_ACCOUNT_ACCESS_TOKEN** (`string`) - The account access token of the user running the test.
* **ROLLBAR_PROJECT_ACCESS_TOKEN** (`string`) - The project access token of the user running the test.
* **ROLLBAR_PD_API_KEY** (`string`) - A PagerDuty API key.
* **ROLLBAR_USER_EMAIL** (`string`) - A Rollbar user email address.
* **ROLLBAR_TEAM_ID** (`string`) - The ID of a rollbar team.
* **ROLLBAR_EMAIL_ADDRESS** (`string`) - An email address.

Please note: if you run the entire acceptance suite, you will need to set BOTH `ROLLBAR_ACCOUNT_ACCESS_TOKEN` & `ROLLBAR_PROJECT_ACCESS_TOKEN`.
Otherwise, certain tests require either token.

**For example:**
```bash
export TF_ACC=...
export ROLLBAR_ACCOUNT_ACCESS_TOKEN=...
export ROLLBAR_PROJECT_ACCESS_TOKEN=...
$ make testacc TEST="./rollbar/" 2>&1 | tee test.log
```
