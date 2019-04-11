# Use the official go docker image built on debian.
FROM golang:latest

# Grab the source code and add it to the workspace.
ADD . /go/src/revelChat

# Install revel and the revel CLI.
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

# Use the revel CLI to start up our application.
ENTRYPOINT revel run revelChat dev 8080

# Open up the port where the app is running.
EXPOSE 8080
