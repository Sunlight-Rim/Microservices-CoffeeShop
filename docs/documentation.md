# Coffee Shop App Docs

- [Endpoints](#endpoints)
- [Register and login](#register-and-login)
- [Orders](#orders)
- [Users](#users)

## Endpoints
| URL          | Description                                   | Method   | Data                              |
|--------------|-----------------------------------------------|----------|-----------------------------------|
| /user/signup | Register user                                 | `POST`   | username<br> password<br> address |
| /user/login  | Login user                                    | `POST`   | username<br> password             |
| /user        | View all users data                           | `GET`    |                                   |
| /user/1      | View data of user with id=1                   | `GET`    |                                   |
| /user        | Change some info of your account              | `PATCH`  | username<br> password<br> address |
| /user        | Delete your account                           | `DELETE` |                                   |
| /order       | Create new order                              | `POST`   | type<br> sugar(int)               |
| /order       | View all your orders                          | `GET`    |                                   |
| /order/1     | View your certain order                       | `GET`    |                                   |
| /order/1     | Update your order if it hasn't been delivered | `PATCH`  | type <br> sugar(int)              |
| /order/1     | Delete your certain order                     | `DELETE` |                                   |

## Register and login

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
{
  "user": {
    "id":        1,
    "username":  "ScottPilgrim",
    "address":   "65 Alberta Ave, Regal Heights, Toronto",
    "regdate":   "12.06.2023",
    "order_ids": []
  }
}
```

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
{
  "id":    1,
  "token": "NyBldmlsIGV4ZXM="
}
```

Now, you can use recieved token in your orders and users requests.
The token is regenerated after you change your password.

## Orders

### Create new Order

You can order one coffee:
```
curl http://localhost:8080/order -X \
    --include \
    --header 'token: NyBldmlsIGV4ZXM=' \
    --request "POST" \
    --data '{
        "coffees":[{"type":"Espresso", "sugar":10}] \
    }'
```
Response:

```
{
  "order": {
    ...
  }
}
```

## Get one Order
For getting one, just send
```
curl http://localhost:8080/order/25 -sX \
    --include \
    --header "Content-Type: application/json" \
    --request "GET" \           
    --data '{"token":"T3VyIERlbW9jcmFjeSBoYXMgYmVlbiBoYWNrZWQ="}'
```
Response:
```
{
  "order": {
    ...
  }
}
```

## Users