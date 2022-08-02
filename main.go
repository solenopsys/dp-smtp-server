package main

func main() {

}

func sleep() {

}

func testConnectService() {
	model := createModel()
	model.addService(2001, "service1")
	sleep()
	chekService()
}

func testConnectWs() {
	model := createModel()
	model.addClient("client1", "trueKey")
	sleep()
	chekClient()
}

func testAuthErorWs() {
	model := createModel()
	model.addClient("client1", "falseKey")
	sleep()
	chekClient()
}

func testMessageWsToService() {
	model := createModel()
	model.addService(2001, "service1")
	model.addClient("client1", "trueKey")
	const streamId = 212
	const funcId = 212
	body := []byte("")
	model.sendMessage("service1", streamId, funcId, body)
}

func testMessageWsToServiceAndResponse() {
	model := createModel()
	model.addService(2001, "service1")
	model.addClient("client1", "trueKey")
	const streamId = 212
	const funcId = 212
	body := []byte("")
	model.sendMessage("service1", streamId, funcId, body)
}
