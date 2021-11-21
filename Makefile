include config/.env
export
migrateup: 
	migrate -source file\://go-migrate -database postgres\://$(user):$(password)@0.0.0.0\:$(port)/$(dbname)?sslmode=disable up 2
migratedown: 
	migrate -source file\://go-migrate -database postgres\://$(user):$(password)@0.0.0.0\:$(port)/$(dbname)?sslmode=disable down 2