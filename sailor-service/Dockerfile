FROM golang:1.24.11

COPY app /opt/af/


COPY cmd/server/config/static /opt/af/static

WORKDIR /opt/af/


RUN chmod 757 -R  /opt/af/  \
   && mkdir -p /opt/log/af  \
   && chmod 757 -R /opt/log/af

ENTRYPOINT ["/opt/af/app"]