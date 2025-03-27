package utils

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	globalSnowflakeIns *Snowflake
)

func init() {
	var err error
	globalSnowflakeIns, err = NewSnowflakeWithIP()
	if err != nil {
		panic(err)
	}
}

const (
	epoch          int64 = 1735689600000 // 2025-01-01 00:00:00
	machineIDBits  uint8 = 10
	sequenceBits   uint8 = 12
	machineIDShift uint8 = sequenceBits
	timestampShift uint8 = machineIDBits + sequenceBits
	sequenceMask   int64 = -1 ^ (-1 << sequenceBits)
)

var (
	ErrInvalidMachineID = errors.New("machine ID out of range")
)

type Snowflake struct {
	mu        sync.Mutex
	lastTime  int64
	machineID int64
	sequence  int64
}

// NewSnowflake 初始化 Snowflake 实例
func NewSnowflake(machineID int64) (*Snowflake, error) {
	if machineID < 0 || machineID >= (1<<machineIDBits) {
		return nil, ErrInvalidMachineID
	}
	return &Snowflake{
		lastTime:  0,
		machineID: machineID,
		sequence:  0,
	}, nil
}

// Generate 生成唯一 ID
func (s *Snowflake) Generate() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli()
	if now < s.lastTime {
		panic("clock moved backwards")
	}

	if now == s.lastTime {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.lastTime {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTime = now
	return uint64(((now - epoch) << timestampShift) | (s.machineID << machineIDShift) | s.sequence)
}

// GetContainerIP 获取容器 IP 地址
func GetContainerIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("no IP address found")
}

// IPToMachineID 将 IP 地址转换为机器 ID
func IPToMachineID(ip string) int64 {
	var machineID int64
	_, err := fmt.Sscanf(ip, "%*d.%*d.%*d.%d", &machineID)
	if err != nil {
		return 0
	}
	return machineID
}

// NewSnowflakeWithIP 自动获取容器 IP 并初始化 Snowflake
func NewSnowflakeWithIP() (*Snowflake, error) {
	ip, err := GetContainerIP()
	if err != nil {
		return nil, err
	}
	machineID := IPToMachineID(ip)
	return NewSnowflake(machineID)
}

func GenSnowflakeId() uint64 {
	return globalSnowflakeIns.Generate()
}
