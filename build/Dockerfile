FROM golang:1.11.4 as builder

WORKDIR /go/src/github.com/cgetzen/ssh-operator/
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/ssh-operator cmd/manager/main.go

FROM alpine:3.8

ENV OPERATOR=/usr/local/bin/ssh-operator \
    USER_UID=1001 \
    USER_NAME=ssh-operator

# install operator binary
COPY --from=builder /go/bin/ssh-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENV WATCH_NAMESPACE=kubeapps
ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
