ARG GO_VERSION=1.21
ARG OS_NAME=alpine
ARG OS_VERSION=3.17

ARG BUILD_IMAGE=golang:${GO_VERSION}-${OS_NAME}${OS_VERSION}
ARG RUN_IMAGE=${OS_NAME}:${OS_VERSION}

FROM ${BUILD_IMAGE} AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ARG LDFLAGS
COPY . .
RUN go build -ldflags="${LDFLAGS}" -mod readonly -o /app/bin/golang-api-template /app/cmd/api/main.go

FROM ${RUN_IMAGE}

COPY --from=build /app/bin /bin

EXPOSE 3000
CMD ["/bin/golang-api-template"]
