include ./.env
DBURL=postgres://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=disable
MIGRATIONPATH=db/migrations

# make database migrations NAME=create_user_table (use : make migrate-create NAME=table_name)
migrate-create:
	migrate create -ext sql -dir $(MIGRATIONPATH) -seq create_$(NAME)_table

# make database migrate-up
migrate-createUp:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) up
# make database migrate-down
migrate-createDown:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) down

migrate-status:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) version

migrate-force:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) force $(v)