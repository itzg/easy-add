A utility for easily adding a file from a downloaded archive during Docker builds

## Usage

Run the binary with `--help` to view the usage documentation.

## Example usage within `Dockerfile`

```
FROM ubuntu

ADD https://github.com/itzg/easy-add/releases/download/0.1.1/easy-add_0.1.1_linux_amd64 /usr/bin/easy-add
RUN chmod +x /usr/bin/easy-add

RUN easy-add --file restify --from https://github.com/itzg/restify/releases/download/1.2.0/restify_1.2.0_linux_amd64.tar.gz
```