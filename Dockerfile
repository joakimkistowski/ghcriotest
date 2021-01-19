FROM golang:alpine
#build container
RUN mkdir -p /go/src/ghcriotest/
WORKDIR /go/src/ghcriotest/
COPY ./*.go .
RUN GOOS=linux go build -v ghcriotest.go

FROM alpine
# Execution Container
COPY --from=0 /go/src/ghcriotest/ghcriotest /bin/ghcriotest

EXPOSE 8080

CMD ["ghcriotest"]
