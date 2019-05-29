package codegen

import (
	"github.com/viant/toolbox"
	"reflect"
	"strings"
	"sort"
)

func firstLetterToUppercase(text string) string {
	return strings.ToUpper(string(text[0:1])) + string(text[1:])
}

func firstLetterToLowercase(text string) string {
	return strings.ToLower(string(text[0:1])) + string(text[1:])
}

func extractReceiverAlias(structType string) string {
	var result = string(structType[0])
	for i := len(structType) - 1; i > 0; i-- {
		aChar := string(structType[i])
		lowerChar := strings.ToLower(aChar)
		if lowerChar != aChar {
			result = lowerChar
			break
		}
	}
	return strings.ToLower(result)
}

func getTagOptions(tag, key string) []string {
	if tag == "" {
		return nil
	}
	var structTag = reflect.StructTag(strings.Replace(tag, "`", "", len(tag)))
	options, ok := structTag.Lookup(key)
	if !ok {
		return nil
	}
	return strings.Split(options, ",")
}

func getSliceHelperTypeName(typeName string, isPointer bool) string {
	if typeName == "" {
		return ""
	}

	var pluralName = firstLetterToUppercase(typeName) + "s"
	if isPointer {
		pluralName += "Ptr"
	}
	return strings.Replace(pluralName, ".", "", -1)
}

func isSkipable(options *Options, field *toolbox.FieldInfo) bool {
	if options := getTagOptions(field.Tag, options.TagName); len(options) > 0 {
		for _, candidate := range options {
			if candidate == "-" {
				return true
			}
		}
	}
	return false
}

func wrapperIfNeeded(text, wrappingChar string) string {
	if strings.HasPrefix(text, wrappingChar) {
		return text
	}
	return wrappingChar + text + wrappingChar
}

func getPoolName(typeName string) string {
	typeName = strings.Replace(typeName, "*", "", 1)
	return strings.Replace(typeName+"Pool", ".", "", -1)
}

func getJSONKey(options *Options, field *toolbox.FieldInfo) string {
	var key = field.Name
	if field.Tag != "" {
		if options := getTagOptions(field.Tag, options.TagName); len(options) > 0 {
			key = options[0]
		}
	}
	return key
}

func normalizeTypeName(typeName string) string {
	return strings.Replace(typeName, "*", "", strings.Count(typeName, "*"))
}

func sortedKeys(m map[string]string) ([]string) {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}