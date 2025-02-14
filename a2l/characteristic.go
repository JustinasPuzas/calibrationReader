package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Characteristic struct {
	Name              string
	NameSet           bool
	LongIdentifier    string
	LongIdentifierSet bool
	Type              TypeEnum
	TypeSet           bool
	Address           string
	AddressUint32     uint32
	AddressSet        bool
	//Deposit is the identifier of the corresponding record layout
	Deposit       string
	DepositSet    bool
	Encoding      encodingEnum
	MaxDiff       float64
	MaxDiffSet    bool
	Conversion    string
	ConversionSet bool
	LowerLimit    float64
	LowerLimitSet bool
	UpperLimit    float64
	UpperLimitSet bool
	Annotation    []annotation
	AxisDescr     []axisDescr
	BitMask       bitMask
	//byteOrder can be used to overwrite the standard byte order defined in mod par
	ByteOrder               ByteOrder
	CalibrationAccess       calibrationAccessEnum
	ComparisonQuantity      comparisonQuantity
	DependentCharacteristic []DependentCharacteristic
	Discrete                discreteKeyword
	DisplayIdentifier       DisplayIdentifier
	EcuAddressExtension     ecuAddressExtension
	ExtendedLimits          extendedLimits
	Format                  format
	FunctionList            []FunctionList
	GuardRails              guardRailsKeyword
	IfData                  []IfData
	MapList                 []MapList
	MatrixDim               MatrixDim
	MaxRefresh              MaxRefresh
	ModelLink               modelLink
	Number                  Number
	PhysUnit                physUnit
	ReadOnly                readOnlyKeyword
	RefMemorySegment        refMemorySegment
	StepSize                StepSize
	SymbolLink              symbolLink
	VirtualCharacteristic   []VirtualCharacteristic
}

func parseCharacteristic(tok *tokenGenerator) (Characteristic, error) {
	c := Characteristic{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAnnotationToken:
			var buf annotation
			buf, err = parseAnnotation(tok)
			if err != nil {
				log.Err(err).Msg("characteristic annotation could not be parsed")
				break forLoop
			}
			c.Annotation = append(c.Annotation, buf)
			log.Info().Msg("characteristic annotation successfully parsed")
		case beginAxisDescrToken:
			var buf axisDescr
			buf, err = parseAxisDescr(tok)
			if err != nil {
				log.Err(err).Msg("characteristic axisDescr could not be parsed")
				break forLoop
			}
			c.AxisDescr = append(c.AxisDescr, buf)
			log.Info().Msg("characteristic axisDescr successfully parsed")
		case bitMaskToken:
			c.BitMask, err = parseBitMask(tok)
			if err != nil {
				log.Err(err).Msg("characteristic bitMask could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic bitMask successfully parsed")
		case byteOrderToken:
			c.ByteOrder, err = parseByteOrder(tok)
			if err != nil {
				log.Err(err).Msg("characteristic byteOrder could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic byteOrder successfully parsed")
		case calibrationAccessToken:
			c.CalibrationAccess, err = parseCalibrationAccessEnum(tok)
			if err != nil {
				log.Err(err).Msg("characteristic calibrationAccess could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic calibrationAccess successfully parsed")
		case comparisonQuantityToken:
			c.ComparisonQuantity, err = parseComparisonQuantity(tok)
			if err != nil {
				log.Err(err).Msg("characteristic comparisonQuantity could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic comparisonQuantity successfully parsed")
		case beginDependentCharacteristicToken:
			var buf DependentCharacteristic
			buf, err = parseDependentCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("characteristic dependentCharacteristic could not be parsed")
				break forLoop
			}
			c.DependentCharacteristic = append(c.DependentCharacteristic, buf)
			log.Info().Msg("characteristic dependentCharacteristic successfully parsed")
		case discreteToken:
			c.Discrete, err = parseDiscrete(tok)
			if err != nil {
				log.Err(err).Msg("characteristic discrete could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic discrete successfully parsed")
		case displayIdentifierToken:
			c.DisplayIdentifier, err = parseDisplayIdentifier(tok)
			if err != nil {
				log.Err(err).Msg("characteristic displayIdentifier could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic displayIdentifier successfully parsed")
		case encodingToken:
			c.Encoding, err = parseEncodingEnum(tok)
			if err != nil {
				log.Err(err).Msg("characteristic encoding could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic encoding successfully parsed")
		case ecuAddressExtensionToken:
			c.EcuAddressExtension, err = parseECUAddressExtension(tok)
			if err != nil {
				log.Err(err).Msg("characteristic ecuAddressExtension could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic ecuAddressExtension successfully parsed")
		case extendedLimitsToken:
			c.ExtendedLimits, err = parseExtendedLimits(tok)
			if err != nil {
				log.Err(err).Msg("characteristic bufExtendedLimits could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic extendedLimits successfully parsed")
		case formatToken:
			c.Format, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("characteristic format could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic format successfully parsed")
		case beginFunctionListToken:
			var buf FunctionList
			buf, err = parseFunctionList(tok)
			if err != nil {
				log.Err(err).Msg("characteristic functionList could not be parsed")
				break forLoop
			}
			c.FunctionList = append(c.FunctionList, buf)
			log.Info().Msg("characteristic functionList successfully parsed")
		case guardRailsToken:
			c.GuardRails, err = parseGuardRails(tok)
			if err != nil {
				log.Err(err).Msg("characteristic guardRails could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic guardRails successfully parsed")
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("characteristic ifData could not be parsed")
				break forLoop
			}
			c.IfData = append(c.IfData, buf)
			log.Info().Msg("characteristic ifData successfully parsed")
		case beginMapListToken:
			var buf MapList
			buf, err = parseMapList(tok)
			if err != nil {
				log.Err(err).Msg("characteristic mapList could not be parsed")
				break forLoop
			}
			c.MapList = append(c.MapList, buf)
			log.Info().Msg("characteristic mapList successfully parsed")
		case matrixDimToken:
			c.MatrixDim, err = parseMatrixDim(tok)
			if err != nil {
				log.Err(err).Msg("characteristic matrixDim could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic matrixDim successfully parsed")
		case maxRefreshToken:
			c.MaxRefresh, err = parseMaxRefresh(tok)
			if err != nil {
				log.Err(err).Msg("characteristic maxRefresh could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic maxRefresh successfully parsed")
		case modelLinkToken:
			c.ModelLink, err = parseModelLink(tok)
			if err != nil {
				log.Err(err).Msg("measurement modelLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement modelLink successfully parsed")
		case numberToken:
			c.Number, err = parseNumber(tok)
			if err != nil {
				log.Err(err).Msg("characteristic number could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic number successfully parsed")
		case physUnitToken:
			c.PhysUnit, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("characteristic physUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic physUnit successfully parsed")
		case readOnlyToken:
			c.ReadOnly, err = parseReadOnly(tok)
			if err != nil {
				log.Err(err).Msg("characteristic readOnly could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic readOnly successfully parsed")
		case refMemorySegmentToken:
			c.RefMemorySegment, err = parseRefMemorySegment(tok)
			if err != nil {
				log.Err(err).Msg("characteristic refMemorySegment could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic refMemorySegment successfully parsed")
		case stepSizeToken:
			c.StepSize, err = parseStepSize(tok)
			if err != nil {
				log.Err(err).Msg("characteristic stepSize could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic stepSize successfully parsed")
		case symbolLinkToken:
			c.SymbolLink, err = parseSymbolLink(tok)
			if err != nil {
				log.Err(err).Msg("characteristic symbolLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic symbolLink successfully parsed")
		case beginVirtualCharacteristicToken:
			var buf VirtualCharacteristic
			buf, err = parseVirtualCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("characteristic virtualCharacteristic could not be parsed")
				break forLoop
			}
			c.VirtualCharacteristic = append(c.VirtualCharacteristic, buf)
			log.Info().Msg("characteristic virtualCharacteristic successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("characteristic could not be parsed")
				break forLoop
			} else if tok.current() == endCharacteristicToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("characteristic could not be parsed")
				break forLoop
			} else if !c.NameSet {
				c.Name = tok.current()
				c.NameSet = true
				log.Info().Msg("characteristic name successfully parsed")
			} else if !c.LongIdentifierSet {
				c.LongIdentifier = tok.current()
				c.LongIdentifierSet = true
				log.Info().Msg("characteristic longIdentifier successfully parsed")
			} else if !c.TypeSet {
				c.Type, err = parseTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("characteristic type could not be parsed")
					break forLoop
				}
				c.TypeSet = true
				log.Info().Msg("characteristic type successfully parsed")
			} else if !c.AddressSet {
				c.Address = tok.current()
				c.AddressSet = true
				log.Info().Msg("characteristic Address successfully parsed")
			} else if !c.DepositSet {
				c.Deposit = tok.current()
				c.DepositSet = true
				log.Info().Msg("characteristic deposit successfully parsed")
			} else if !c.MaxDiffSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("characteristic maxDiff could not be parsed")
					break forLoop
				}
				c.MaxDiff = buf
				c.MaxDiffSet = true
				log.Info().Msg("characteristic maxDiff successfully parsed")
			} else if !c.ConversionSet {
				c.Conversion = tok.current()
				c.ConversionSet = true
				log.Info().Msg("characteristic conversion successfully parsed")
			} else if !c.LowerLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("characteristic lowerLimit could not be parsed")
					break forLoop
				}
				c.LowerLimit = buf
				c.LowerLimitSet = true
				log.Info().Msg("characteristic lowerLimit successfully parsed")
			} else if !c.UpperLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("characteristic upperLimit could not be parsed")
					break forLoop
				}
				c.UpperLimit = buf
				c.UpperLimitSet = true
				log.Info().Msg("characteristic upperLimit successfully parsed")
			}
		}
	}
	return c, err
}
