FROM golang:alpine as build
WORKDIR /usr/app
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/contactsapi/

FROM gcr.io/distroless/static as release
WORKDIR /usr/app
COPY --from=0 /usr/app/contactsapi .
USER nonroot
CMD ["./contactsapi"]

FROM release