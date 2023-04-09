package xml_data

import (
	"encoding/xml"

	"github.com/abhisekp/conf42-golang-2023-building-cli-app/pkg/person"
)

type Root struct {
	_            struct{}
	XMLName      xml.Name `xml:"Persona"`
	XMLSchema    string   `xml:"xsi:noNamespaceSchemaLocation,attr"`
	XMLNamespace string   `xml:"xmlns:xsi,attr"`
	Name         string   `xml:"Name"`
	DateCreated  string   `xml:"DateCreated"`

	Items []*person.Person `xml:"Person"`
}
