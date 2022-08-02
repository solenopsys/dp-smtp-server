package main

type TestModel struct {
	services  map[uint16]string
	clientsId uint16
}

func createModel() TestModel {
	return TestModel{services: nil}
}

func (m TestModel) addService(port int, service string) {

}

func (m TestModel) addClient(clientName string, key string) {

}

func (m TestModel) sendMessage(s string, id int, id2 int, body []byte) {

}
