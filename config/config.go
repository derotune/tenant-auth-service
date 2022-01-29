package config

import (
	"regexp"

	"github.com/spf13/viper"
)

type Provider struct {
	Name   string `mapstructure:"Name"`
	Issuer string `mapstructure:"Issuer"`
}

type Tenant struct {
	Name         string   `mapstructure:"Name"`
	ClientId     string   `mapstructure:"ClientId"`
	ClientSecret string   `mapstructure:"ClientSecret"`
	Provider     Provider `mapstructure:"Provider"`
}

type Host struct {
	Regex   string   `mapstructure:"Regex"`
	Tenants []Tenant `mapstructure:"Tenants"`
}

type Config struct {
	Hosts []Host `mapstructure:"Hosts"`
}

var LoadedConfiguration Config

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&LoadedConfiguration)
}

func FindMatchingConfiguration(url string) (tenant *Tenant) {

	for _, hostConfiguration := range LoadedConfiguration.Hosts {
		var re = regexp.MustCompile(hostConfiguration.Regex)
		matches := re.FindStringSubmatch(url)

		if len(matches) == 0 {
			return nil
		}

		// we found matching configuration so lets extract the client config
		matchedClient := matches[re.SubexpIndex("Client")]

		for _, tenantConfig := range hostConfiguration.Tenants {
			if tenantConfig.Name == matchedClient {
				return &tenantConfig
			}
		}
	}

	return nil
}
