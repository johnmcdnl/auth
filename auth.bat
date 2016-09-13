cls
go get
godep save
godep update
go install
cls
docker build -t auth .
docker rm -f auth
docker run --name auth -p 8600:8600 auth
