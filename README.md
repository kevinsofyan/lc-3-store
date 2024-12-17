
# Live Coding 3 - P2

Benedict Kevin Sofyan - store order lc 3

## Demo 

the app deployed to heroku and can be accessible with this url
```
https://orders-lc3-717e5d36a486.herokuapp.com/swagger/index.html
```

use this sample username for login
```
{
  "email": "alice.johnson@example.com",
  "password": "hashed_password1"
}
```

## Database

this app use supabase postgre as database
use this env to run in local

```
DB_HOST=aws-0-ap-southeast-1.pooler.supabase.com
DB_USER=postgres.eayclkabtrbpmtybtynf
DB_PASSWORD=asdQWE123!
DB_NAME=postgres
DB_SCHEMA=store                  
DB_PORT=5432    
```

## Running in local

to run this in local just use this command

```
go run main.go
```

the default port will be :8080