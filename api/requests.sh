# For vscode https://marketplace.visualstudio.com/items?itemName=humao.rest-client

### GET all Orders
curl http://localhost:8080/orders

### GET one Order
curl http://localhost:8080/orders/1

### POST new Order
curl http://localhost:8080/orders \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"Type":"Espresso", "Sugar":10}'

### PATCH updated Order
curl -X "PATCH" http://localhost:8080/orders/1 \
    --include \
    --header "Content-Type: application/json" \
    --data '{"Type":"Mocha"}'

### DELETE one Order
curl -X "DELETE" http://localhost:8080/orders/1