############################################
#          DO NOT TOUCH THIS FILE          #
############################################

FROM golang:latest

RUN mkdir -p "/lhgames"
WORKDIR "/lhgames"
COPY . .

RUN go get -d -v "./game/"

EXPOSE 3000

CMD ["/bin/sh", "-c", "go run ./game/*.go"]
