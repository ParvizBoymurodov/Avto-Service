package app

func (receiver *server) InitRoutes() {
	//routes for auth, login and other...
	receiver.router.HandleFunc("/api/registered", receiver.Registered)
	receiver.router.HandleFunc("/api/login", receiver.Login)
	receiver.router.HandleFunc("/api/password/edit/{id}", receiver.ChangePass)
	receiver.router.HandleFunc("/api/personal/list", receiver.PersonalList)
	receiver.router.HandleFunc("/api/delete/personal/{id}", receiver.DeletePersonal)

	//routes for service setting
	receiver.router.HandleFunc("/api/add/service", receiver.AddService)
	receiver.router.HandleFunc("/api/service/list", receiver.ServiceList)
	receiver.router.HandleFunc("/api/edit/service/{id}", receiver.EditService)
	receiver.router.HandleFunc("/api/remove/service/{id}", receiver.RemoveService)

	//routes for service market
	receiver.router.HandleFunc("/api/add/market", receiver.AddMarket)
	receiver.router.HandleFunc("/api/market/list", receiver.MarketList)
	receiver.router.HandleFunc("/api/delete/market/{id}", receiver.RemoveMarket)
}
