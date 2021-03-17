go mod init example.com/postgre

docker run --rm --name mydb -e POSTGRES_PASSWORD=password -e POSTGRES_USER=admin -d -p 5432:5432 postgres
docker exec -it mydb psql -U admin
#https://www.datacamp.com/community/tutorials/10-command-line-utilities-postgresql

https://www.postgresqltutorial.com/

# list all db
\l
# select db postgres
\c postgres
# list all tables
\dt
# show table resources
\d resources
# show data
select * from resources;

# time string to time.Time
time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
