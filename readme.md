# simple server makes things simple which are not simpler

[![Go](https://github.com/mugund10/simpleserver/actions/workflows/go.yml/badge.svg)](https://github.com/mugund10/simpleserver/actions/workflows/go.yml)

### To Run

```
git clone https://github.com/mugund10/simpleserver.git
cd simpleserver
docker build --platform linux/arm64 -t simpleserver .
docker run -d --name simpleserver   -p 443:443   simpleserver
```