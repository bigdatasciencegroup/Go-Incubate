#Start from a Debian image with the latest version of GO 
#installed and a workspace (GOPATH) configure at /go.
FROM golang:latest

#Author label
LABEL Author Adaickalavan Meiyappan

# install dependencies
RUN go get github.com/Shopify/sarama

#Entrypoint
ENTRYPOINT ["/app/docker-entrypoint.sh"]