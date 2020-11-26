FROM golang
ADD . /go/src/github.com/rays/go-pagerduty
WORKDIR /go/src/github.com/rays/go-pagerduty
RUN go get ./... && go test -v -race -cover ./...
