package config

type Config struct {
  UseHTTP           bool
  UseTLS            bool
  HttpAddr          string
  HttpsAddr         string
  SkipVerify        string
  CommonName        string
  ClientAuth        string
  CACommonName      string
  CAOrganization    string
  CAOrganizationUnit    string
  CACountry         string
  CAProvince        string
  CALocality        string
  CAStreetAddress   string
  CAPostalCode      string
  CAYearsValid      string
  CertOrganization  string
  CertOrganizationUnit  string
  CertCountry       string
  CertProvince      string
  CertLocality      string
  CertStreetAddress string
  CertPostalCode    string
  CertIPAddresses   string
  CertDomains       string
  CertYearsValid    string
  ClientCertCommonName    string
  ClientCertOrganization  string
  ClientCertOrganizationUnit  string
  ClientCertCountry       string
  ClientCertProvince      string
  ClientCertLocality      string
  ClientCertStreetAddress string
  ClientCertPostalCode    string
  ClientCertIPAddresses   string
  ClientCertDomains       string
  ClientCertYearsValid    string
  RSAKeySize        string
  MinTLS            string
  MaxTLS            string
}

func NewConfig() *Config {
  return &Config{
    UseHTTP:          false,
    UseTLS:           false,
    HttpAddr:         ":8080",
    HttpsAddr:        ":8443",
    SkipVerify:       "false",
    CommonName:       "server.localhost",
    ClientAuth:       "none",
    CACommonName:     "root.localhost",
    CAOrganization:   "",
    CAOrganizationUnit: "",
    CACountry:        "US",
    CAProvince:       "",
    CALocality:       "",
    CAStreetAddress:  "",
    CAPostalCode:     "",
    CAYearsValid:     "10",
    CertOrganization: "",
    CertOrganizationUnit: "",
    CertCountry:      "US",
    CertProvince:     "",
    CertLocality:     "",
    CertStreetAddress:  "",
    CertPostalCode:   "",
    CertIPAddresses:  "",
    CertDomains:      "",
    CertYearsValid:   "1",
    ClientCertCommonName: "client.localhost",
    ClientCertOrganization: "",
    ClientCertOrganizationUnit: "",
    ClientCertCountry:      "US",
    ClientCertProvince:     "",
    ClientCertLocality:     "",
    ClientCertStreetAddress:  "",
    ClientCertPostalCode:   "",
    ClientCertIPAddresses:  "",
    ClientCertDomains:      "",
    ClientCertYearsValid:   "1",
    RSAKeySize:       "4096",
    MinTLS:           "1.0",
    MaxTLS:           "1.3",
  }
}

func (c *Config) CAOptions() map[string]string {
  return map[string]string{
    "cn":  c.CACommonName,
    "org": c.CAOrganization,
    "ou": c.CAOrganizationUnit,
    "country": c.CACountry,
    "province": c.CAProvince,
    "locality": c.CALocality,
    "street": c.CAStreetAddress,
    "postal": c.CAPostalCode,
    "isCA": "true",
    "client": "true",
    "server": "true",
    "years": c.CAYearsValid,
  }
}

func (c *Config) CertOptions() map[string]string {
  return map[string]string{
    "cn":  c.CommonName,
    "org": c.CertOrganization,
    "ou": c.CertOrganizationUnit,
    "country": c.CertCountry,
    "province": c.CertProvince,
    "locality": c.CertLocality,
    "street": c.CertStreetAddress,
    "postal": c.CertPostalCode,
    "isCA": "true",
    "client": "true",
    "server": "true",
    "ips": c.CertIPAddresses,
    "domains":  c.CertDomains,
    "years": c.CertYearsValid,
  }
}
func (c *Config) ClientCertOptions() map[string]string {
  return map[string]string{
    "cn":  c.ClientCertCommonName,
    "org": c.ClientCertOrganization,
    "ou": c.ClientCertOrganizationUnit,
    "country": c.ClientCertCountry,
    "province": c.ClientCertProvince,
    "locality": c.ClientCertLocality,
    "street": c.ClientCertStreetAddress,
    "postal": c.ClientCertPostalCode,
    "isCA": "true",
    "client": "true",
    "server": "true",
    "ips": c.ClientCertIPAddresses,
    "domains":  c.ClientCertDomains,
    "years": c.ClientCertYearsValid,
  }
}

func (c *Config) KeyOptions() map[string]string {
  return map[string]string{
    "keysize": c.RSAKeySize,
  }
}
