FROM alpine:3.23.0 AS base
WORKDIR /app


FROM oven/bun:1 AS buildfrontend
WORKDIR /Chatsper.Frondend

COPY /Chatsper.Frondend/ .

RUN bun install
RUN bun run build

FROM golang:tip-alpine3.23 AS buildbackend
WORKDIR /koop-feedback.Backend

COPY /Chatsper.Backend/ .

RUN go build ./cmd/main.go

FROM base AS final

ENV FRONTEND_BUILD=/app/frontend

COPY --from=buildbackend /Chatsper.Backend/main .
COPY --from=buildfrontend /Chatsper.Frondend/build/ ./frontend

CMD ["./main"]