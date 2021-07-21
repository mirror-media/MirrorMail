FROM golang:1.16.3

WORKDIR /MirrorMail
RUN export GIN_MODE=release
ADD . /MirrorMail
RUN cd /MirrorMail && go build
VOLUME /var/log
EXPOSE 8080

CMD ["./MirrorMail"]