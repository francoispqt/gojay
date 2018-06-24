package main

import (
	"go/ast"
	"log"

	"github.com/fatih/structtag"
)

const gojayTag = "gojay"
const hideTag = "-"
const unmarshalHideTag = "-u"
const marshalHideTag = "-m"
const omitEmptyTag = "omitempty"

func getGojayTagValue(tags *ast.BasicLit) (*structtag.Tag, error) {
	t, err := structtag.Parse(tags.Value[1 : len(tags.Value)-1])
	if err != nil {
		return nil, err
	}
	v, err := t.Get(gojayTag)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func hasTagUnmarshalHide(tags *ast.BasicLit) bool {
	v, err := getGojayTagValue(tags)
	if err != nil {
		log.Print(err)
		return false
	}
	return (v.Name == unmarshalHideTag || v.Name == hideTag) || v.HasOption(unmarshalHideTag)
}

func hasTagMarshalHide(tags *ast.BasicLit) bool {
	v, err := getGojayTagValue(tags)
	if err != nil {
		log.Print(err)
		return false
	}
	return (v.Name == marshalHideTag || v.Name == hideTag) || v.HasOption(marshalHideTag)
}

func hasTagOmitEmpty(tags *ast.BasicLit) bool {
	v, err := getGojayTagValue(tags)
	if err != nil {
		log.Print(err)
		return false
	}
	return v.Name == omitEmptyTag || v.HasOption(omitEmptyTag)
}

func tagKeyName(tags *ast.BasicLit) string {
	v, err := getGojayTagValue(tags)
	if err != nil {
		log.Print(err)
		return ""
	}
	if v.Name == hideTag || v.Name == unmarshalHideTag || v.Name == marshalHideTag {
		return ""
	}
	return v.Name
}
