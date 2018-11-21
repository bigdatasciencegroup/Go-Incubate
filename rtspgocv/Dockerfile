FROM denismakogon/gocv-alpine:3.4.2-buildstage as build-stage

RUN go get -u -d gocv.io/x/gocv
RUN cd $GOPATH/src/gocv.io/x/gocv && go build -o $GOPATH/bin/gocv ./cmd/version/main.go

# #Install tools required to build project
# RUN apk add --no-cache git
# RUN go get github.com/golang/dep/cmd/dep

# #Install library dependencies
# COPY Gopkg.lock Gopkg.toml /go/src/app/
# WORKDIR /go/src/app/
# RUN dep ensure -vendor-only

#Copy the local package files (from development computer) to the container's workspace (docker image)
COPY . /go/src/app 
# BUild the app command inside the container
RUN go install app


FROM denismakogon/gocv-alpine:3.4.2-runtime

#Author label
LABEL Author Adaickalavan Meiyappan

# COPY --from=build-stage /go/bin/gocv /go/bin/gocv
COPY --from=build-stage /go/bin/app /go/bin/app
# CMD ["/go/bin/gocv"]

# ENV DISPLAY 10.0.75.1:0.0
# VOLUME /tmp/.X11-unix:/tmp/.X11-unix

ENTRYPOINT ["/go/bin/app"]
