package repository

var (
	listRetailer string = `select id, name, address_city,address_street,address_house,owner_first_name,owner_last_name,open_time,close_time,created_at,actor max(version) 
	from retailer 
	where deleted_at is not null 
	order by id`
	createRetailer string = `INSERT INTO retailer (id, name, address_city, address_street, address_house, owner_first_name, owner_last_name, open_time, close_time, created_at, actor) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	updateRetailer        string = `UPDATE `
	getRetailerByID       string = `err = db.Get(&p, "SELECT * FROM retailer deleted_at==null LIMIT 1")`
	getReteilerVersion    string = `err = db.Get(&p, "SELECT * FROM retailer deleted_at==null LIMIT 1")`
	deleteRetailer        string = `UPDATE WHERE id=? deleted_at=?`
	deleteRetailerVersion string = `UPDATE WHERE id=? deleted_at=?`
)

//INSERT INTO retailer (id, name, address_city, address_street, address_house, owner_first_name, owner_last_name, open_time, close_time, created_at, actor) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)

//INSERT INTO retailer (id, name, address_city,address_street,address_house,owner_first_name,owner_last_name,open_time,close_time,created_at,actor) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)

// INSERT INTO retailer (
// 	id,
// 	name,
// 	address_city,
// 	address_street,
// 	address_house,
// 	owner_first_name,
// 	owner_last_name,
// 	open_time,
// 	close_time,
// 	created_at,
// 	actor) VALUES (?,?,?,?,?,?,?,?,?,?,?)
