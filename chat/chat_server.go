package chat


type ChatServer struct {

	Rooms map[string]*Room
	Bind_to string
}
