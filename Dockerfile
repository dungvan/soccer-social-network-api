#this is only for development
FROM golang:latest

ENV ENV_API local
ENV SSN_FRONTEND_HOST http://localhost:3000
ENV SSN_API_DIR /go/src/github.com/dungvan2512/soccer-social-network-api

# install dependency tool
RUN go get -u github.com/golang/dep/cmd/dep

# Fresh for rebuild on code change, no need for production
RUN go get -u github.com/pilu/fresh

# for development, pilu/fresh is used to automatically build the application everytime you save a Go or template file
CMD fresh

EXPOSE 8080
