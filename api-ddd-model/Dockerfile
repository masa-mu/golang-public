FROM golang:1.10

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 \
        && chmod +x /usr/local/bin/dep

RUN go get github.com/oxequa/realize

WORKDIR /go/src/golang-public/api-ddd-model
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only
COPY . .

ENTRYPOINT [ "bash", "-c", "dep check || dep ensure -v -vendor-only && realize start --run --no-config" ]

EXPOSE 8000
