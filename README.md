# soccer-Social-Network
### setup:
- Install docker, docker-compose.
### run local development:
- build docker
```bash
$ git clone https://github.com/dungvan2512/soccer-social-network
$ cd soccer-social-network
$ docker-compose build
$ docker-compose up
```
##### open ngrok for local development, export the minio server to internet
- install ngrok https://ngrok.com/
- export port 9000 to internet
```bash
$ ./ngrok http 9000
```
- waiting for ngrok export successfully, copy the `endpoint` which ngrok generate to `soccer-social-network/config/local.json`, replace to value of `endpoint` of `objectstorage`
