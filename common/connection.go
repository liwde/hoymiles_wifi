package common

import (
	"fmt"
	"net"
	"time"
)

const (
	ConnectionType = "tcp"
)

type ConnectionData struct {
	Host string
	Port int32
	Uri  string
}

type Connection interface {
	SendRequest(command []byte, message []byte) []byte
}

// newConnection is unused at the moment
func newConnection(host string, port int32) *ConnectionData {
	var uri string = fmt.Sprintf("tcp://%s:%d", host, port)

	return &ConnectionData{host, port, uri}
}

// Should SendRequest match the `Connection` interface?
// If yes, function signatur of the interface should match the function signature of the implementation
func (connData *ConnectionData) SendRequest(command []byte, message []byte, sequence uint16) ([]byte, error) {

	tcpServer, err := net.ResolveTCPAddr(ConnectionType, connData.Uri)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP(ConnectionType, nil, tcpServer)
	if err != nil {
		return nil, err
	}

	// Set Connection TimeOut
	t := time.Now().Add(time.Second * 5)
	err = conn.SetReadDeadline(t)
	if err != nil {
		return nil, err
	}

	// `var` keyword can be dropped, when assignment is done via `:=`
	// as the type is already defined (either by declaration or return type of a function)
	var requestData = &RequestData{
		Header:   command,
		Sequence: sequence,
		Message:  message,
	}
	var messageData = GetByteMessage(requestData) //send message

	_, err = conn.Write(messageData)
	if err != nil {
		return nil, err
	}

	// receive answer
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		return nil, err
	}

	responseMessage, err := GetResponseMessage(received)
	if err != nil {
		return nil, err
	}

	err = conn.Close()
	if err != nil {
		return nil, err
	}

	return responseMessage.Message, nil
}
