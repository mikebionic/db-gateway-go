GOOS=windows GOARCH=amd64 go build -o bin/sqlapi_console-amd64.exe .
GOOS=windows GOARCH=386 go build -o bin/sqlapi_console-386.exe .

GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o bin/sqlapi-amd64.exe .
GOOS=windows GOARCH=386 go build -ldflags -H=windowsgui -o bin/sqlapi-386.exe .