# Coffee Shop App Docs

- [Endpoints](#endpoints)
- [Auth](#auth)
- [Users](#users)
- [Orders](#orders)
- [DB Schema](#db-schema)

## Endpoints

| **Function** | URL                     | Method | Request Body                                  | Request header | Response Body                | Description                                                       |
|--------------|-------------------------|--------|-----------------------------------------------|----------------|------------------------------|-------------------------------------------------------------------|
| **Signup**   | /auth/signup            | POST   | username<br> address<br> password             |                | `User`                       | Register user                                                     |
| **Login**    | /auth/login             | POST   | username<br> password                         |                | accessToken<br> refreshToken | Login user and get pair of tokens                                 |
| **Refresh**  | /auth/refresh           | POST   | refreshToken                                  |                | accessToken                  | Refresh access token                                              |
| **GetMe**    | /user                   | GET    |                                               | accessToken    | `User`                       | View data of your account                                         |
| **GetOther** | /user/\<nickname\>      | GET    |                                               | accessToken    | `User`                       | View data of user with id=1                                       |
| **Update**   | /user                   | PATCH  | username<br> address<br> password<br> regdate | accessToken    | `User`                       | Change some info of your account                                  |
| **Delete**   | /user                   | DELETE |                                               | accessToken    | `User`                       | Delete your account                                               |
| **Create**   | /order                  | POST   | coffee<br> topping<br> sugar                  | accessToken    | `Order`                      | Create new order with specified coffees                           |
| **GetOne**   | /order/\<id\>           | GET    |                                               | accessToken    | `Order`                      | View your certain order                                           |
| **GetSome**  | /order?shift=\<number\> | GET    |                                               | accessToken    | [5]`Order`                   | View some your orders. Returns 5 orders starting from the "shift" |
| **Cancel**   | /order/\<id\>           | PATCH  |                                               | accessToken    | `Order`                      | Cancel your order (if order status hasn't been "DELIVERED")       |
| **Delete**   | /order/\<id\>           | DELETE |                                               | accessToken    | `Order`                      | Delete your certain order                                         |

<table>
<tr><td>

| Mark    | Fileds                                                                       |
|---------|------------------------------------------------------------------------------|
| `User`  | id<br> username<br> address<br> regdate                                      |

</td><td>

| Mark    | Fileds                                                                       |
|---------|------------------------------------------------------------------------------|
| `Order` | id<br> userid<br> status<br> coffee<br> topping<br> sugar<br> total<br> date |

</td></tr> </table>

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

Now, you can use recieved accessToken in your orders and users requests. Remember that the lifetime of accessToken is short. Payload of accessToken contains only "id" and "exp".

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
        "coffee": "Espresso",
        "topping": "Banana",
        "sugar": 2
    }'
```
Response:

```json
  "order": {
    "id": 2,
    "userid": 1,
    "status": "PENDING",
    "coffee": "Espresso",
    "topping": "Banana",
    "sugar": 2,
    "total": 3.5,
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
    "id": 2,
    "userid": 1,
    "status": "PENDING",
    "coffee": "Espresso",
    "topping": "Banana",
    "sugar": 2,
    "total": 3.5,
    "date": "2023-05-06T22:07:51Z"
  }
```

### Get some Orders

Get some orders by specifying a shift (offset):
```shell
curl -X GET http://localhost:8080/order?shift=0 \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ"
```
Response:
```json
"orders": [
  {
    "id": 1,
    "userid": 1,
    "status": "DELIVERED",
    "coffee": "Americano",
    "topping": "",
    "sugar": 1,
    "total": 3.5,
    "date": "2023-05-06T15:43:17Z"
  },
  {
    "id": 2,
    "userid": 1,
    "status": "PENDING",
    "coffee": "Espresso",
    "topping": "Banana",
    "sugar": 2,
    "total": 3.5,
    "date": "2023-05-06T22:07:51Z"
  }
]
```

### Cancel Order

You can cancel order only if its current status is not DELIVERED:
```shell
curl -X PATCH http://localhost:8080/order/2 \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ"
```
Response:

```json
  "order": {
    "id": 2,
    "userid": 1,
    "status": "CANCELLED",
    "coffee": "Espresso",
    "topping": "Banana",
    "sugar": 2,
    "total": 3.5,
    "date": "2023-05-06T22:07:51Z"
  }
```

### Delete Order

To delete order:
```shell
curl -X POST http://localhost:8080/order/2 \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTE2MjM5MDIyfQ.s_miVYCHBD3ZEFbfwOsdtrMsVtU7JM-ByB_ecLM8LrQ"
```
Response:

```json
  "order": {
    "id": 2,
    "userid": 1,
    "status": "CANCELLED",
    "coffee": "Espresso",
    "topping": "Banana",
    "sugar": 2,
    "total": 3.5,
    "date": "2023-05-06T22:07:51Z"
  }
```

## Users

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

Database corresponds to 3NF.
![schema](https://i.imgur.com/sg8K5zB.png)