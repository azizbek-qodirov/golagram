migrate_up:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/botapi_test -verbose up

migrate_down:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/botapi_test -verbose down

migrate_force:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/botapi_test -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table