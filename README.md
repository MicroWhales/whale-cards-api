docker pull golang

under root
docker build -t whale-card-api .

docker run -dp 5000:5000 whale-card-api