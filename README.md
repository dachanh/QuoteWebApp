<h1>Build:</h1>

MySQL

```
            docker run  --name quoteAppmysql -e --privileged=true -e MYSQL_ROOT_PASSWORD="ead8686ba57479778a76e" -e MYSQL_USER="dev1" -e MYSQL_PASSWORD="19e5a718a54a9fe0559dfbce6908" -e MYSQL_DATABASE="Quote_Database" -p 3306:3306 bitnami/mysql:5.7
  ``` 
<h1>Deploy:</h1>

```
 chmod +x build.sh && ./build.sh
```
