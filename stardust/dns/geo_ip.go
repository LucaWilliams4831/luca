package dns

import (
	"github.com/oschwald/geoip2-golang"
<<<<<<< HEAD
	"github.com/sipt/panda141035/assets"
	"github.com/sipt/panda141035/log"
=======
	"github.com/sipt/shuttle/assets"
	"github.com/sipt/shuttle/log"
>>>>>>> ba7fba0 (built with no error)
	"net"
)

var geoipDB *geoip2.Reader

func InitGeoIP(dbFile string) error {
	var err error
	geoipFileBytes, err := assets.ReadFile(dbFile)
	if err != nil {
		log.Logger.Errorf("[GeoIP] read failed [%v]", err)
		return err
	}
	geoipDB, err = geoip2.FromBytes(geoipFileBytes)
	if err != nil {
		return err
	}
	return nil
}

func GeoLookUp(ip string) (countryCode string) {
	if geoipDB == nil {
		return
	}
	netIP := net.ParseIP(ip)
	if netIP == nil {
		return
	}
	country, err := geoipDB.Country(netIP)
	if err == nil && country != nil {
		log.Logger.Debugf("[GeoIP] lookup [%s] country -> [%s]", ip, country.Country.IsoCode)
		return country.Country.IsoCode
	}
	log.Logger.Debugf("[GeoIP] lookup [%s] country failed: %s", ip, err.Error())
	return
}

func CloseGeoDB() error {
	return geoipDB.Close()
}
