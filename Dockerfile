FROM golang:latest

# Install parallel and accept the citation notice (we aren't using this in a
# context where it make sense to cite GNU Parallel).

RUN ln -s /opt/bats/bin/bats /usr/sbin/bats
COPY ./bats-core-master /opt/bats/

ENTRYPOINT ["bash", "/usr/sbin/bats"]