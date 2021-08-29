package services

import (
	"net"

	geoip2 "github.com/oschwald/geoip2-golang"
)

type (
	GeoService interface {
		GetIPInfo(ip string) (*GeoInfo, error)
	}

	GeoIP2Service struct {
		lang   string
		cityDB *geoip2.Reader
	}

	GeoInfo struct {
		IP       string `json:"ip,omitempty"`
		City     string `json:"city,omitempty"`
		Region   string `json:"region,omitempty"`
		Location struct {
			AccuracyRadius uint16  `json:"accuracy_radius,omitempty"`
			Latitude       float64 `json:"latitude,omitempty"`
			Longitude      float64 `json:"longitude,omitempty"`
			MetroCode      uint    `json:"metro_code,omitempty"`
			TimeZone       string  `json:"time_zone,omitempty"`
		} `json:"location,omitempty"`
	}
)

func NewGeoService() (*GeoIP2Service, error) {
	cityDB, err := geoip2.Open("data/GeoLite2-City.mmdb")
	if err != nil {
		return nil, err
	}

	return &GeoIP2Service{
		lang:   "en",
		cityDB: cityDB,
	}, nil
}

func (s *GeoIP2Service) Close() {
	if s != nil {
		s.Close()
	}
}

func (s *GeoIP2Service) GetIPInfo(ip string) (*GeoInfo, error) {
	ipAddress := net.ParseIP(ip)
	city, err := s.cityDB.City(ipAddress)
	if err != nil {
		return nil, err
	}

	return &GeoInfo{
		IP:     ip,
		City:   city.City.Names[s.lang],
		Region: city.Continent.Names[s.lang],
		Location: struct {
			AccuracyRadius uint16  `json:"accuracy_radius,omitempty"`
			Latitude       float64 `json:"latitude,omitempty"`
			Longitude      float64 `json:"longitude,omitempty"`
			MetroCode      uint    `json:"metro_code,omitempty"`
			TimeZone       string  `json:"time_zone,omitempty"`
		}(city.Location),
	}, nil
}
