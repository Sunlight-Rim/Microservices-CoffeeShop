# Coffee Shop App Docs

- [Endpoints](#endpoints)
- [Auth](#auth)
- [Orders](#orders)
- [Users](#users)

## Endpoints
| URL          | Description                                   | Method   | Data                                 |
|--------------|-----------------------------------------------|----------|--------------------------------------|
| /user/signup | Register user                                 | `POST`   | username<br> password<br> address    |
| /user/login  | Login user                                    | `POST`   | username<br> password                |
| /user        | View info of your account                     | `GET`    |                                      |
| /user        | View info of users with ID specified in Data  | `GET`    | ids                                  |
| /user        | Change some info of your account              | `PATCH`  | username<br> password<br> address    |
| /user        | Delete your account                           | `DELETE` |                                      |
| /order       | Create new order                              | `POST`   | (token in header)<br> type<br> sugar |
| /order       | View all your orders                          | `GET`    |                                      |
| /order/1     | View your certain order                       | `GET`    |                                      |
| /order/1     | Update your order if it hasn't been delivered | `PATCH`  | (token in header)<br> type<br> sugar |
| /order/1     | Delete your certain order                     | `DELETE` |                                      |

## Auth

### Register

Firslty, you need to register a new user via sending POST request
to /user/signup with `username`, `address` and `password` in data:
```shell
curl -X POST http://localhost:8080/user/signup \
    -H 'Accept: application/json' \
    -d '{
        "username": "ScottPilgrim",
        "password": "ramonalove",
        "address":  "65 Alberta Ave, Regal Heights, Toronto"
    }'
```
Response:
```json
"user": {
  "id":        1,
  "username":  "ScottPilgrim",
  "address":   "65 Alberta Ave, Regal Heights, Toronto",
  "regdate":   "2023-04-30T18:47:35Z"
}
```

### Login

Then you need to login in just created account by sending POST
to /user/login with `username` and `password`:
```shell
curl -X POST http://localhost:8080/user/login \
    -H 'Accept: application/json' \
    -d '{
        "username": "ScottPilgrim",
        "password": "ramonalove"
    }'
```
Response:
```json
"id":    1,
"token": "NyBldmlsIGV4ZXM="
```

Now, you can use recieved token in your orders and users requests.
The token is regenerated after you change your password.

## Orders

### Create new Order

You can order one coffee:
```

```
Response:

```

```

## Get one Order

For getting one, just send
```

```
Response:
```

```

## Users

### GET

You can get information about your account:
```shell
curl -X GET http://localhost:8080/user \
    -H 'Accept: application/json' \
    -H 'token: NyBldmlsIGV4ZXM='
```
Response:
```json
"user": {
  "id":        1,
  "username":  "ScottPilgrim",
  "address":   "65 Alberta Ave, Regal Heights, Toronto",
  "regdate":   "2023-04-30T18:47:35Z"
}
```

Also you can get information about other accounts by specifying their IDs:
```shell
curl -X GET http://localhost:8080/user \
    -H 'Accept: application/json' \
    -H 'token:  NyBldmlsIGV4ZXM=' \
    -d '{
        "ids": [1, 2]
    }'
```
Response:
```json
"users": [
  {
    "id": "1",
    "username": "ScottPilgrim",
    "address": "65 Alberta Ave, Regal Heights, Toronto",
    "regdate": "2023-04-30T18:47:35Z"
  },
  {
    "id": "2",
    "username": "WallaceWells",
    "address": "66 Alberta Ave, Regal Heights, Toronto",
    "regdate": "2023-05-01T15:13:20Z"
  }
]
```

### UPDATE

You can update some information of your account:
```shell
curl -X PATCH http://localhost:8080/user \
    -H 'Accept: application/json' \
    -H 'token:  NyBldmlsIGV4ZXM=' \
    -d '{
        "user": {
            "username": "Scott"
        }
    }'
```
Response:
```json
"user": [
  {
    "id": "1",
    "username": "Scott",
    "address": "65 Alberta Ave, Regal Heights, Toronto",
    "regdate": "2023-04-30T18:47:35Z"
  }
]
```