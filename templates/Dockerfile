FROM golang:alpine as builder

WORKDIR /go/src/github.com/{{.GithubOwner}}/{{.RepoName}}

COPY . .

RUN CGO=0 go build -a -o {{.RepoName}}

FROM alpine:3.8 as runtime

WORKDIR /opt/app

COPY --from=builder /go/src/github.com/{{.GithubOwner}}/{{.RepoName}} .

CMD ["./{{.RepoName}}"]