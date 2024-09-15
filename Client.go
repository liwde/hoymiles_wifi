package hoymiles_wifi

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/BLun78/hoymiles_wifi/common"
	"github.com/BLun78/hoymiles_wifi/hoymiles/models"
	"google.golang.org/protobuf/proto"
	"net"
	"time"
)

type (
	ClientData struct {
		connectionData *ConnectionData
		connection     net.Conn
	}
	ConnectionData struct {
		Host     string
		Port     int32
		Uri      string
		Sequence uint16
	}
	requestData struct {
		Header   []byte
		Sequence uint16
		Message  []byte
	}
	responseData struct {
		Header        []byte
		Sequence      uint16
		Message       []byte
		IsCrc16Valid  bool
		IsLengthValid bool
		IsValid       bool
	}
	Client interface {
		CloseConnection()
		GetRealDataNew(dto models.RealDataNewReqDTO) models.RealDataNewResDTO
	}
)

func NewClient(host string, port int32) *ClientData {
	connectionData := &ConnectionData{
		Host:     host,
		Port:     port,
		Uri:      fmt.Sprintf("%s:%d", host, port),
		Sequence: 0,
	}

	return &ClientData{connectionData: connectionData}
}

func (client *ClientData) GetRealDataNew(request *models.RealDataNewReqDTO) (*models.RealDataNewResDTO, error) {

	request.Offset = common.OFFSET
	var err error = nil
	var response proto.Message = &models.RealDataNewResDTO{}

	response, err = client.sendRequestProtobuf(common.CMD_REAL_RES_DTO, request, response)
	if err != nil {
		return nil, err
	}

	return response.(*models.RealDataNewResDTO), nil
}

func (client *ClientData) CloseConnection() error {
	err := client.connection.Close()

	if client != nil && client.connectionData != nil {
		client.connectionData.Sequence = 0
	}

	client.connection = nil
	return err
}

func (client *ClientData) sendRequestProtobuf(command []byte, requestMessage proto.Message, responseMessage proto.Message) (proto.Message, error) {
	request_as_byte, err := proto.Marshal(requestMessage)
	if err != nil {
		return nil, err
	}

	response, err := client.sendRequest(command, request_as_byte)
	if err != nil {
		return nil, err
	}

	err = proto.Unmarshal(response, responseMessage)
	if err != nil {
		return nil, err
	}

	return responseMessage, nil
}

func (client *ClientData) sendRequest(command []byte, message []byte) ([]byte, error) {

	err := client.createOrCheckConnection()
	if err != nil {
		return nil, err
	}

	client.connectionData.Sequence++

	requestData := &requestData{
		Header:   command,
		Sequence: client.connectionData.Sequence,
		Message:  message,
	}
	var messageData = getByteMessage(requestData) //send message

	_, err = client.connection.Write(messageData)
	if err != nil {
		return nil, err
	}

	// receive answer
	received := make([]byte, 1024)
	_, err = client.connection.Read(received)
	if err != nil {
		return nil, err
	}

	responseMessage, err := getResponseMessage(received)
	if err != nil {
		return nil, err
	}

	return responseMessage.Message, nil
}

func getByteMessage(request *requestData) []byte {
	// prepare data
	var sequence uint16 = request.Sequence & 0xFFFF
	var crc uint16 = common.CRC16(request.Message)
	var length = uint16(len(request.Message) + common.HEADER_LENGTH)

	// The sequence, crc, length must order by BigEndian
	sequenceBytes := toBigEndianByteArray16(sequence)
	crc16Bytes := toBigEndianByteArray16(crc)
	lengthBytes := toBigEndianByteArray16(length)

	// prepare message
	message := make([]byte, 0)
	message = append(message, common.CMD_HEADER[0], common.CMD_HEADER[1])
	message = append(message, request.Header[0], request.Header[1])
	message = append(message, sequenceBytes[0], sequenceBytes[1])
	message = append(message, crc16Bytes[0], crc16Bytes[1])
	message = append(message, lengthBytes[0], lengthBytes[1])
	message = append(message, request.Message...)
	return message
}

func getResponseMessage(response []byte) (*responseData, error) {
	var err error = nil

	if len(response) < common.HEADER_LENGTH {
		err = errors.New("Buffer is too short - message format is invalid of response.")
		return nil, err
	}

	// read data
	var header []byte = response[:4]
	// The sequence, crc, length must order by LittleEndian
	var sequence uint16 = binary.BigEndian.Uint16(response[4:6])
	var crc16Response uint16 = binary.BigEndian.Uint16(response[6:8])
	var length = binary.BigEndian.Uint16(response[8:10])
	var message []byte = response[10:length]

	// create new crc16 to validate message
	var crc = common.CRC16(message)

	// Checks
	var crcCheck = crc == crc16Response
	if !crcCheck {
		err = fmt.Errorf("CRC16 mismatch: %d != %d", crc16Response, crc)
		return nil, err
	}
	var lengthCheck = (length - common.HEADER_LENGTH) == uint16(len(message))
	if !lengthCheck {
		err = fmt.Errorf("Length mismatch: %d != %d", len(message), common.HEADER_LENGTH)
		return nil, err
	}

	// prepare response
	var result = &responseData{
		Header:        header,
		Sequence:      sequence,
		Message:       message,
		IsCrc16Valid:  crcCheck,
		IsLengthValid: lengthCheck,
		IsValid:       crcCheck && lengthCheck,
	}
	return result, err
}

// Review: Names returns are possible in Go, but very uncommon
// The idiomatic way would be to return the array directly
func toBigEndianByteArray16(i uint16) (arr [2]byte) {
	binary.BigEndian.PutUint16(arr[0:2], i)
	return
}

func (client *ClientData) createOrCheckConnection() error {
	if client.connection == nil {

		tcpServer, err := net.ResolveTCPAddr(common.CCONNECTION_TYPE, client.connectionData.Uri)
		if err != nil {
			return err
		}

		client.connection, err = net.DialTCP(common.CCONNECTION_TYPE, nil, tcpServer)
		if err != nil {
			return err
		}

		// Set Connection TimeOut
		t := time.Now().Add(time.Second * 15)
		err = client.connection.SetReadDeadline(t)
		if err != nil {
			return err
		}
	}
	return nil
}
