FROM golang:latest as builder
WORKDIR /build
COPY *.go go.* ./
RUN GOOS=linux go build -v .


FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /build/netmaker-gui .
ADD /images/* images/
ADD /html/* html/
CMD ["./netmaker-gui"]
