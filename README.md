[![Test Status](https://github.com/davidji99/terraform-provider-rollbar/workflows/tests/badge.svg)](https://github.com/davidji99/terraform-provider-rollbar/actions?query=workflow%3Atests)

Terraform Provider Rollbar
=========================

This provider is used to configure certain resources supported by [Rollbar API](https://docs.rollbar.com/reference).

**NOTE**: This provider is unofficial and not created by the Rollbar team.
If you have questions about Rollbar functionality, please kindly refer to the [Rollbar API documentation](https://explorer.docs.rollbar.com/).

For provider bugs/questions, please open an issue on this repository.

Documentation
------------

Documentation about resources and data sources can be found [here](https://registry.terraform.io/providers/davidji99/rollbar/latest/docs).

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) `v0.12+`. (No support for terraform `v0.11`)
- [Go](https://golang.org/doc/install) `v1.14` (to build the provider plugin)

Usage
-----

For Terraform `v0.12+` compatibility, the configuration should specify version 1.0.0 or higher:
```hcl-terraform
provider "rollbar" {
  version = ">= 1.0.0"
}
```
- This requires the provider binary installed somewhere on the host machine.


For Terraform `v0.13+` compatibility, the following configuration should be used:
```hcl-terraform
terraform {
  required_providers {
    rollbar = {
      source = "davidji99/rollbar"
      version = "1.0.0"
    }
  }
}
```

Releases
------------

Provider binaries can be found [here](https://github.com/davidji99/terraform-provider-rollbar/releases).

Development
-----------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.12+ is *required*).

If you wish to bump the provider version, you can do so in the file `version/version.go`.

### Clone the Provider

This repository supports Go modules so you can clone this repository anywhere you wish and does not have to be in `$GOPATH`.

### Build the Provider

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-rollbar
...
```

### Using the Provider

To use the dev provider with local Terraform, copy the freshly built plugin into Terraform's local plugins directory:

```sh
cp $GOPATH/bin/terraform-provider-rollbar ~/.terraform.d/plugins/
```

Set the Rollbar provider without a version constraint:

```hcl
provider "rollbar" {}
```

Then, initialize Terraform:

```sh
terraform init
```

### Testing

Please see the [TESTING](TESTING.md) guide for detailed instructions on running tests.

### Updating or adding dependencies

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.

Dependencies can be added or updated as follows:

```bash
$ GO111MODULE=on go get github.com/some/module@release-tag
$ GO111MODULE=on go mod tidy
$ GO111MODULE=on go mod vendor
```

This example will fetch a module at the release tag and record it in your project's go.mod and go.sum files.
It's a good idea to tidy up afterward and then copy the dependencies into `vendor/` directory.

If a module does not have release tags, then `module@master` can be used instead.

#### Removing dependencies

Remove all usage from your codebase and run:

```bash
$ GO111MODULE=on go mod tidy
$ GO111MODULE=on go mod vendor
```
