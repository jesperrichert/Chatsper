FROM debian:trixie-slim AS base
WORKDIR /app

FROM node:24-trixie-slim AS buildfrontend
WORKDIR /Chatsper.Frontend

COPY /Chatsper.Frontend/ .

RUN npm install
RUN npm run build

FROM golang:tip-trixie AS buildbackend
WORKDIR /Chatsper.Backend

COPY /Chatsper.Backend/ .

RUN apt update -y && apt install -y sudo && sudo apt install -y build-essential
ENV CGO_ENABLED=1
ENV IS_DOCKER_BUILD=YES

RUN go build ./cmd/main.go

FROM base AS final

ENV FRONTEND_BUILD=/app/frontend

COPY --from=buildbackend /Chatsper.Backend/main .
COPY --from=buildfrontend /Chatsper.Frontend/build/ ./frontend

CMD ["./main"]