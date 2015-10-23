# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/deligo

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)

RUN go get github.com/Masterminds/glide
RUN go build github.com/Masterminds/glide
#turns out RUN cd src/deligo && [some command] --also works, but lets do it nicely
WORKDIR src/deligo
RUN GO15VENDOREXPERIMENT=1 glide install
RUN GO15VENDOREXPERIMENT=1 go build

# Run the outyet command by default when the container starts.
ENTRYPOINT ["./deligo"]
#ENTRYPOINT /bin/bash --use this to look around in the image; passing --entrypoint /bin/bash when docker run, also works

# Document that the service listens on port 8080.
EXPOSE 8080

