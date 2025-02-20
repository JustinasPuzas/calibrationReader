package calibrationReader

import (
	"errors"

	"github.com/JustinasPuzas/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getOffsetX retrieves the offset operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getOffsetX(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.OffsetX.DatatypeSet {
		err := errors.New("offsetX datatype not set")
		log.Err(err).Msg("could not retrieve offsetX value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.OffsetX.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve offsetX value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.OffsetX.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve offsetX value")
		return 0, err
	}
	*curPos += uint32(rl.OffsetX.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getOffsetY retrieves the offset operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getOffsetY(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.OffsetY.DatatypeSet {
		err := errors.New("offsetY datatype not set")
		log.Err(err).Msg("could not retrieve offsetY value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.OffsetY.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve offsetY value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.OffsetY.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve offsetY value")
		return 0, err
	}
	*curPos += uint32(rl.OffsetY.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getOffsetZ retrieves the offset operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getOffsetZ(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.OffsetZ.DatatypeSet {
		err := errors.New("offsetZ datatype not set")
		log.Err(err).Msg("could not retrieve offsetZ value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.OffsetZ.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve offsetZ value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.OffsetZ.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve offsetZ value")
		return 0, err
	}
	*curPos += uint32(rl.OffsetZ.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getOffset4 retrieves the offset operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getOffset4(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.Offset4.DatatypeSet {
		err := errors.New("offset4 datatype not set")
		log.Err(err).Msg("could not retrieve offset4 value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.Offset4.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve offset4 value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.Offset4.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve offset4 value")
		return 0, err
	}
	*curPos += uint32(rl.Offset4.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getOffset5 retrieves the offset operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getOffset5(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.Offset5.DatatypeSet {
		err := errors.New("offset5 datatype not set")
		log.Err(err).Msg("could not retrieve offset5 value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.Offset5.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve offset5 value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.Offset5.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve offset5 value")
		return 0, err
	}
	*curPos += uint32(rl.Offset5.Datatype.GetDatatypeLength())
	return int64(val), err
}
