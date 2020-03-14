[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/itzg/easy-add)](https://github.com/itzg/easy-add/releases/latest)
![CircleCI](https://img.shields.io/circleci/build/github/itzg/easy-add)

A utility for easily adding a file from a downloaded archive during Docker builds

## Usage

Running the binary with `--help` can be used to obtain usage at any time.

```
  -file path
    	The path to executable to extract within archive
  -from URL
    	URL of a tar.gz archive to download. May contain Go template references to 'var' entries.
  -mkdirs
    	Attempt to create the directory path specified by to
  -to path
    	The path where executable will be placed (default "/usr/local/bin")
  -var name=value
    	Sets variables that can be referenced in 'from'. Format is name=value (default os=linux,arch=amd64)
  -version
    	Show version and exit
```

## Template variables in `from`

The `from` argument is process as a Go template with `var` as the context. For example, repetition in the URL can be simplified such as:

```
--var version=1.2.0 \
--from https://github.com/itzg/restify/releases/download/{{.version}}/restify_{{.version}}_{{.os}}_{{.arch}}.tar.gz
```

## Example usage within `Dockerfile`

```
FROM ubuntu

ARG EASY_ADD_VER=0.5.3
ADD https://github.com/itzg/easy-add/releases/download/${EASY_ADD_VER}/easy-add_${EASY_ADD_VER}_linux_amd64 /usr/bin/easy-add
RUN chmod +x /usr/bin/easy-add

RUN easy-add --var version=1.2.0 --var app=restify --file restify --from https://github.com/itzg/{{.app}}/releases/download/{{.version}}/{{.app}}_{{.version}}_linux_amd64.tar.gz
```

## `Dockerfile` usage with BuildKit and multi-arch builds

```
# hook into docker BuildKit --platform support
# see https://docs.docker.com/engine/reference/builder/#automatic-platform-args-in-the-global-scope
ARG TARGETOS=linux
ARG TARGETARCH=amd64
ARG TARGETVARIANT=""

ARG EASY_ADD_VER=0.5.3
ADD https://github.com/itzg/easy-add/releases/download/${EASY_ADD_VER}/easy-add_${TARGETOS}_${TARGETARCH}${TARGETVARIANT} /usr/bin/easy-add
RUN chmod +x /usr/bin/easy-add
```