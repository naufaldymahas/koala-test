## API Spec
### Response Format:
```json
{
  "status": "boolean",
  "message": "string",
  "data": "dynamic"
}
```
### Register
- Headers:
  - Content-Type: "application/json"
- Method: POST
- Endpoint: /auth/register
- Request:
```json
{
  "customer_name": "string, required",
  "email": "string, required",
  "phone_number": "string, required, min:9, max:13",
  "dob": "string, required, format_string:'YYYY-MM-DD'",
  "sex": "boolean, required",
  "password": "string, required"
}
```
*note, for sex: true = male, false = female
- Example Request
```json
{
  "customer_name": "Naufaldy Mahas",
  "email": "naufaldymahas@gmail.com",
  "phone_number": "081213200349",
  "dob": "1996-06-22",
  "sex": "true",
  "password": "password123"
}
```
- Example Response:
```json
{
  "status": true,
  "message": "Customer created successfully",
  "data": null
}
```

### Login
- Headers:
  - Content-Type: "application/json"
- Method: POST
- Endpoint: /auth/register
- Request:
```json
{
  "phone_number_or_email": "string, required",
  "password": "string, required"
}
```
- Example Request:
```json
{
  "phone_number_or_email": "naufaldymahas@gmail.com",
  "password": "password123"
}
```

- Example Response:
```json
{
  "status": true,
  "message": "",
  "data": {
    "token": "randomtoken",
    "refresh_token": "randomtoken"
  }
}
```

### Refresh Token
- Headers:
  - Authorization: "Bearer <token>"
- Method: PATCH
- Endpoint: /auth/refresh

- Example Response:
```json
{
  "status": true,
  "message": "",
  "data": {
    "token": "randomtoken",
    "refresh_token": "randomtoken"
  }
}
```

### Create Order
- Headers:
  - Content-Type: "application/json"
  - Authorization: "Bearer <token>"
- Method: POST
- Endpoint: /order
- Request:
```json
{
  "payment_method_id": "string, required",
  "order_details": [
    {
      "product_id": "string, required",
      "qty": "number, required"
    }
  ]
}
```
- Example Request:
```json
{
  "payment_method_id": "332d5fda-b787-4e78-886b-f576b8886828",
  "order_details": [
    {
      "product_id": "dca0368f-7ad2-4005-9e34-c9c5ec584b77",
      "qty": 3
    }
  ]
}
```
- Example Response:
```json
{
  "status": true,
  "message": "",
  "data": {
    "order_id": "9bdfd70f-1fad-4b98-b751-768d091bd98c",
    "customer_id": "b94b4ef1-99ca-44e8-baf5-08ad8b8ea0d3",
    "order_number": "PO-130/XII/2020",
    "order_date": "2020-12-31T00:46:12.373659+07:00",
    "payment_method_id": "332d5fda-b787-4e78-886b-f576b8886828",
    "order_details": [
      {
        "order_detail_id": "5d02f1c0-0d75-4171-8225-7d199ee93ca4",
        "order_id": "9bdfd70f-1fad-4b98-b751-768d091bd98c",
        "product_id": "dca0368f-7ad2-4005-9e34-c9c5ec584b77",
        "qty": 3,
        "created_date": "2020-12-31T00:46:12.375327+07:00",
      }
    ]
  }
}
```