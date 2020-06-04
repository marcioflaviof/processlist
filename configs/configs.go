package configs

//=========================
//	Server HTTP
//=========================
const (

	// defines ip and port address for server instance
	SERVER_ADDR = ":"
	//SERVER_ADDR = "localhost:8080"

	// host for mongo db
	MONGO_HOST = "mongodb+srv://admin:1010@cluster0-ga3ne.mongodb.net/Processes?retryWrites=true&w=majority"
	//MONGO_HOST = "mongodb://localhost:27017"
)

//========================
//  DB Collections
//========================
const (
	PROCESSES_COLLECTION = "processes"
)
