# Coffee Shop App Docs

- [Endpoints](#endpoints)
- [Auth](#auth)
- [Users](#users)
- [Orders](#orders)
- [DBSchema](#db-schema)

## Endpoints
| URL            | Method   | Request Body                      | Request header | Response Body                                          | Description                                                           |
|----------------|----------|-----------------------------------|----------------|--------------------------------------------------------|-----------------------------------------------------------------------|
| /auth/signup   | `POST`   | username<br> address<br> password |                | id<br> username<br> address<br> regdate                | Register user                                                         |
| /auth/login    | `POST`   | username<br> password             |                | accessToken<br> refreshToken                           | Login user and get pair of tokens                                     |
| /auth/refresh  | `POST`   | refreshToken                      |                | accessToken                                            | Refresh access token                                                  |
| /user          | `GET`    |                                   | accessToken    | id<br> username<br> address<br> regdate                | View data of your account                                             |
| /user/1        | `GET`    |                                   | accessToken    | id<br> username<br> address<br> regdate                | View data of user with id=1                                           |
| /user          | `PATCH`  | username<br> address<br> password | accessToken    | id<br> username<br> address<br> regdate                | Change some info of your account                                      |
| /user          | `DELETE` |                                   | accessToken    | id<br> username<br> address<br> regdate                | Delete your account                                                   |
| /order         | `POSR`   | [type<br> sugar]<br> ...          | accessToken    | id<br> status<br> date<br> coffees<br> total           | Create new order with specified coffees                               |
| /order/1       | `GET`    |                                   | accessToken    | id<br> status<br> date<br> coffees<br> total           | View your certain order                                               |
| /order?shift=0 | `GET`    |                                   | accessToken    | [id<br> status<br> date<br> coffees<br> total]<br> ... | View some your orders. Returns 6 orders<br> starting from the "shift" |
| /order/1       | `PATCH`  | [type<br> sugar]<br> ...          | accessToken    | id<br> status<br> date<br> coffees<br> total           | Update your order (if order status hasn't been "DELIVERED")           |
| /order/1       | `DELETE` |                                   | accessToken    | id<br> status<br> date<br> coffees<br> total           | Delete your certain order                                             |

## Auth

### Register

Firslty, you need to register a new user via sending POST request
to /user/signup with `username`, `password` and `address` in data:
```shell
curl -X POST http://localhost:8080/auth/signup \
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
curl -X POST http://localhost:8080/auth/login \
    -H 'Accept: application/json' \
    -d '{
        "username": "ScottPilgrim",
        "password": "ramonalove"
    }'
```
Response:
```json
"accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ",
"refreshToken": "V2UgYXJlIFNleCBCb2ItT21i"
```

Now, you can use recieved accessToken in your orders and users requests. Remember that the lifetime of accessToken is short.

P.S.: In accessToken payload only contains "id" and "exp".

### Refresh

#TODO

If accessToken is expired you need to refresh it by using refreshToken:
```shell
curl -X POST http://localhost:8080/auth/refresh \
    -H 'Accept: application/json' \
    -d '{
        "refreshToken": "V2UgYXJlIFNleCBCb2ItT21i"
    }'
```
Response:
```json
"accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ",
"refreshToken": "SSBoYXZlIHRvIGdvIHBlZS4="
```

## Orders

### Create new Order

You can order one coffee:
```shell
curl -X POST http://localhost:8080/order \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ" \
    -d '{
        "coffees": [
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
      {"type": "Espresso", "sugar": 10},
      {"type": "Americano", "sugar": 8}
    ],
    "total": 4.5,
    "date": "2023-05-06T22:07:51Z"
  }
```

### Get one Order

For getting one, just send:
```shell
curl -X GET http://localhost:8080/order/2 \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ"
```
Response:
```json
  "order": {
    "id": "2",
    "coffees": [
      {"type": "Espresso", "sugar": 10},
      {"type": "Americano", "sugar": 8}
    ],
    "total": 4.5,
    "date": "2023-05-06T22:07:51Z"
  }
```

### Get some Orders

Get some orders:
```shell
curl -X GET http://localhost:8080/order?shift=0 \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ" \
```
Response:
```json
"orders": [
  {
    "id": "1",
    "coffees": [
      {"type": "Americano"}
    ],
    "total": 2.5,
    "date": "2023-05-05T23:25:13Z"
  },
  {
    "id": "2",
    "coffees": [
      {"type": "Espresso", "sugar": 10},
      {"type": "Americano", "sugar": 8}
    ],
    "total": 4.5,
    "date": "2023-05-06T22:07:51Z"
  }
]
```

### Update Order

You can update order only if its status is not DELIVERED:
```shell
curl -X POST http://localhost:8080/order \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ" \
    -d '{
        "coffees": [
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
      {"type": "Espresso", "sugar": 10},
      {"type": "Americano", "sugar": 8}
    ],
    "total": 4.5,
    "date": "2023-05-06T22:07:51Z"
  }
```

### Delete Order

To delete order:
```shell
curl -X POST http://localhost:8080/order \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ" \
    -d '{
        "id": 2
    }'
```
Response:

```json
  "order": {
    "id": "2",
    "coffees": [
      {"type": "Espresso", "sugar": 10},
      {"type": "Americano", "sugar": 8}
    ],
    "total": 4.5,
    "date": "2023-05-06T22:07:51Z"
  }
```

## Users

#TODO

### Get your Account info

You can get information about your account:
```shell
curl -X GET http://localhost:8080/user \
    -H 'Accept: application/json' \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ" \
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

### Get other Users info

Also you can get information about other account by specifying their IDs:
```shell
curl -X GET http://localhost:8080/user/2 \
    -H 'Accept: application/json' \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ" \
```
Response:
```json
"user": [
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
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ" \
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
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ"
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

## DB Schema

![schema](https://i.imgur.com/u3WRbz9.png)