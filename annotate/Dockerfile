FROM golang:alpine AS build-stage
ADD . /work
WORKDIR /work
RUN go build -o go-annotate .

FROM alpine:latest
COPY --from=build-stage /work/go-annotate /usr/local/bin/go-annotate
ENTRYPOINT ["/usr/local/bin/go-annotate"]
