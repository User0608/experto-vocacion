package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Cors struct {
	AllowOrigins []string `json:"allow_origins"`
	AllowMethods []string `json:"allow_methods"`
}
type Certificates struct {
	Public  string `json:"public"`
	Private string `json:"private"`
}

type ServiceConfig struct {
	Address      string       `json:"address"`
	Cors         Cors         `json:"CORS"`
	Certificates Certificates `json:"certificates"`
	ExpertApiUrl string       `json:"expert_api_url"`
}
type DBConfigs struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	UserDB   string `json:"user_db"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

func (d *DBConfigs) GetConnectionString() string {
	cadena := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		d.Host,
		d.UserDB,
		d.Password,
		d.DBName,
		d.Port,
	)
	return cadena
}

func LoadServiceConfigs(path string) (*ServiceConfig, error) {
	configs := &ServiceConfig{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("No se pudo abrir el archivo de configuraciones: %s", err.Error())
	}
	if err := json.Unmarshal(file, configs); err != nil {
		return nil, fmt.Errorf("Archivo de configuraciones incorrecto, ERR: %s", err.Error())
	}
	return configs, nil
}
func LoadDBConfigs(path string) (*DBConfigs, error) {
	configs := &DBConfigs{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("No se pudo abrir el archivo de configuraciones db: %s", err.Error())
	}
	if err := json.Unmarshal(file, configs); err != nil {
		return nil, fmt.Errorf("Archivo de configuraciones incorrecto db, ERR: %s", err.Error())
	}
	return configs, nil
}
