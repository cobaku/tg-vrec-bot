package dto

import "encoding/xml"

type YandexXmlBody struct {
	XMLName  xml.Name                      `xml:"recognitionResults"`
	Variants []YandexXmlRecognitionVariant `xml:"variant"`
}

type YandexXmlRecognitionVariant struct {
	Key   string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Value string `xml:",chardata"`
}
