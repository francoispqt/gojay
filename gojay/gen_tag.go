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

func getGojayTagValue(tags *ast.BasicLit, tagName string) (*structtag.Tag, error) {
	t, err := structtag.Parse(tags.Value[1 : len(tags.Value)-1])
	if err != nil {
		return nil, err
	}

	v, err := t.Get(tagName)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func hasTagUnmarshalHide(tags *ast.BasicLit, tagName string) bool {
	v, err := getGojayTagValue(tags, tagName)
	if err != nil {
		log.Print(err)
		return false
	}
	return (v.Name == unmarshalHideTag || v.Name == hideTag) || v.HasOption(unmarshalHideTag)
}

func hasTagMarshalHide(tags *ast.BasicLit, tagName string) bool {
	v, err := getGojayTagValue(tags, tagName)
	if err != nil {
		log.Print(err)
		return false
	}
	return (v.Name == marshalHideTag || v.Name == hideTag) || v.HasOption(marshalHideTag)
}

func hasTagOmitEmpty(tags *ast.BasicLit, tagName string) bool {
	v, err := getGojayTagValue(tags, tagName)
	if err != nil {
		log.Print(err)
		return false
	}
	return v.Name == omitEmptyTag || v.HasOption(omitEmptyTag)
}

func tagKeyName(tags *ast.BasicLit, tagName string) string {
	v, err := getGojayTagValue(tags, tagName)
	if err != nil {
		log.Print(err)
		return ""
	}
	if v.Name == hideTag || v.Name == unmarshalHideTag || v.Name == marshalHideTag {
		return ""
	}
	return v.Name
}
