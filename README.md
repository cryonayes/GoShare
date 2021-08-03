# GoShare
### GoShare is a new way to share files securely.

### How to use

Create your own database with name go_share or follow these steps if you don't have mysql installed.
```bash
docker pull mysql:latest
docker run -it --name GoShareDB --network host -e MYSQL_ROOT_PASSWORD=[ROOT_PASS] -d mysql:latest
docker exec -it GoShareDB /bin/bash
mysql -p # Enter ROOT_PASS
CREATE USER '[username]'@'%' IDENTIFIED BY '[passwd]';
GRANT ALL PRIVILEGES ON *.* TO 'cryonayes'@'%';
FLUSH PRIVILEGES;
```

Compile & run GoShare services
```bash
go build && ./GoShare
```
Run next.js server
```bash
cd nextjs
yarn build && yarn start
```

**Note:** If you having **"Error while saving data!"** error, you probably forgot to create your uploads folder. 