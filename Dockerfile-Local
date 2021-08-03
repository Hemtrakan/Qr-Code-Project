FROM golang:1.15-alpine

RUN apk update && apk upgrade && apk add --no-cache build-base bash git openssh
LABEL maintainer="Chirapon Hemtrakan <chirapon.tdev@gmail.com>"

WORKDIR /app
COPY . .

RUN go mod download
RUN go get github.com/pilu/fresh

EXPOSE 8000
CMD ["sh", "-c", "fresh"]