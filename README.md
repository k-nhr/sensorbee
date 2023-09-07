[![Build Status](https://travis-ci.org/sensorbee/sensorbee.svg?branch=master)](https://travis-ci.org/sensorbee/sensorbee)
[![Coverage Status](https://coveralls.io/repos/github/sensorbee/sensorbee/badge.svg?branch=master)](https://coveralls.io/github/sensorbee/sensorbee?branch=master)

**This project has been discontinued.**

# SensorBee: Lightweight stream processing engine for IoT

### Install SensorBee
```
$ go get gopkg.in/sensorbee/sensorbee.v0/...
```
`sensorbee` and `build_sensorbee` commands are installed under $GOPATH/bin.

### Run SensorBee
```
$ sensorbee run
```

### Build Custom SensorBee Command
Because SensorBee is written in Go, all plugins are statically linked. In order to add a plugin, a custom `sensorbee` command must be built.
```
$ ls
build.yaml

$ cat build.yaml
plugins:
  - github.com/sensorbee/twitter/plugin

$ build_sensorbee
sensorbee_main.go

$ ls
build.yaml
sensorbee
sensorbee_main.go
```

See http://docs.sensorbee.io/en/latest/ for details.
