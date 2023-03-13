# Coffee Shop App Docs

```
message Coffee {
    string type = 1;
    float price = 2;
}

message Order {
    enum Status {
        PENDING = 0;
        DELIVERED = 1;
    }
    int64 id = 1;
    repeated Coffee coffees = 2;
    float total = 3;
    google.protobuf.Timestamp date = 4;
    Status status = 5;
}
```

```
message User {
    int64 id = 1;
    string name = 2;
    string address = 3;
    google.protobuf.Timestamp date = 4;
    repeated int64 order_ids = 5;
    // also have token in database
}
```