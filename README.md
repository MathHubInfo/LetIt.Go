# letit.go

A tiny server to let environment variables go as JSON. 

## Usage

```
Usage of ./letit:
  -bind string
        address to bind to (default "0.0.0.0:3000")
  -path string
        server path to serve the resulting json on (default "/")
  -vars string
        comma-seperated list of variables and globs to use
```

## Examples

To expose only the home variable on port 80

```
./letit -bind 127.0.0.1:80 -vars HOME
```

To expose all variables starting with LC on port 8080

```
./letit -bind 127.0.0.1:80 -vars 'LC*'
```

## Tests

[![Build Status](https://travis-ci.org/MathHubInfo/LetIt.Go.svg?branch=master)](https://travis-ci.org/MathHubInfo/LetIt.Go)

Travis Tests check if the project builds. 

## Docker

A [from scratch](https://hub.docker.com/_/scratch/) Dockerfile exists. 
This is on DockerHub as an automated build [mathhub/letitgo](https://hub.docker.com/r/mathhub/letitgo/).

By default, it will expose 3000. 
The environment variables to be exposed can be specified as arguments. 
For example:

```
      docker run --rm -p 3000:3000 -e HELLO=world -e ANSWER=42 mathhub/letitgo 'HELLO,ANSWER'
```

## License

Public Domain / Unlicense