#Compile
docker build -t gohello:latest -f Dockerfile .

#Run
docker  run -d -p 8080:8081 --name gohello gohello:latest
# go-docker
