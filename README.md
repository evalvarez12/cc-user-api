
## Setup
Run in this order:
```
make images

make database

make api
```

*Address is 127.0.0.1:8082*

## End-points
### success:
```
{
  status: “successs”
  data: ...
}
```

### Error:
```
{
  status: “error”,
  message: “Error message”
}
```
## User autentification
To add a user use:
```
POST     /user                     // Adds a new User
```
Where the body is expected to have something like:
```
{
  "name": "JhonyBoy66",
  "complete_name": "Juan Perez Perez"
  "created_at": "2016-03-23T13:14:29.141081-06:00",
  "updated_at": "2016-03-23T13:14:29.141081-06:00",
  "password": "juanito"
}
```

Login uses takes only name and password in the body
```
POST     /user/login               // User login

{
    "name" : "JhonyBoy66",
    "password": "juanito"
}
```
If successful a `token` is generated and returned on a `success` type message.
After login the `token` must be attached as a `X-Auth-Token` header for each request.

## Charges
### Routes:
```
GET     /charges               // Get a list of all the charges
GET     /charges/day           // Get a list of all charges ordered by day
POST    /charges               // Add a charge - returns the id
GET     /charges/id?id=#       // Get charge with id
PUT     /charges?id=#          // Modify a charge
DELETE  /charges?id=#          // Delete a charge

```



`/charges` can take 1 or 2 parameters `since` and `to` to specify the time intervals of created charges to list.
Examples:
```
0.0.0.0:8082/charges?since=03/21/2012
0.0.0.0:8082/charges?to=01/21/2016
0.0.0.0:8082/charges?since=01/21/2014&to=01/27/2016
```

Example of listed charges:
```
{
  "success": true,
  "data": [
    {
      "id": 4,
      "user_id": 123,
      "category_id": 3,
      "account_id": 4,
      "name": "proident, commodo nostrud",
      "description": "esse in culpa",
      "created_at": "2015-04-23T19:00:40-05:00",
      "updated_at": "2015-04-29T04:05:40-05:00",
      "expected_date": "2015-09-21T01:42:04-05:00",
      "real_amount": {
        "value": 5439,
        "scale": 0
      },
      "amount": 5439,
      "currency_code": "USD",
      "kind": "income",
      "source": "Source B",
      "destination": "Destination E"
    },
    {
      "id": 5,
      "user_id": 123,
      "category_id": 4,
      "account_id": 3,
      "name": "amet, aute excepteur",
      "description": "fugiat laboris dolore",
      "created_at": "2015-09-08T01:06:29-05:00",
      "updated_at": "2015-11-26T19:05:30-06:00",
      "expected_date": "2016-01-19T14:39:15-06:00",
      "real_amount": {
        "value": 141,
        "scale": 0
      },
      "amount": 141,
      "currency_code": "MXN",
      "kind": "expense",
      "source": "Source C",
      "destination": "Destination D"
    },
    {
      "id": 6,
      "user_id": 123,
      "category_id": 5,
      "account_id": 4,
      "name": "voluptate in cillum",
      "description": "qui ad sed",
      "created_at": "2015-06-10T13:25:44-05:00",
      "updated_at": "2015-08-06T12:10:05-05:00",
      "expected_date": "2016-03-12T21:46:31-06:00",
      "real_amount": {
        "value": 6164,
        "scale": 4
      },
      "amount": 0.62,
      "currency_code": "EUR",
      "kind": "expense",
      "source": "Source C",
      "destination": "Destination E"
    }
  ]
}
```

To `add`/`update` the new `charge` is received in the body as:
```
{
   "name": "nueva carga",
   "description": "taquitos y tequila",
   "created_at": "2015-02-17T14:03:16-06:00",
   "updated_at": "2015-12-16T10:06:16-06:00",
   "expected_date": "2015-01-05T17:14:12-06:00",
   "real_amount": {
     "value": 123,
     "scale": 1
   },
   "amount": 12.3,
   "coin": "MXN",
   "exchange_rate": {
     "value": 1,
     "scale": 0
   },
   "kind": "expense",
   "source": "Source B",
   "destination": "Destination D"
}
```

##PlannedCharges
### Routes:
```
GET     /charges/planned          // Get a list of all planned charges
GET     /charges/planned/next     // Get the amount of next income and expense
POST    /charges/planned          // Add a planned charge
DELETE  /charges/planned          // Delete a planned charge
```

Example of listed planned charges:
```
{
  "success": true,
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "category_id": 4,
      "account_id": 1,
      "name": "laborum deserunt veniam,",
      "description": "quis sed sunt",
      "amount": {
        "value": 6577,
        "scale": 3
      },
      "currency_code": "USD",
      "exchange_rate": 0,
      "kind": "expense",
      "source": "Source C",
      "destination": "Destination F",
      "periodicity": 1,
      "created_at": "2016-04-07 22:18:33",
      "updated_at": "2016-04-20 10:35:32",
      "since": "2016-05-09 18:03:55",
      "until": "2016-10-09 18:03:55"
    }
  ]
}
```

Example of listed next income/expense:
```
{
  "success": true,
  "data": {
    "expense": {
      "value": 6577,
      "scale": 3
    },
    "income": {
      "value": 0,
      "scale": 0
    }
  }
}
```


## Categories
### Routes:
```
GET     /categories               // Get a list of all the categories
GET     /categories/pie           // Get a list of all non zero categories
POST    /categories               // Add a category - returns the id
PUT     /categories?id=#          // Modify a category
DELETE  /categories?id=#          // Delete a category

```

Example of listed categories:
```
{
  "success": true,
  "data": [
    {
      "category_id": 1,
      "name": "Drinks",
      "created_at": "2016-03-23T13:14:29.141081-06:00",
      "updated_at": "2016-03-23T13:14:29.141081-06:00",
      "total": {
        "value": 0,
        "scale": 0
      }
    },
    {
      "category_id": 2,
      "name": "Food",
      "created_at": "2016-03-23T13:14:29.141081-06:00",
      "updated_at": "2016-03-23T13:14:29.141081-06:00",
      "total": {
        "value": 0,
        "scale": 0
      }
    }
  ]
}
```


To `add`/`update` the new `category` is received in the body as:
```
{
      "name": "New Category",
      "created_at": "2016-03-23T13:14:29.141081-06:00",
      "updated_at": "2016-03-23T13:14:29.141081-06:00",
      "total": {
        "value": 0,
        "scale": 0
      }
}
```

## Accounts
### Routes:
```
GET     /accounts                 // Get a list of all the accounts
POST    /accounts                 // Add a account - returns the id
PUT     /accounts?id=#            // Modify a account
DELETE  /accounts?id=#            // Delete a account

```

Example of listed accounts:
```
{
  "success": true,
  "data": [
    {
      "account_id": 2,
      "user_id" : 123,
      "name": "Retirement",
      "bank": "Banamex",
      "description": "For the future",
      "created_at": "2016-03-23T13:14:26.643708-06:00",
      "updated_at": "2016-03-23T13:14:26.643708-06:00",
      "total": {
        "value": 0,
        "scale": 0
      }
    },
    {
      "account_id": 3,
      "user_id" : 123,
      "name": "Paycheck",
      "bank": "IXE",
      "description": "Regular expenses",
      "created_at": "2016-03-23T13:14:26.643708-06:00",
      "updated_at": "2016-03-23T13:14:26.643709-06:00",
      "total": {
        "value": 0,
        "scale": 0
      }
    }
  ]
}
```


To `add`/`update` the new `account` is received in the body as:
```
{
  "name": "New Account",
  "bank": "Cartera",
  "description": "Gastar",
  "created_at": "2016-03-23T13:14:26.643708-06:00",
  "updated_at": "2016-03-23T13:14:26.643708-06:00",
  "total": {
    "value": 0,
    "scale": 0
  }
}
```
