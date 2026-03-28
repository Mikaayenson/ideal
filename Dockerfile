# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS build
WORKDIR /src
RUN apk add --no-cache ca-certificates git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION=dev
ARG COMMIT=unknown
ARG BUILD_DATE=unknown
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X github.com/stryker/ideal/internal/version.Version=${VERSION} -X github.com/stryker/ideal/internal/version.Commit=${COMMIT} -X github.com/stryker/ideal/internal/version.BuildDate=${BUILD_DATE}" -o /out/ideal ./cmd/ideal

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /
COPY --from=build /out/ideal /ideal
USER nonroot:nonroot
ENTRYPOINT ["/ideal"]
