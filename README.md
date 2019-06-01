# paperless-cli
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/stgarf/paperless-cli/master/LICENSE)
[![Build Status](https://travis-ci.org/stgarf/paperless-cli.svg?branch=master)](https://travis-ci.org/stgarf/paperless-cli)

# Table of Contents

- [Description](#description)
- [Installation and Usage](#Installation-and-usage)
  * [Prerequisites](#prerequisites)
  * [Precompiled binary](#precompiled-binary)
  * [Self-compiled binary](#self-compiled-binary)
  * [Usage](#usage)
    - [Check paperless-cli version](#check-paperless-cli-version)
    - [Setting up a config](#setting-up-a-config)
- [Development](#development)
  * [Prerequisites](#prerequisites)
  * [Get the code](#get-the-code)
  * [Running the tests](#running-the-tests)
- [Built With](#built-with)
- [Contributing](#contributing)
- [Versioning](#versioning)
- [Authors](#authors)
- [License](#license)
- [Acknowledgments](#acknowledgments)


## Description

A CLI tool written in Go to interface with a [Paperless](https://github.com/the-paperless-project/paperless) instance.

## Installation and usage

### Prerequisites

You should have a working Go environment and have `$GOPATH/bin` in your `$PATH`.

### Precompiled binary

You can get a precompiled binary from the releases page (soon).

### Self-compiled binary

To download source, compile, and install, run:
```shell
$ go get -u github.com/stgarf/paperless-cli
$ $GOPATH/src/github.com/stgarf/paperless-cli/build.sh
```

The source code will be located in `$GOPATH/src/github.com/stgarf/paperless-cli`.

A newly compiled binary will be in `$GOPATH/src/github.com/stgarf/paperless-cli/bin/`

You can place the binary in you `$PATH` for easy usage. e.g.

`$ cp $GOPATH/src/github.com/stgarf/paperless-cli/bin/paperless-cli $GOPATH/bin`

`$ which paperless-cli` should return the path to the newly installed binary.

### Usage

#### Check paperless-cli version

```shell
$ paperless-cli version
paperless-cli v0.6.0 built on 2019-03-10T16:28:33Z from git:c8dd2e7-clean (master) by user@chrx
```

#### Setting up a config

You can set up a basic YAML-based configuration to be read by placing it in
`$HOME/.paperless-cli.yaml`.

A base config can be autogenerated with `paperless-cli config create`.

Here's an example configuration:
```yaml
# A basic paperless-cli configuration file.

# The hostname of the Paperless instance.
hostname: localhost
# Connect via HTTP or HTTPS.
use_https: false
# Port the Paperless instance is listening on.
port: 8000
# Path to the API root.
root: /api
# Username to auth to Paperless instance.
username: myUsername
# Username to auth to Paperless instance.
password: mYpa$$w0rd.is.h0p3fully.s3cur3!!1
```
## Development

### Prerequisites

You should have a working Go environment and have `$GOPATH/bin` in your `$PATH`.

### Get the code

To download the source code, run:
```shell
$ go get -du github.com/stgarf/paperless-cli
```

The source code will be located in `$GOPATH/src/github.com/stgarf/paperless-cli`.

### Running the tests

`$ go test`

Really though, there are none and need to be written.

## Built With

* [cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
* [logrus](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go
* [gjson](https://github.com/tidwall/gjson) - Get JSON values quickly - JSON Parser for Go
* [viper](https://github.com/spf13/viper) - Go configuration with fangs
* [go-fqdn](https://github.com/Showmax/go-fqdn) - Simple wrapper around net and os golang standard libraries providing Fully Qualified Domain Name of the machine
* [go-homedir](https://github.com/mitchellh/go-homedir) - Go library for detecting and expanding the user's home directory without cgo
* [govvv](https://github.com/ahmetb/govvv) - "go build" wrapper to add version info to Golang applications

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on the code of conduct, and the process for submitting pull requests to the project.

## Versioning

We use [SemVer](https://semver.org/) for versioning, roughly. For the versions available, see the [tags on this repository](https://github.com/stgarf/paperless-cli/tags).

## Authors

* **Steve Garf** - *Initial CLI work* - [stgarf](https://github.com/stgarf)

See also the list of [contributors](https://github.com/stgarf/paperless-cli/contributors) who participated in this project.

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

* [Daniel Quinn](https://github.com/danielquinn) for starting Paperless and the awesome community of people maintaining [Paperless](https://github.com/the-paperless-project/paperless)
* Seth Junot for rubber duck debugging - [xSetech](https://github.com/xSetech)
* Hat tip to anyone whose code was imported!
