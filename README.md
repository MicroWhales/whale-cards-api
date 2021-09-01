docker pull golang

under root
docker build -t whale-card-api .

docker run -dp 5000:5000 whale-card-api

To install dependencies, use the go get command, which will also update the go. mod file automatically. Since the package is not currently used anywhere in the project, it's marked as indirect. This comment may also appear on an indirect dependency package; that is, a dependency of another dependency.