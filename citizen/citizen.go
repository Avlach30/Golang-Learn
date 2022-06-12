package citizen

//* Define function which is returned a slice with element of map of multiple keys
func Citizens() []map[string]string {

	//* Declataion of slice with element of map which is have multiple key-value pair
	citizens := []map[string]string {
		{"name": "arnold", "address": "diponegoro street 45"},
		{"name": "bobby", "address": "soekarno street 8"},
		{"name": "charlie", "address": "bukittingi, 87"},
		{"name": "doddy", "address": "asia afrika, 7"},
		{"name": "evan", "address": "achmad yani, 88"},
	}

	return citizens
}