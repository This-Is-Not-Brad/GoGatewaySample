$env:GOOS = "linux"
go build -o .\bin\sample .\src\main.go

serverless deploy