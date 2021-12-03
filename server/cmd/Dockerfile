ARG ENV_URI 
ARG ENV_PORT
FROM golang:1.16-alpine
ENV URI=$ENV_URI
ENV PORT=$ENV_PORT
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD . /app
RUN go build -o /go-app

#CMD echo "/go-app -uri ${uri} -port ${port}"
#CMD ["/go-app","-uri=" ]
RUN echo "$ENV_PORT $ENV_URI"
CMD "/go-app" "-uri" $ENV_URI "-port" $ENV_PORT
EXPOSE 3000