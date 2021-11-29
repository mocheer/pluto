# docker 


docker pull kartoza/postgis
docker run -t --name postgressql --restart always  -e POSTGRES_USER=postgres   -e POSTGRES_PASSWORD=istrong9889  -e ALLOW_IP_RANGE=0.0.0.0/0 -p 5432:5432 -d kartoza/postgis
