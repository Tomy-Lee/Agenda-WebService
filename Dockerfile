FROM golang:1.8
COPY . "$GOPATH/src/github.com/Tomy-Lee/Agenda-WebService"
RUN cd "$GOPATH/src/github.com/Tomy-Lee/Agenda-WebService/cli" && go get -v && go install -v
RUN cd "$GOPATH/src/github.com/Tomy-Lee/Agenda-WebService/service" && go get -v && go install -v
WORKDIR /
EXPOSE 8080
VOLUME ["/data"]
