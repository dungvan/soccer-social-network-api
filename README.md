# soccer-Social-Network
### setup:
- Install docker, docker-compose.
### run local development:
- step 1 config the enviroment of project in the Dockerfile with the following variable:
    - ENV_API (local, development, staging, production,... with each value, need to create a json config file in config folder for e.g ENV_API = local, so json config file is config/local.json)
    - SSN_FRONTEND_HOST is the host of client using this api server for cors
    - SSN_API_DIR the base folder of this project
- step 2 build docker
```bash
$ git clone https://github.com/dungvan/soccer-social-network-api
$ cd soccer-social-network-api
$ docker-compose build
$ docker-compose up
```
##### open ngrok for local development, export the minio server to internet
- install ngrok https://ngrok.com/
- export port 9000 to internet
```bash
$ ./ngrok http 9000
```
- waiting for ngrok export successfully, copy the `endpoint` which ngrok generate to `soccer-social-network-api/config/local.json`, replace to value of `endpoint` of `objectstorage`
