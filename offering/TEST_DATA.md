# Values
### IDs
All ID values are UUID in offering service.

### SetCourierLocation:
grpcui request:
```json
{
  "method": "SetCourierLocation",
  "data": {
    "@type": "type.googleapis.com/google.protobuf.Struct",
    "value": {
      "location": {"lat": 12,"lon": 50},"CourierId": "10000000-0000-0000-0000-000000000000","CourierType": 1
    }
  }
}
```

### GetNearbyCouriers
grpcui request:
```json
{
  "method": "GetNearbyCouriers",
  "data": {
    "@type": "type.googleapis.com/google.protobuf.Struct",
    "value": {
      "location": {"lat": 12,"lon": 50},"CourierType": 1,"RadiusMeter": 5000
    }
  }
}
```

### RejectOffer
grpcui request:
```json
{
  "method": "RejectOffer",
  "data": {
    "@type": "type.googleapis.com/google.protobuf.Struct",
    "value": {
      "CourierId": "10000000-0000-0000-0000-000000000000","OfferId": "10000000-0000-0000-0000-000000000001"
    }
  }
}
```

### CancelOffer
grpcui request:
```json
{
  "method": "CancelOffer",
  "data": {
    "@type": "type.googleapis.com/google.protobuf.Struct",
    "value": {
      "OfferId": "10000000-0000-0000-0000-000000000001"
    }
  }
}
```

# gRPC Stream API
Please use grpcui in order to test stream API.

#### SetCourierLiveLocation
grpcui request:
```json
[
    {
        "courierId": "10000000-0000-0000-0000-000000000000",
        "courierType": "TRUCK",
        "location": {
            "lat": 12,
            "lon": 13
        },
        "time": "1970-01-01T00:00:00Z"
    },
    {
        "courierId": "10000000-0000-0000-0000-000000000000",
        "courierType": "TRUCK",
        "location": {
            "lat": 12,
            "lon": 13.2
        },
        "time": "1970-01-01T00:00:00Z"
    }
]
```

#### GetCourierLiveLocation
grpcui request:
```json
{
  "intervalSeconds": 10,
  "courierId": "10000000-0000-0000-0000-000000000000"
}
```
