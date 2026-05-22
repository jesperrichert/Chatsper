FROM alpine:3.23.0 AS base
WORKDIR /app


FROM oven/bun:1 AS buildfrontend
WORKDIR /Chatsper.Frontend

COPY /Chatsper.Frontend/ .

RUN bun install
RUN bun run build

FROM golang:tip-alpine3.23 AS buildbackend
WORKDIR /Chatsper.Backend

COPY /Chatsper.Backend/ .

RUN sudo apt install build-essential
ENV CGO_ENABLED=1
RUN go build ./cmd/main.go

FROM base AS final

ENV FRONTEND_BUILD=/app/frontend

COPY --from=buildbackend /Chatsper.Backend/main .
COPY --from=buildfrontend /Chatsper.Frontend/build/ ./frontend

CMD ["./main"]