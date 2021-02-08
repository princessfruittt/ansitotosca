FROM golang:1.15-alpine AS build

## Optionally set HUGO_BUILD_TAGS to "extended" or "nodeploy" when building like so:
##   docker build --build-arg HUGO_BUILD_TAGS=extended .
#ARG HUGO_BUILD_TAGS

ARG CGO=1
ENV CGO_ENABLED=${CGO}
ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR /go/src/github.com/princessfruittt/AnsiToTosca

COPY . /go/src/github.com/princessfruittt/AnsiToTosca/

# gcc/g++ are required to build SASS libraries for extended version
RUN apk update && \
    apk add --no-cache gcc g++ musl-dev && \
    go get github.com/magefile/mage

RUN mage hugo && mage install

# ---

FROM alpine:3.12

COPY --from=build /go/bin/ansitotosca /usr/bin/ansitotosca

## libc6-compat & libstdc++ are required for extended SASS libraries
## ca-certificates are required to fetch outside resources (like Twitter oEmbeds)
#RUN apk update && \
#    apk add --no-cache ca-certificates libc6-compat libstdc++ git

VOLUME /site
WORKDIR /site

# Expose port for live server
EXPOSE 1313

ENTRYPOINT ["ansitotosca"]
CMD ["--help"]