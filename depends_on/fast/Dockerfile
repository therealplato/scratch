FROM golang:1.10
ENV F=/go/src/github.com/therealplato/scratch/fast
RUN mkdir -p $F
COPY main.go $F/main.go
RUN go build -o /fast $F/main.go 
ENTRYPOINT /fast
