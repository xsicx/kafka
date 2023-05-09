ARG GOLANG_VERSION=1.20.4

### GOLANG BUILD
FROM golang:${GOLANG_VERSION}-alpine as build

RUN apk --no-cache add ca-certificates
RUN apk --no-cache add \
    git \
    ;

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod verify && go install -v ./... ;

### GOLANG DEV
FROM build as dev

CMD ["sleep", "86400"]

### GOLANG PROD
FROM alpine:latest AS prod

WORKDIR /go/src/app

RUN apk add --no-cache tzdata;
###Added user 1000uid/guid and use it###
ENV USER=app
ENV UID=1000
ENV GID=1000
###Added user sensei and use it###
RUN addgroup --gid "$GID" "$USER"
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)"\
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"
USER $USER

FROM prod AS cli
COPY --from=build --chown=$USER:$USER /go/bin/cli /go/bin/cli
CMD ["sleep", "86400"]