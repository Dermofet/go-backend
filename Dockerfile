FROM golang:1.20

WORKDIR /app/

COPY . /app

RUN go mod tidy

RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag
# RUN swag init -g ./cmd/go-backend/main.go --output ./docs

CMD swag init -g ./cmd/go-backend/main.go --output ./docs && go run cmd/go-backend/main.go
