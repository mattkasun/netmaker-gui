FROM golang:latest as builder
ARG VERSION=dev
WORKDIR /build
COPY *.go go.* ./
RUN GOOS=linux go build -v -ldflags="-X 'main.version=${VERSON}'" .


FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /build/netmaker-gui .
ADD /images/* images/
ADD /html/* html/
CMD ["./netmaker-gui"]
