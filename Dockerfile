FROM golang:1.19-alpine AS build-env

WORKDIR /code

# Install minimum necessary dependencies
RUN apk add --no-cache ca-certificates build-base git

# Add source files
COPY . .

# install binary
RUN go mod download
RUN make build

############################################################################
FROM alpine:3.16

# Copy over binaries from the build-env
COPY --from=build-env /code/bin/oracle-consumerd /usr/bin/oracle-consumerd

EXPOSE 26656 26657 1317 9090
WORKDIR /opt

CMD ["oracle-consumerd", "version"]  
