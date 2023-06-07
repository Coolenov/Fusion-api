FROM golang:alpine


WORKDIR /app

COPY . .

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/Coolenov/Fusion-library
RUN go get -u gorm.io/driver/mysql

EXPOSE 10000

CMD ["go", "run", "main.go"]

