package common

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/basvdlei/gotsmart/crc16"
)

const (
	HeaderLength = 10
)

type RequestData struct {
	Header   []byte
	Sequence uint16
	Message  []byte
}

type ResponseData struct {
	Header        []byte
	Sequence      uint16
	Message       []byte
	IsCrc16Valid  bool
	IsLengthValid bool
	IsValid       bool
}

func GetByteMessage(request *RequestData) []byte {
	// prepare data
	var sequence uint16 = request.Sequence & 0xFFFF
	var crc uint16 = crc16.Checksum(request.Message)
	var length = uint16(len(request.Message) + HeaderLength)

	// The sequence, crc, length must order by BigEndian
	sequenceBytes := toBigEndianByteArray16(sequence)
	crc16Bytes := toBigEndianByteArray16(crc)
	lengthBytes := toBigEndianByteArray16(length)

	// prepare message
	message := make([]byte, length)
	message = append(message, request.Header...)
	message = append(message, sequenceBytes[0], sequenceBytes[1])
	message = append(message, crc16Bytes[0], crc16Bytes[1])
	message = append(message, lengthBytes[0], lengthBytes[1])
	message = append(message, request.Message...)
	return message
}

func GetResponseMessage(response []byte) (*ResponseData, error) {
	var err error = nil

	if len(response) < HeaderLength {
		err = errors.New("Buffer is too short - message format is invalid of response.")
		return nil, err
	}

	// read data
	var header []byte = response[:4]
	// The sequence, crc, length must order by LittleEndian
	var sequence uint16 = binary.LittleEndian.Uint16(response[4:6])
	var crc16Response uint16 = binary.LittleEndian.Uint16(response[6:8])
	var length = binary.LittleEndian.Uint16(response[8:10])
	var message []byte = response[10:]

	// create new crc16 to validate message
	var crc = crc16.Checksum(message)

	// Checks
	var crcCheck = crc == crc16Response
	if !crcCheck {
		err = fmt.Errorf("CRC16 mismatch: %d != %d", crc16Response, crc)
		return nil, err
	}
	var lengthCheck = (length - HeaderLength) == uint16(len(message))
	if !lengthCheck {
		err = fmt.Errorf("Length mismatch: %d != %d", len(message), HeaderLength)
		return nil, err
	}

	// prepare response
	var result = &ResponseData{
		Header:        header,
		Sequence:      sequence,
		Message:       message,
		IsCrc16Valid:  crcCheck,
		IsLengthValid: lengthCheck,
		IsValid:       crcCheck && lengthCheck,
	}
	return result, err
}

func toBigEndianByteArray16(i uint16) (arr [2]byte) {
	binary.BigEndian.PutUint16(arr[0:2], i)
	return
}
