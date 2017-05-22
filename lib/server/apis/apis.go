package apis

var apiMap map[string]func 
func initialize() {
	apiMap["/"]	= homeHandler
	apiMap["/letsee"] = letseeHandler
}