# docker 



## postgis
docker pull kartoza/postgis
docker run -t --name postgressql --restart always  -e POSTGRES_USER=postgres  -e POSTGRES_PASSWORD=xxx  -e ALLOW_IP_RANGE=0.0.0.0/0 -p 5432:5432 -d kartoza/postgis


## 关于latest
build和push的时候都不指定版本号，不带`:v{version}`后缀，默认推送到`latest`分支，pull的时候也就不需要指定版本，run的时候同理
