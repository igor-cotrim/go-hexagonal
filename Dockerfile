FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go install github.com/spf13/cobra-cli@latest && go install github.com/golang/mock/mockgen@v1.6.0
RUN apt-get update && apt-get install sqlite3 -y

CMD ["tail", "-f", "/dev/null"]
