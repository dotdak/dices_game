# syntax=docker/dockerfile:1
FROM golang:1.17-alpine AS base-dev
RUN apk --no-cache add build-base git mercurial gcc bash

WORKDIR /app
COPY go.mod .
RUN go mod download

COPY . .
RUN go build .

FROM alpine:3.14.2

RUN apk --no-cache add ca-certificates tzdata htop tini bash curl
COPY --from=base-dev /app/dices_game /dices_game
COPY --from=base-dev /app/index.html /index.html
ARG PORT
ENV PORT=${PORT}
ENTRYPOINT ./dices_game

