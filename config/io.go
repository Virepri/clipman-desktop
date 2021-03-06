package config

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

func SaveCFG() {
	config := cfg{
		AdminHash: AdminHash,
		UserHash:  UserHash,
		ServerIP:  ServerIP,
		Buffer:    Buffer,
	}

	_ = os.MkdirAll(CfgDir, 0600) //Who cares

	data, _ := json.Marshal(config)
	if f, err := os.OpenFile(CfgLocation, os.O_CREATE|os.O_RDWR, 0666); err == nil {
		if _, err := f.Write(data); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}

func LoadCFG() bool {
	if fi, err := os.Stat(CfgLocation); err == nil {
		buffer := make([]byte, fi.Size())
		if f, err := os.Open(CfgLocation); err == nil {
			if _, err := f.Read(buffer); err == nil {
				var config cfg
				if err := json.Unmarshal(buffer, &config); err == nil {
					AdminHash = config.AdminHash
					UserHash = config.UserHash
					Buffer = config.Buffer
					ServerIP = config.ServerIP
					TLSConfig.InsecureSkipVerify = config.TLSInsecure
					ServerUsesTLS = config.ServerUsesTLS
					if TLSConfig.InsecureSkipVerify && ServerUsesTLS {
						fmt.Println("WARNING: Using TLS without verification of certificate. This is unsafe and prone to MITM attacks.")
						fmt.Println("Add the server's certificate to your certpool if you trust it (or get the server to switch to a cert authority)")
					}
					if !ServerUsesTLS {
						fmt.Println("WARNING: Do not handle sensitive data without TLS enabled. This is unsafe and prone to MITM attacks.")
					}
					return true
				}
			}
		}
	}

	fmt.Println("Config doesn't seem to exist, generating a config.")

	ServerIP = "127.0.0.1:7606"
	Buffer = 1024

	hash := sha256.New()
	hash.Write([]byte("password"))
	AdminHash = hex.EncodeToString(hash.Sum(nil))
	UserHash = hex.EncodeToString(hash.Sum(nil))
	TLSConfig.InsecureSkipVerify = false
	ServerUsesTLS = false

	SaveCFG()
	return false
}
