FROM golang:1.24-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION
ARG BUILDTIME
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'ctf01d/internal/handler.version=${VERSION}' -X 'ctf01d/internal/handler.buildTime=${BUILDTIME}'" -o server ./cmd/main.go

FROM alpine:latest AS prod
LABEL maintainer="Evgenii Sopov <mrseakg@gmail.com>"
LABEL repository="https://github.com/sea5kg/ctf01d-training-platform"
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/html ./html
COPY --from=builder /app/api/openapi.yaml api/openapi.yaml

EXPOSE 4102

CMD ["./server"]
