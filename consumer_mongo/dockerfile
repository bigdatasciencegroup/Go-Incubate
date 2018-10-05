# # Start from a Debian image with the latest version of GO 
# # installed and a workspace (GOPATH) configure at /go.
# FROM golang:latest

# #Author label
# LABEL Author Adaickalavan Meiyappan

# #Copy the local package files (from development computer) to the container's workspace (docker image)
# COPY . /go/src/app 

# #Install dependencies
# RUN go get github.com/Shopify/sarama
# RUN go get github.com/joho/godotenv
# RUN go get github.com/gorilla/mux
# RUN go get github.com/wvanbergen/kafka/consumergroup
# RUN go get gopkg.in/mgo.v2
# RUN go get gopkg.in/mgo.v2/bson

# # BUild the app command inside the container
# RUN go install app

# #Run the app command by default when the container starts
# ENTRYPOINT /go/bin/app

FROM golang:onbuild
