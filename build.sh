
if [[ "mac" == $1 ]]; then
    go build notifier.go
else
    GOOS=linux GOARCH=amd64 go build notifier.go
fi