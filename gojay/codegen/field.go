package codegen

import (
	"fmt"
	"github.com/viant/toolbox"
	"strings"
)

//Field represents a field.
type Field struct {
	Key                string
	Init               string
	OmitEmpty          string
	TimeLayout         string
	NullType           string
	Name               string
	Accessor           string
	Mutator            string
	Receiver           string //alias and type name
	Alias              string //object alias name
	Var                string //variable for this field
	Type               string
	RawType            string
	HelperType         string
	ComponentType      string
	RawComponentType   string
	IsPointerComponent bool

	PointerModifier     string //takes field pointer, "&" if field is not a pointer type
	DereferenceModifier string //take pointer value, i.e "*" if field has a pointer type

	ComponentPointerModifier     string //takes item pointer if needed,i.e
	ComponentDereferenceModifier string //de reference value if needed, i.e
	ComponentInitModifier        string //takes item pointer if type is not a pointer type
	ComponentInit                string //initialises component type

	DecodingMethod  string
	EncodingMethod  string
	PoolName        string //pool name associated with this field
	ResetDependency string
	Reset           string
	IsAnonymous     bool
	IsPointer       bool
	IsSlice         bool

	GojayMethod string
}

//NewField returns a new field
func NewField(owner *Struct, field *toolbox.FieldInfo, fieldType *toolbox.TypeInfo) (*Field, error) {
	typeName := normalizeTypeName(field.TypeName)
	var result = &Field{
		IsAnonymous:        field.IsAnonymous,
		Name:               field.Name,
		RawType:            field.TypeName,
		IsPointer:          field.IsPointer,
		Key:                getJSONKey(owner.options, field),
		Receiver:           owner.Alias + " *" + owner.TypeInfo.Name,
		Type:               typeName,
		Mutator:            owner.Alias + "." + field.Name,
		Accessor:           owner.Alias + "." + field.Name,
		ComponentType:      field.ComponentType,
		IsPointerComponent: field.IsPointerComponent,
		Var:                firstLetterToLowercase(field.Name),
		Init:               fmt.Sprintf("%v{}", typeName),
		TimeLayout:         "time.RFC3339",
		IsSlice:            field.IsSlice,
		PoolName:           getPoolName(field.TypeName),
		Alias:              owner.Alias,
		Reset:              "nil",
	}
	var err error
	if field.IsPointer {
		result.DereferenceModifier = "*"
		result.Init = "&" + result.Init
	} else {
		result.PointerModifier = "&"

	}
	if field.IsSlice {
		result.HelperType = getSliceHelperTypeName(field.ComponentType, field.IsPointerComponent)
		result.PoolName = getPoolName(field.ComponentType)
	} else if fieldType != nil {
		result.HelperType = getSliceHelperTypeName(fieldType.Name, field.IsPointerComponent)
	}

	if options := getTagOptions(field.Tag, "timeLayout"); len(options) > 0 {
		result.TimeLayout = wrapperIfNeeded(options[0], `"`)
	} else if options := getTagOptions(field.Tag, "timeFormat"); len(options) > 0 {
		result.TimeLayout = wrapperIfNeeded(toolbox.DateFormatToLayout(options[0]), `"`)
	}
	if strings.Contains(field.Tag, "omitempty") {
		result.OmitEmpty = "OmitEmpty"
	}
	if strings.Contains(field.Tag, "nullempty") {
		result.OmitEmpty = "NullEmpty"
	}

	if owner.options.PoolObjects {
		if field.IsPointer && !strings.HasSuffix(field.TypeName, ".Time") && !strings.Contains(field.TypeName, "sql.Null") {
			poolName := getPoolName(field.TypeName)
			result.Init = fmt.Sprintf(`%v.Get().(*%v)`, poolName, field.TypeName)
		}
	}

	encodingMethod := field.ComponentType
	if encodingMethod == "" {
		encodingMethod = result.Type
	}
	result.DecodingMethod = firstLetterToUppercase(encodingMethod)
	result.EncodingMethod = firstLetterToUppercase(encodingMethod)

	switch typeName {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		result.Reset = "0"
	case "float32", "float64":
		result.Reset = "0.0"
	case "string":
		result.Reset = `""`

	case "bool":
		result.Reset = "false"
	default:
		if field.IsSlice && owner.Type(field.ComponentType) != nil {
			var itemPointer = ""
			if !field.IsPointerComponent {
				itemPointer = "&"
			}
			result.ResetDependency, err = expandFieldTemplate(poolSliceInstanceRelease, struct {
				PoolName        string
				Accessor        string
				PointerModifier string
			}{PoolName: result.PoolName, Accessor: result.Accessor, PointerModifier: itemPointer})
			if err != nil {
				return nil, err
			}

		} else if field.IsPointer && fieldType != nil {
			result.ResetDependency, err = expandFieldTemplate(poolInstanceRelease, struct {
				PoolName string
				Accessor string
			}{PoolName: result.PoolName, Accessor: result.Accessor})
			if err != nil {
				return nil, err
			}
		}

	}
	if field.IsSlice || field.IsPointer {
		result.Reset = "nil"
	}

	if result.IsPointerComponent {
		result.ComponentInit = "&" + result.ComponentType + "{}"
		result.RawComponentType = "*" + result.ComponentType

		result.ComponentDereferenceModifier = "*"
		result.ComponentInitModifier = "&"

	} else {
		result.ComponentInit = result.ComponentType + "{}"
		result.RawComponentType = result.ComponentType

		result.ComponentPointerModifier = "&"
	}

	return result, nil
}
