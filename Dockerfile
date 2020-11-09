FROM golang:1.15.0-alpine as builder

COPY . /go/src
WORKDIR /go/src/hello-go

RUN CGO_ENABLED=0 GOOS=linux go build

FROM gcr.io/cloud-builders/gcloud
RUN apt-get install ca-certificates

COPY --from=builder /go/src/hello-go .

ARG GCP_PROJECT=ca-mizushima-sandbox
RUN gcloud config set project $GCP_PROJECT

ENTRYPOINT ["./hello-go"]
