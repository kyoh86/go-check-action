FROM golang:alpine AS build-stage
ADD . /work
WORKDIR /work
RUN go build -o go-check .

FROM alpine:latest
COPY --from=build-stage /work/go-check /usr/local/bin/go-check
ENTRYPOINT ["/usr/local/bin/go-check"]
