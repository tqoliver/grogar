FROM golang:1.10.3-alpine3.7 AS build

# Install required tools for the project (both to build and to run). This is the requirements section
# Run `trash --file vendor.yaml` to get application dependencies
RUN apk add --no-cache git

# Copy the entire project to the image directory
COPY . /go/src/github.com/tqoliver/grogar/
WORKDIR /go/src/github.com/tqoliver/grogar/

RUN CGO_ENABLED=0 GOOS=linux go build -o /grogar

FROM alpine
COPY --from=build /grogar /
EXPOSE 8000
ENTRYPOINT [ "/grogar" ]