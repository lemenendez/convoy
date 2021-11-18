# Convoy

This is a database connection builder. The naming of the package itself is an ongoing project ;(

## Run docker stand alone

`docker build -t convoy .`

`docker run --rm  -it -v $PWD:/go/src convoy bash`

## Usage

### Get

Inline `go get github.com/lemenendez/convoy @v1.0.0`

```go
if db, err = convoy.NewDB(convoy.Options{
   Host: dbhost,
   User: dbusername,
   Pass: dbuserpass,
   Port: dbport,
   DB: dbname,
   }); err != nil {
     panic(err)
}
```
