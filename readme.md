# simple server makes things simple which are not simpler

### To Run
```
 docker build --platform linux/arm64 -t simpleserver .
 docker run -d --name simpleserver   -p 443:443   simpleserver
```