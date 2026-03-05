# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/skillsd ./cmd/skillsd

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /out/skillsd /usr/local/bin/skillsd
COPY examples/skills ./examples/skills
EXPOSE 8080
ENV SKILLS_API_TOKEN=dev-token
ENTRYPOINT ["skillsd", "serve", "--skills-dir", "/app/examples/skills", "--addr", ":8080"]
