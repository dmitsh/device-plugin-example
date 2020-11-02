FROM ubuntu:20.04

COPY dp /usr/local/bin/dp

ENTRYPOINT /usr/local/bin/dp

