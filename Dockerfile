FROM golang:1.22.4-alpine AS builder

WORKDIR /build
COPY . .
#To copy everything into some directory 
#COPY go.mod /build/
#COPT main.go /build/
#But Copy Everything COPY . .

RUN go mod download
RUN go build -o ./userapi
#copying into gcr.io/distroless/base-debian12 is a best practise for us
#because it is a lightweight image from golang

FROM gcr.io/distroless/base-debian12

WORKDIR ./app
COPY --from=builder /build/userapi ./userapi 

CMD ["/app/userapi"]