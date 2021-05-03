PROJECT := huddlet
DBNAME := huddlet

# Development
.PHONY: develop
develop:
	make db.setup
	go run .

.PHONY: db.setup
db.setup:
	createdb -h localhost -U postgres $(DBNAME) 2> /dev/null || true
	./db/shmig -d $(DBNAME) up

