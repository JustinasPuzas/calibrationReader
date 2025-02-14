package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type measurement struct {
	Name                string
	NameSet             bool
	LongIdentifier      string
	LongIdentifierSet   bool
	Datatype            DataTypeEnum
	DatatypeSet         bool
	Conversion          string
	ConversionSet       bool
	Resolution          uint16
	ResolutionSet       bool
	Accuracy            float64
	AccuracySet         bool
	LowerLimit          float64
	LowerLimitSet       bool
	UpperLimit          float64
	UpperLimitSet       bool
	Annotation          []annotation
	ArraySize           arraySize
	BitMask             bitMask
	BitOperation        bitOperation
	ByteOrder           ByteOrder
	Discrete            discreteKeyword
	DisplayIdentifier   DisplayIdentifier
	EcuAddress          ecuAddress
	EcuAddressExtension ecuAddressExtension
	ErrorMask           errorMask
	Format              format
	FunctionList        FunctionList
	IfData              []IfData
	Layout              layout
	MatrixDim           MatrixDim
	MaxRefresh          MaxRefresh
	ModelLink           modelLink
	PhysUnit            physUnit
	ReadWrite           readWriteKeyword
	RefMemorySegment    refMemorySegment
	SymbolLink          symbolLink
	Virtual             virtual
}

func parseMeasurement(tok *tokenGenerator) (measurement, error) {
	m := measurement{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAnnotationToken:
			var buf annotation
			buf, err = parseAnnotation(tok)
			if err != nil {
				log.Err(err).Msg("measurement annotation could not be parsed")
				break forLoop
			}
			m.Annotation = append(m.Annotation, buf)
			log.Info().Msg("measurement annotation successfully parsed")
		case arraySizeToken:
			m.ArraySize, err = parseArraySize(tok)
			if err != nil {
				log.Err(err).Msg("measurement arraySize could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement arraySize successfully parsed")
		case bitMaskToken:
			m.BitMask, err = parseBitMask(tok)
			if err != nil {
				log.Err(err).Msg("measurement bitMask could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement bitMask successfully parsed")
		case beginBitOperationToken:
			m.BitOperation, err = parseBitOperation(tok)
			if err != nil {
				log.Err(err).Msg("measurement bitOperation could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement bitOperation successfully parsed")
		case byteOrderToken:
			m.ByteOrder, err = parseByteOrder(tok)
			if err != nil {
				log.Err(err).Msg("measurement byteOrder could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement byteOrder successfully parsed")
		case discreteToken:
			m.Discrete, err = parseDiscrete(tok)
			if err != nil {
				log.Err(err).Msg("measurement discrete could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement discrete successfully parsed")
		case displayIdentifierToken:
			m.DisplayIdentifier, err = parseDisplayIdentifier(tok)
			if err != nil {
				log.Err(err).Msg("measurement displayIdentifier could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement displayIdentifier successfully parsed")
		case ecuAddressToken:
			m.EcuAddress, err = parseEcuAddress(tok)
			if err != nil {
				log.Err(err).Msg("measurement ecuAddress could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement ecuAddress successfully parsed")
		case ecuAddressExtensionToken:
			m.EcuAddressExtension, err = parseECUAddressExtension(tok)
			if err != nil {
				log.Err(err).Msg("measurement ecuAddressExtension could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement ecuAddressExtension successfully parsed")
		case errorMaskToken:
			m.ErrorMask, err = parseErrorMask(tok)
			if err != nil {
				log.Err(err).Msg("measurement errorMask could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement errorMask successfully parsed")
		case formatToken:
			m.Format, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("measurement format could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement format successfully parsed")
		case beginFunctionListToken:
			m.FunctionList, err = parseFunctionList(tok)
			if err != nil {
				log.Err(err).Msg("measurement functionList could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement functionList successfully parsed")
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("measurement ifData could not be parsed")
				break forLoop
			}
			m.IfData = append(m.IfData, buf)
			log.Info().Msg("measurement ifData successfully parsed")
		case layoutToken:
			m.Layout, err = parseLayout(tok)
			if err != nil {
				log.Err(err).Msg("measurement layout could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement layout successfully parsed")
		case matrixDimToken:
			m.MatrixDim, err = parseMatrixDim(tok)
			if err != nil {
				log.Err(err).Msg("measurement matrixDim could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement matrixDim successfully parsed")
			log.Info().Str("current token", tok.current()).Msg("measurement current token:")
		case maxRefreshToken:
			m.MaxRefresh, err = parseMaxRefresh(tok)
			if err != nil {
				log.Err(err).Msg("measurement maxRefresh could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement maxRefresh successfully parsed")
		case modelLinkToken:
			m.ModelLink, err = parseModelLink(tok)
			if err != nil {
				log.Err(err).Msg("measurement modelLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement modelLink successfully parsed")
		case physUnitToken:
			m.PhysUnit, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("measurement physUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement physUnit successfully parsed")
		case readWriteToken:
			m.ReadWrite, err = parseReadWrite(tok)
			if err != nil {
				log.Err(err).Msg("measurement readWrite could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement readWrite successfully parsed")
		case refMemorySegmentToken:
			m.RefMemorySegment, err = parseRefMemorySegment(tok)
			if err != nil {
				log.Err(err).Msg("measurement refMemorySegment could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement refMemorySegment successfully parsed")
		case symbolLinkToken:
			m.SymbolLink, err = parseSymbolLink(tok)
			if err != nil {
				log.Err(err).Msg("measurement symbolLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement symbolLink successfully parsed")
		case beginVirtualToken:
			m.Virtual, err = parseVirtual(tok)
			if err != nil {
				log.Err(err).Msg("measurement virtual could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement virtual successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("measurement could not be parsed")
				break forLoop
			} else if tok.current() == endMeasurementToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("measurement could not be parsed")
				break forLoop
			} else if !m.NameSet {
				m.Name = tok.current()
				m.NameSet = true
				log.Info().Msg("measurement name successfully parsed")
			} else if !m.LongIdentifierSet {
				m.LongIdentifier = tok.current()
				m.LongIdentifierSet = true
				log.Info().Msg("measurement longIdentifier successfully parsed")
			} else if !m.DatatypeSet {
				m.Datatype, err = parseDataTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("measurement datatype could not be parsed")
					break forLoop
				}
				m.DatatypeSet = true
				log.Info().Msg("measurement datatype successfully parsed")
			} else if !m.ConversionSet {
				m.Conversion = tok.current()
				m.ConversionSet = true
				log.Info().Msg("measurement conversion successfully parsed")
			} else if !m.ResolutionSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("measurement resolution could not be parsed")
					break forLoop
				}
				m.Resolution = uint16(buf)
				m.ResolutionSet = true
				log.Info().Msg("measurement resolution successfully parsed")
			} else if !m.AccuracySet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("measurement accuracy could not be parsed")
					break forLoop
				}
				m.Accuracy = buf
				m.AccuracySet = true
				log.Info().Msg("measurement accuracy successfully parsed")
			} else if !m.LowerLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("measurement lowerLimit could not be parsed")
					break forLoop
				}
				m.LowerLimit = buf
				m.LowerLimitSet = true
				log.Info().Msg("measurement lowerLimit successfully parsed")
			} else if !m.UpperLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("measurement upperLimit could not be parsed")
					break forLoop
				}
				m.UpperLimit = buf
				m.UpperLimitSet = true
				log.Info().Msg("measurement upperLimit successfully parsed")
			}
		}
	}
	return m, err
}
