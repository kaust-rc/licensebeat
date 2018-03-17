[![Build Status](https://travis-ci.org/kaust-rc/licensebeat.svg?branch=master)](https://travis-ci.org/kaust-rclicensebeat)
[![codecov.io](http://codecov.io/github/kaust-rc/licensebeat/coverage.svg?branch=master)](http://codecov.io/github/kaut-rc/licensebeat?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kaust-rc/licensebeat)](https://goreportcard.com/report/github.com/kaust-rc/licensebeat)
[![license](https://img.shields.io/github/license/kaust-rc/licensebeat.svg)](https://github.com/kaust-rc/licensebeat)
[![Github All Releases](https://img.shields.io/github/downloads/kaust-rc/licensebeat/total.svg)](https://github.com/kaust-rc/licensebeat)

![Elastic Beats 6.2](https://img.shields.io/badge/Elastic%20Beats-v6.2-blue.svg)
![Golang 1.9](https://img.shields.io/badge/Golang-v1.9-blue.svg)


# Licensebeat

Welcome to Licensebeat.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/kaust-rc/licensebeat`

## Getting Started with Licensebeat

### Requirements

* [Golang](https://golang.org/dl/) 1.9

### Init Project
To get running with Licensebeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Licensebeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/kaust-rc/licensebeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Licensebeat run the command below. This will generate a binary
in the same directory with the name licensebeat.

```
make
```


### Run

To run Licensebeat with debugging output enabled, run:

```
./licensebeat -c licensebeat.yml -e -d "*"
```


### Test

To test Licensebeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Licensebeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Licensebeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/kaust-rc/licensebeat
git clone https://github.com/kaust-rc/licensebeat ${GOPATH}/src/github.com/kaust-rc/licensebeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
