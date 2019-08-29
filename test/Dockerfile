FROM ubuntu

ARG EASY_ADD_VERSION=0.1.1

ADD https://github.com/itzg/easy-add/releases/download/${EASY_ADD_VERSION}/easy-add_${EASY_ADD_VERSION}_linux_amd64 /usr/bin/easy-add
RUN chmod +x /usr/bin/easy-add

RUN easy-add --file restify --from https://github.com/itzg/restify/releases/download/1.2.0/restify_1.2.0_linux_amd64.tar.gz

ENTRYPOINT ["/usr/bin/restify"]