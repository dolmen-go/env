
# env - The POSIX env utility implemented in Go

See the [IEEE Std 1003.1-2017 specification](https://pubs.opengroup.org/onlinepubs/9699919799/).

## Synopsis

```
env [-i] [name=value ...] [utility [argument ...]]
```

## Install

You probably don't want to install this program because Unix-like environments already have an implementation, so this is redundant.

However you might want to run the command as a portable tool for both Unix and Windows environments in a Go developer environment:

```console
$ go run github.com/dolmen-go/env@latest -i a=1 b=2
a=1
b=2
```

## Why?

This pure Go implementation of `env` allows to use set environment variables in a portable way for [`//go:generate` commands](https://pkg.go.dev/cmd/go@latest#hdr-Generate_Go_files_by_processing_source):

```
//go:generate -command env go run github.com/dolmen-go/env@v0.1.0
//go:generate env GOTELEMETRY=off GOPROXY=off go run process.go
```

Or even:

```
//go:generate -command go-run go run github.com/dolmen-go/env@v0.1.0 GOTELEMETRY=off GOPROXY=off go run
//go:generate go-run process.go
``````

## License

``````
Copyright 2023 Olivier Mengué

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
``````