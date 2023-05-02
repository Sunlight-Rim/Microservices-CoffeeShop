# For vscode you can use https://marketplace.visualstudio.com/items?itemName=humao.rest-client

### CREATE new User account
curl -sX POST http://localhost:8080/user/signup \
    -H 'Accept: application/json' \
    -d '{
            "username": "test",
            "password": "test",
            "address":  "test"
    }'

### LOGIN into your account
curl -sX POST http://localhost:8080/user/login \
    -H 'Accept: application/json' \
    -d '{
          "username": "test",
          "password": "test"
    }'

# USERS

### GET information of your account
curl -sX GET http://localhost:8080/user \
    -H 'Accept: application/json' \
    -H 'token:  testToken'

### GET information of accounts with id 1 and 3
curl -sX GET http://localhost:8080/user \
    -H 'Accept: application/json' \
    -H 'token:  testToken' \
    -d '{
        "ids": [1, 3]
    }'

### UPDATE information of your account
curl -sX PATCH http://localhost:8080/user \
    -H 'Accept: application/json' \
    -H 'token:  testToken' \
    -d '{
        "user": {
            "username": "newName"
        }
    }'

### DELETE your account
curl -sX "DELETE" http://localhost:8080/user \
    -H 'Accept: application/json' \
    -H 'token:  testToken'

# ORDERS

### POST new Order
curl -sX POST http://localhost:8080/order \
    -H "Content-Type: application/json" \
    -H 'token: testToken' \
    -d '{
        "coffees":[
            {"type": "Espresso", "sugar": 10}
        ]
    }'

### GET all your Orders
curl -sX GET http://localhost:8080/order \
    -H "Content-Type: application/json" \
    -H 'token: testToken'

### GET one your Order
curl -sX GET http://localhost:8080/order/1 \
    -H "Content-Type: application/json" \
    -H 'token: testToken'

### PATCH updated Order
curl -X "PATCH" http://localhost:8080/order/1 \
    --include \
    --header "Content-Type: application/json" \
    --data '{"Type":"Mocha"}'

### DELETE one Order
curl -X "DELETE" http://localhost:8080/order/1