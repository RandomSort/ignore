go build
docker run -v $(pwd):/code bats/bats:latest /code/test
