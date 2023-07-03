FROM   golang:1.20
WORKDIR /app
ENV TimeZone=Asia/Shanghai

RUN ln -snf /usr/share/zoneinfo/$TimeZone /etc/localtime && echo $TimeZone > /etc/timezone

COPY go.mod go.sum ./

COPY hello_server ./

RUN go env -w GOSUMDB=off
RUN go env -w GOPROXY=https://goproxy.cn,direct


RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /hello_server

CMD ["/hello_server"]