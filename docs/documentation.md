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
to /user/signup with `username`, `password` and `address` in data:
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
```shell
curl -X POST http://localhost:8080/order \
    -H "Content-Type: application/json" \
    -H 'token: NyBldmlsIGV4ZXM=' \
    -d '{
        "coffees":[
            {"type": "Espresso", "sugar": 10},
            {"type": "Americano", "sugar": 8}
        ]
    }'
```
Response:

```json
  "order": {
    "id": "2",
    "coffees": [
      {
        "type": "Espresso",
        "sugar": 10
      },
      {
        "type": "Americano",
        "sugar": 8
      }
    ],
    "total": 4.5,
    "date": "2023-05-06T22:07:51.599183433Z"
  }
```

### Get one Order

For getting one, just send:
```shell
curl -X GET http://localhost:8080/order/1 \
    -H "Content-Type: application/json" \
    -H 'token: yBldmlsIGV4ZXM='
```
Response:
```json
  "order": {
    "id": "31",
    "coffees": [
      {
        "type": "Espresso",
        "sugar": 10
      },
      {
        "type": "Americano",
        "sugar": 8
      }
    ],
    "total": 4.5,
    "date": "2023-05-06T22:07:51Z"
  }
```

### Get all Orders

To get all orders:
```shell
curl -X GET http://localhost:8080/order/1 \
    -H "Content-Type: application/json" \
    -H 'token: yBldmlsIGV4ZXM='
```
Response:
```json

```

## Users

### Get your Account info

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

#### Get other Users info

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

### Update your Account

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

### Delete Account

To delete an account, you can use the DELETE method:
```shell
curl -X DELETE http://localhost:8080/user \
    -H 'Accept: application/json' \
    -H 'token:  NyBldmlsIGV4ZXM='
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