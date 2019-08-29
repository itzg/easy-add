A utility for easily adding a file from a downloaded archive during Docker builds

## Example usage within `Dockerfile`

```
RUN easy-add --file restify --from https://github.com/itzg/restify/releases/download/1.2.0/restify_1.2.0_linux_amd64.tar.gz
```