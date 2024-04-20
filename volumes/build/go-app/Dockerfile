# =============== build stage ===============
FROM golang:alpine AS build
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN go build -ldflags "-s -w" -o app -v ./cmd/bot/main.go

# =============== final stage ===============
FROM chromedp/headless-shell:93.0.4535.3 AS dev

WORKDIR /build
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/app /build/
COPY --from=build /build/settings/config.yaml /build/settings
RUN mkdir data 
RUN cd data \
    && mkkdir free_account \
    && mkkdir work_account
# COPY --from=build /build/data /build/data
ENTRYPOINT ["/build/app", "-other", "flags"]