package database

var GlobalDB DB

func InitGlobalDB() {
	GlobalDB = New()
}

type DB struct {
	data map[string]string
}

func New() DB {
	return DB{
		data: map[string]string{},
	}
}

func (db *DB) Set(key string, value string) {
	db.data[key] = value
}

func (db *DB) GetAll() []map[string]string {
	var arr []map[string]string

	for k, v := range db.data {
		arr = append(arr, map[string]string{
			"Key":   k,
			"Value": v,
		})
	}

	return arr
}

func (db *DB) Delete(key string) {
	delete(db.data, key)
}
