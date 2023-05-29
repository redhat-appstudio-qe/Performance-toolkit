FROM registry.access.redhat.com/ubi9/go-toolset:1.19.9-3

USER root
WORKDIR /tests
COPY . /tests

RUN go get -d ./...

RUN go build -o /build 

CMD [ "/build" ]
