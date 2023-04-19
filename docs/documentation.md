# Coffee Shop App Docs

## Register User

Firslty, you need to register new user via sending POST request to /users:
```
...
```
Response:
```
{
  "token": "T3VyIERlbW9jcmFjeSBoYXMgYmVlbiBoYWNrZWQ="
}
```

Now you can use recieved token in your orders requests.

## Create new Order
You can order one coffee:
```
curl http://localhost:8080/order -sX \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{
        "token":"T3VyIERlbW9jcmFjeSBoYXMgYmVlbiBoYWNrZWQ=", \
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