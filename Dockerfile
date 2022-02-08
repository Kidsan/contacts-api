FROM golang:alpine as build
WORKDIR /usr/app
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/contacts-api/

FROM gcr.io/distroless/static as release
WORKDIR /usr/app
COPY --from=0 /usr/app/contacts-api .
USER nonroot
CMD ["./contacts-api"]

FROM release