PROJECT := huddlet
DBNAME := huddlet

# Development
.PHONY: develop
develop:
	go run .

.PHONY: db.setup
db.setup:
	createdb -h localhost -U postgres $(DBNAME) 2> /dev/null || true
	./migrations/shmig -d $(DBNAME) up

