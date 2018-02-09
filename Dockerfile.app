# Start from the base Go image
FROM golang

# Set /go/src/github.com/davidjulien/tesla as the CWD
WORKDIR /go/src/github.com/davidjulien/tesla

# Copy package source files to container
ADD . .

# Download and install dependency manager
RUN go get github.com/Masterminds/glide
RUN go install github.com/Masterminds/glide

# Install dependencies
RUN glide install

# Build Tesla
RUN go install github.com/davidjulien/tesla

# Start Tesla
ENTRYPOINT [ "tesla" ]