package webserver

// create http router
func (ws *WebServer) router() {

	ws.server.Post("login", ws.login)

	ws.server.Post("retailer", ws.create)
	ws.server.Put("retailer", ws.update)
	ws.server.Get("retailer/:id", ws.getRetailerByID)
	ws.server.Get("retailer", ws.getRetailerList)
	ws.server.Delete("retailer/:id", ws.delete)
	ws.server.Get("retailer/:id/version", ws.getRetailerVersionList)
	ws.server.Get("retailer/:id/version/:version_id", ws.getRetailerVersion)
	ws.server.Delete("retailer/:id/version/:version_id", ws.deleteRetailerVersion)
}
