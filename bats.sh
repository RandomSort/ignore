go build
docker run -v $(pwd):/code randomsort/gobats:latest /code/test
