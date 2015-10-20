package marketplace

import (
	"fmt"
)

var EndPoints = map[string]string{
	"A2EUQ1WTGCTBG2": "mws.amazonservices.ca",
	"ATVPDKIKX0DER":  "mws.amazonservices.com",
	"A1PA6795UKMFR9": "mws-eu.amazonservices.com",
	"A1RKKUPIHCS9HS": "mws-eu.amazonservices.com",
	"A13V1IB3VIYZZH": "mws-eu.amazonservices.com",
	"A21TJRUUN4KGV":  "mws.amazonservices.in",
	"APJ6JRA9NG5V4":  "mws-eu.amazonservices.com",
	"A1F83G8C2ARO7P": "mws-eu.amazonservices.com",
	"A1VC38T7YXB528": "mws.amazonservices.jp",
	"AAHKV2X7AFYLW":  "mws.amazonservices.com.cn",
}

var MarketPlaceIds = map[string]string{
	"CA": "A2EUQ1WTGCTBG2",
	"US": "ATVPDKIKX0DER",
	"DE": "A1PA6795UKMFR9",
	"ES": "A1RKKUPIHCS9HS",
	"FR": "A13V1IB3VIYZZH",
	"IN": "A21TJRUUN4KGV",
	"IT": "APJ6JRA9NG5V4",
	"UK": "A1F83G8C2ARO7P",
	"JP": "A1VC38T7YXB528",
	"CN": "AAHKV2X7AFYLW",
}

type MarketplaceError struct {
	errorType string
	value     string
}

func (e MarketplaceError) Error() string {
	return fmt.Sprintf("Invalid %v: %v", e.errorType, e.value)
}

type MarketPlace struct {
	Region   string
	Id       string
	EndPoint string
}

func New(region string) (*MarketPlace, error) {
	mp := MarketPlace{Region: region}

	marketPlaceId, idError := mp.MarketPlaceId()
	if idError != nil {
		return nil, idError
	}
	mp.Id = marketPlaceId

	endPoint, endPointError := mp.MarketPlaceEndPoint()
	if endPointError != nil {
		return nil, endPointError
	}
	mp.EndPoint = endPoint
	return &mp, nil
}

// Endpoint get the MWS end point for the region.
func (mp *MarketPlace) MarketPlaceEndPoint() (string, error) {
	if mp.EndPoint != "" {
		return mp.EndPoint, nil
	}
	if val, ok := EndPoints[mp.Id]; ok {
		return val, nil
	}
	return "", MarketplaceError{"marketplace id", mp.Id}
}

// MarketPlaceID get the marketpalce id for the region.
func (mp *MarketPlace) MarketPlaceId() (string, error) {
	if mp.Id != "" {
		return mp.Id, nil
	}
	if val, ok := MarketPlaceIds[mp.Region]; ok {
		return val, nil
	}
	return "", MarketplaceError{"region", mp.Region}
}

// Encoding get the ecoding for file upload and parsing
// TODO add encoding for JP.
func Encoding(region string) string {
	switch region {
	case "CN":
		return "UTF-16"
	default:
		return "ISO-8859-1"
	}
}
