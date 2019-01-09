FROM adaickalavan/gocvkafka as build-stage

# Install Go application
WORKDIR /go/src/app/
COPY . .
RUN GOOS=linux go install -a -tags static .

FROM alpine

#Author label
LABEL Author Adaickalavan Meiyappan

# Install runtime tools
RUN apk add --update git openssl cyrus-sasl-dev

# Copy necessary files
COPY --from=build-stage /go/bin/app /go/bin/app
COPY --from=build-stage /go/pkg /go/pkg

ENTRYPOINT ["/go/bin/app"]