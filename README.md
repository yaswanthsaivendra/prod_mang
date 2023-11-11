# prod_mang

## How to run the project

1. clone the repo
2. make a .env.local using .env file
```
cp .env .env.local
```
3. create a postgres db using your local installation of postgres
4. Replace the db creds in .env.local file
5. Run the project
```
go run main.go
```