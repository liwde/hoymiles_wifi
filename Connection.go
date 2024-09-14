package hoymiles_wifi

import (
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/BLun78/hoymiles_wifi/common"
	"github.com/BLun78/hoymiles_wifi/hoymiles/models"
)

func UpdateState(conn *common.Connection) *models.RealDataNewResDTO {

	// `header` is defined, but unused
	// Isn't the byte array not the same like common.CMD_HEADER ?
	header := []byte{'H', 'M', 0xa3, 0x11}

	dto := &models.RealDataNewReqDTO{}

	// int32 would not be Year 2038 safe
	// See https://en.wikipedia.org/wiki/Year_2038_problem
	// Not 100% sure if the models are self-defined or provided by hoymiles
	dto.Time = int32(time.Now().Unix())

	// `request_as_bytes` is defined, but unused
	request_as_bytes, err := proto.Marshal(dto)
	if err != nil {
		return nil
	}

	result := &models.RealDataNewResDTO{}
	// `received` is not defined/initialized
	err = proto.Unmarshal(received, result)
	if err != nil {
		return nil
	}

	return result
}
