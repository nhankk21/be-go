FROM golang:1.16-alpine as build
WORKDIR /demo
COPY . .
RUN go mod vendor
RUN CGO_ENADBLED=0 go build --mod=vendor -o ./godk .

FROM alpine
COPY --from=build ./demo/godk ./bin/godk
EXPOSE 3000
CMD [ "./bin/godk" ]
