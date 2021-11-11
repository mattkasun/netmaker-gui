FROM golang:latest as builder
ARG VERSION=dev
WORKDIR /build
COPY *.go go.* ./
RUN GOOS=linux go build -ldflags="-X 'main.version=${VERSION}'" .


FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /build/netmaker-gui .
ADD /images/* images/
ADD /html/* html/
CMD ["./netmaker-gui"]
