FROM golang:1.18.6-alpine AS build

WORKDIR /outdoorsy

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go install ./...

# Final image
FROM alpine:3.16.2

# copy binary
COPY --from=build /go/bin/outdoorsy /usr/local/bin/outdoorsy

# copy input data
COPY --from=build /outdoorsy/data/* /data/

ENTRYPOINT [ "outdoorsy" ]