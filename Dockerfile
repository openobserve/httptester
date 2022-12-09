# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM public.ecr.aws/docker/library/golang:1.19 as build
ARG TARGETOS TARGETARCH

WORKDIR /go/src/app
COPY . .
ENV CGO_ENABLED=0 

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /go/bin/httptester
# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/httptester /
CMD ["/httptester"]

