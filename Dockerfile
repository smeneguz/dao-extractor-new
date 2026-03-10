# This Docker allows running the extractor binary.
#
# How to build the image:
# > docker build --secret id=github_token,env=GITHUB_TOKEN --tag dao-portal/extractor .
#
# How to run the image:
# > docker run dao-portal/extractor

FROM golang:1.24-alpine AS build-env

# Set up dependencies
ENV PACKAGES="curl make git libc-dev bash gcc linux-headers eudev-dev build-base"
RUN set -eux; apk add --no-cache $PACKAGES;

# Set working directory for the build
WORKDIR /code

# Setup git to make sure that we are using GOPRIVATE correctly
RUN --mount=type=secret,id=github_token \
    git config --global url."https://$(cat /run/secrets/github_token):x-oauth-basic@github.com/".insteadOf "https://github.com/"

# Add sources files
COPY . /code/

# Set the entrypoint, so that the user can set the config using the CMD
RUN BUILD_TAGS=muslc GOOS=linux GOARCH=amd64 make build


# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /home

# Install bash
RUN apk add --no-cache bash

# Copy over binaries from the build-env
COPY --from=build-env /code/build/extractor /usr/bin/extractor
ENTRYPOINT ["extractor"]

