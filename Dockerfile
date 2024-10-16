# SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
#
# SPDX-License-Identifier: MIT

FROM golang:1.22-alpine as builder
WORKDIR /cli

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build -o disclosure-cli

FROM scratch
WORKDIR /
COPY --from=builder /cli/disclosure-cli .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/disclosure-cli"]
