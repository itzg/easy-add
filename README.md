[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/itzg/easy-add)](https://github.com/itzg/easy-add/releases/latest)

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

ARG EASY_ADD_VER=0.2.1
ADD https://github.com/itzg/easy-add/releases/download/${EASY_ADD_VER}/easy-add_${EASY_ADD_VER}_linux_amd64 /usr/bin/easy-add
RUN chmod +x /usr/bin/easy-add

RUN easy-add --file restify --from https://github.com/itzg/restify/releases/download/1.2.0/restify_1.2.0_linux_amd64.tar.gz
```