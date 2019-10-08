package page

var defaultSize int32 = 0

// ConvertToMongo - convert page , size to mongo skip and limit
func ConvertToMongo(page, size int32) (skip int, limit int) {
	if page < 1 {
		page = 1
	}

	if size == 0 {
		size = defaultSize
	}

	skip = int(size * (page - 1))
	limit = int(size)
	return
}

// SetDefaultSize -  set the page defualt size
func SetDefaultSize(size int32) {
	defaultSize = size
}
