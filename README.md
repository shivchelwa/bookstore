# bookstore

Setup Postgres DB:

  1. Install postgres database
  2. export PGPASSWORD=<postgres master password>
  3. psql -U postgres -h localhost -p 5432 -f dbscripts.sql

Build Go:
  
$GOOS=linux GARCH=amd64 go build bookstore
  
Build Docker:
  
  $docker build -t bookstore:1.0.0
  
Run Docker:
  
  $docker run -e "PGHOST=localhost" -e "PGPORT=5432" -e "PGDATABASE=postgres" -e "PGUSER=postgres" -e "PGPASSWORD=postgres" -p 3000:3000 bookstore:1.0.0
  
Test:
  
  $curl http://localhost:3000/books
