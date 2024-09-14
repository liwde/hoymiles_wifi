package hoymiles_wifi

import (
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/BLun78/hoymiles_wifi/common"
	"github.com/BLun78/hoymiles_wifi/hoymiles/models"
)

func UpdateState(conn *common.Connection) *models.RealDataNewResDTO {

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
