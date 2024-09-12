package hoymiles_wifi

import (
	"google.golang.org/protobuf/proto"
	"hoymiles_wifi/hoymiles/models"
	"time"
)

func UpdateState(conn *Connection) *models.RealDataNewResDTO {

	header := []byte{'H', 'M', 0xa3, 0x11}

	dto := &models.RealDataNewReqDTO{}
	dto.Time = int32(time.Now().Unix())

	request_as_bytes, err := proto.Marshal(dto)
	if err != nil {
		return nil
	}

	result := &models.RealDataNewResDTO{}
	err = proto.Unmarshal(received, result)
	if err != nil {
		return nil
	}

	return result
}
