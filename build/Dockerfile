FROM golang:1.17.3 as build

WORKDIR /go/src/app
COPY . .

RUN go install -v ./...

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -o ${BIN_FILE} cmd/collector/*

FROM alpine:3.15

ENV BIN_FILE "/opt/static_collector/collector"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /go/src/app/config/config.yaml
COPY /configs/config.yaml ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}