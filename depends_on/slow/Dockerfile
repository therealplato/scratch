FROM golang:1.10
ENV S=/go/src/github.com/therealplato/scratch/slow
RUN mkdir -p $S
COPY main.go $S/main.go
RUN go build -o /slow $S/main.go 
ENTRYPOINT /slow
