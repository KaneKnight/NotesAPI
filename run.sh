export MONGO_URL=mongodb://127.0.0.1:27017 # MongoDB server URL.
export MONGO_DATABASE=thirdfort # MongoDB Project db name
export MONGO_COLLECTION=users # Collection to register all users
export MONGO_NOTES_COLLECTION=notes # Collection to register the notes.
go run server.go handlers.go models.go helpers.go