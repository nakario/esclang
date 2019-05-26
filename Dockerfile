FROM golang:1.12-stretch

RUN go get github.com/nakario/esclang/cmd/esc

ENTRYPOINT ["esc"]
