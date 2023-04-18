# SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
#
# SPDX-License-Identifier: MIT

FROM golang:1.20-alpine as builder
WORKDIR /cli

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build -o disclosure-cli

FROM gcr.io/distroless/static-debian11
WORKDIR /
COPY --from=builder /cli/disclosure-cli .
ENTRYPOINT ["/disclosure-cli"]