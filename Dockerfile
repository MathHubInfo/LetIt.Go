
FROM golang AS builder

# add our code to the build
WORKDIR /go/src/github.com/MathHubInfo/LetIt.Go
ADD letit.go .

# and build statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o letit letit.go

# copy the executable into a static thing
FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/MathHubInfo/LetIt.Go/letit letit

# And add it as the entry point
EXPOSE 3000
ENTRYPOINT ["./letit", "-vars"]