package main

const (
	apiUrl = "http://34.77.36.161"
	apiPort = 3000
	apiPort2 = 3941
)
func main() {
	key:= getRightPort(apiPort, apiUrl)
	postURL(apiUrl, key, apiPort2, filepath)
}
