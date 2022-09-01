package sysinfo

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type IP struct {
	Query string
}

// SysInfo saves the basic system information
type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	RAM      uint64 `bson:ram`
	Disk     uint64 `bson:disk`
}

func NewSysInfo() *SysInfo {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	//diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	//info.Disk = diskStat.Total / 1024 / 1024
	return info
}

func (sysInfo *SysInfo) ToString() string {
	return sysInfo.Hostname + "," + sysInfo.Platform + "," + sysInfo.CPU + "," + strconv.FormatUint(sysInfo.RAM, 10) + "," + strconv.FormatUint(sysInfo.Disk, 10)
}

func (sysInfo *SysInfo) ToHash() string {
	Hostname := HashValue(sysInfo.Hostname)
	Platform := HashValue(sysInfo.Platform)
	CPU := HashValue(sysInfo.CPU)
	RAM := HashValue(strconv.FormatUint(sysInfo.RAM, 10))
	Disk := HashValue(strconv.FormatUint(sysInfo.Disk, 10))
	return HashValue(Hostname + Platform + CPU + RAM + Disk)
}

func HashValue(value string) string {
	hash := sha256.New()
	hash.Write([]byte(value))
	return hex.EncodeToString(hash.Sum(nil))
}

func DecodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

func GetPublicIP() (string, error) {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query, nil
}

// Get preferred outbound ip of this machine
func GeLocalIp() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
