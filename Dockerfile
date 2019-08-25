FROM golang:1.12-stretch as builder

RUN apt-get update && \
    useradd -c 'gitlab-janitor user' -MNr -s /usr/sbin/nologin gitlab-janitor

WORKDIR /go/src/github.com/tlmiller/gitlab-janitor
COPY . .

ENV GO111MODULE=on
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o gitlab-janitor

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/tlmiller/gitlab-janitor/gitlab-janitor /app/gitlab-janitor
USER gitlab-janitor
ENTRYPOINT ["/app/gitlab-janitor"]
CMD ["run"]
