# README
## How to Run
```go run main.go```


## Endpoint 
### Register
```[POST] http://{your_local_host}:8080/auth/register```
```
body:
{
    email: "test@gmail.com",
    password: 1234567
    passwordConfirmation: 1234567
}
```

### Login 
```[POST]http://{your_local_host}:8080/auth/login```
```
body:
{
    email: "test@gmail.com",
    password: 1234567
}
```

### Logout
```[GET]http://{your_local_host}:8080/auth/logout```
```
body:
{
    email: "test@gmail.com",
    password: 1234567
}
```


### GET ALL COIN
```[GET]http://{your_local_host}:8080/coin```
```
header:
{
    token: <token>
}
```

### Track Coin
```[POST]http://{your_local_host}:8080/coin/track-coin```
```
header:
{
    token: <token>
}

body :
{
    coinId: <coinId>
}
```


### Untrack Coin
```[DELETE]http://{your_local_host}:8080/coin/untrack-coin```
```
header:
{
    token: <token>
}

body :
{
    coinId: <coinId>
}
```

# Host
``` https://cointracker-sangga.fly.dev ```
``` Endpoint same with above, just change the base https://cointracker-sangga.fly.dev/auth/login etc```


