# Initial preparations
FROM golang:1.11-alpine
RUN apk update && apk add git make curl

# Let's add the code in
ADD . /go/src/gitlab.com/mvenezia/cma-aws

# Build package in container
WORKDIR /go/src/gitlab.com/mvenezia/cma-aws
RUN make -f build/Makefile container-build

# Now that we have build a static binary, we can stuff it in a bare-bones alpine image
# But we will need bash and openssh right now
# ca-certificates is needed for aws library
FROM alpine
RUN apk update && apk add bash openssh ca-certificates
COPY --from=0 /go/src/gitlab.com/mvenezia/cma-aws/cma-aws /cma-aws

ENTRYPOINT ["/cma-aws"]
