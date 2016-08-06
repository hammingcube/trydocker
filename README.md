export DOCKER_API_VERSION=1.21
http://stackoverflow.com/questions/28468260/does-the-docker-remote-api-have-an-equivalent-for-docker-run-rm

```sh
g++ -std=c++11 -o binary.exe *.cpp && ./binary.exe < testcases/prob-1-input-0001.txt
```

Docker:

```sh
docker run --rm -v "$(pwd)":/app -w /app gcc bash -c 'g++ -std=c++11 -o binary.exe *.cpp && ./binary.exe < testcases/prob-1-input-0001.txt'
```