# How to use this template

Clone project
Create virtual environment and install pipenv

```
python3 -m venv venv
pip install pipenv
```

Install dependencies
```
pipenv install
```

Setup a database for project and update settings in config/settings.toml

Create .env file with follow config
```
FLASK_APP=app
FLASK_ENV=development
FLASK_DEBUG=1
ENV_FOR_DYNACONF=development

```
# Run REST app

```
flask run
```

## Run gRPC app
Install envoy proxy
```
https://www.envoyproxy.io/docs/envoy/latest/start/install
```

Run envoy proxy
```
envoy --config-path envoy.yaml
```
Run gRPC app
```
python grpc_server.py
```

### Run with docker-compose
Build image
```
docker build . -t grpc_app
```

```
docker-compose up -d 
```

# Run client
Build client protoc file

```
protoc -I src/grpc/protobufs user_tracking.proto  --js_out=import_style=commonjs:.  --grpc-web_out=import_style=commonjs,mode=grpcweb:.
```

Build webpack for web client
```
npx webpack client.js
```

Open file test.html with browser and test with file data.json


# Test performance with locust
Install locust
```
pipenv install locust
```

Run locust
```
locust -f locustfile.py
```
