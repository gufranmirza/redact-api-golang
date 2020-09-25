# redact-api-golang

A microserver exposes an API /redact/ which will be used to redact a JSON based on a JSON path specified.
- Only redaction of leaf nodes of a JSON are allowed.
- No limitation on the level of nesting that can be expressed using the JSON path configured.


# Running Locally

Clone repository 
``` 
git clone https://github.com/gufranmirza/redact-api-golang
```
go to `src`  folder
```
cd redact-api-golang/src
```
Run the service
```
make run
```
Building the service
```
make build
```
It will create executable binary into the `/bin` folder


# Testing
go to `src`  folder

Run tests
```
make test
```

# Generating Mocks 
go to `src`  folder

Generate mocks for packages using the command

```
make mock
```
# API Documentation
API documentation is maintained with swagger.[ Here is the link to swagger file](https://github.com/gufranmirza/redact-api-golang/blob/master/api/swagger.yaml)

# Docker Image
A docker files is present in the `src` folder which can be used to create the docker image of this server

### Create Docker Image
To create the docker image run the following command
```
docker build . -t redact-api
```

### Running Docker Image
To run the docker image you have created using the following command
```
docker run -p 8001:8001 redact-api
```
It will run the service on the port:8001

# Tests

#### 1
```
curl -X POST \
  http://localhost:8001/redact/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: aa646d8f-3301-fb24-c921-3f3ca75fd3b4' \
  -d '{
    "json_to_redact": {
        "a": {
            "b": {
                "c": "va12345l",
                "d": "val",
                "e": "val"
            },
            "l": [
                {"k": "a"},
                {"k": {
                	"p": "va12345l"	
                }},
                {"k": "c"}
            ]
        }
    },
    "redact_regexes": [
        {
            "path": "a.l[1].k.p",
            "regexes": [
                "[0-9]{5}"
            ]
        }
    ],
    "redact_completely": [
        "a.l"
    ]
}'

{
    "a": {
        "b": {
            "c": "va12345l",
            "d": "val",
            "e": "val"
        },
        "l": [
            {
                "k": "a"
            },
            {
                "k": {
                    "p": "va*****l"
                }
            },
            {
                "k": "c"
            }
        ]
    }
}
```

#### 2
```
curl -X POST \
  http://localhost:8001/redact/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: f881702a-42d4-264a-72b3-71bdfeea9362' \
  -d '{
    "json_to_redact": {
        "a": {
            "b": {
                "c": "va12345l",
                "d": "val",
                "email": {
                	"name": "Mike",
                	"address": "mike@gmail.com"
                }
            },
            "l": [
                {"k": "a"},
                {"k": {
                	"p": "va12345l"	
                }},
                {"k": "c"}
            ]
        }
    },
    "redact_regexes": [
        {
            "path": "a.b.email.address",
            "regexes": [
                "^[a-zA-Z0-9.!#$%&'\''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
            ]
        }
    ],
    "redact_completely": [
        "a.l[1].k.p"
    ]
}'

{
    "a": {
        "b": {
            "c": "va12345l",
            "d": "val",
            "email": {
                "address": "**************",
                "name": "Mike"
            }
        },
        "l": [
            {
                "k": "a"
            },
            {
                "k": {
                    "p": "********"
                }
            },
            {
                "k": "c"
            }
        ]
    }
}
```
