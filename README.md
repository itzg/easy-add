A utility for easily adding a file from a downloaded archive during Docker builds

## Usage

Running the binary with `--help` can be used to obtain usage at any time.

```
  -file path
    	The path to executable to extract within archive
  -from URL
    	URL of a tar.gz archive to download
  -mkdirs
    	Attempt to create the directory path specified by to
  -to path
    	The path where executable will be placed (default "/usr/local/bin")
  -version
    	Show version and exit
```

## Example usage within `Dockerfile`

```
FROM ubuntu

ARG EASY_ADD_VER=0.2.1
ADD https://github.com/itzg/easy-add/releases/download/${EASY_ADD_VER}/easy-add_${EASY_ADD_VER}_linux_amd64 /usr/bin/easy-add
RUN chmod +x /usr/bin/easy-add

RUN easy-add --file restify --from https://github.com/itzg/restify/releases/download/1.2.0/restify_1.2.0_linux_amd64.tar.gz
```