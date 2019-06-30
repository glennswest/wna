export GIT_COMMIT=$(git rev-parse --short HEAD)
echo $GIT_COMMIT > wna.version
go get ./...
go build -o wna/wna -ldflags "-X main.builddate=`date -u +.%Y%m%d.%H%M%S` -X main.gitversion=$GIT_COMMIT" ./wna 
GOOS=windows GOARCH=386 go build -o "wna.exe" -ldflags "-X main.builddate=`date -u +.%Y%m%d.%H%M%S` -X main.gitversion=$GIT_COMMIT"  ./wna


