package repository

const (
	listRetailer = `select id, name, address_city,address_street,address_house,owner_first_name,owner_last_name,open_time,close_time,created_at,actor max(version) 
	from retailer where deleted_at==null order by id`
	createRetailer = `INSERT INTO retailer (id, name, address_city,address_street,address_house,owner_first_name,owner_last_name,open_time,close_time,created_at,actor) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?)`

	updateRetailer        = `UPDATE `
	getRetailerByID       = `err = db.Get(&p, "SELECT * FROM retailer deleted_at==null LIMIT 1")`
	getReteilerVersion    = `err = db.Get(&p, "SELECT * FROM retailer deleted_at==null LIMIT 1")`
	deleteRetailer        = `UPDATE WHERE id=? deleted_at=?`
	deleteRetailerVersion = `UPDATE WHERE id=? deleted_at=?`
)
