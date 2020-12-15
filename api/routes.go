package api

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/", a.index).Methods("GET")

	a.Router.HandleFunc("/volume", a.volumeGet).Methods("GET")
	a.Router.HandleFunc("/volume", a.volumePost).Methods("POST")
}
