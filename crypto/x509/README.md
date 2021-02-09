# Generate CA

## Generate Root CA

Create root private key.

```bash
openssl genrsa -out ca.key 2048
```

Open openssl config file `openssl.conf`. Note set `basicConstraints` to true, that you can modify the values.

```text
[ req ]
#default_bits		= 2048
#default_md		= sha256
#default_keyfile 	= privkey.pem
distinguished_name	= req_distinguished_name
attributes		= req_attributes
extensions               = v3_ca
req_extensions           = v3_ca

[ req_distinguished_name ]
countryName			= Country Name (2 letter code)
countryName_min			= 2
countryName_max			= 2
stateOrProvinceName		= State or Province Name (full name)
localityName			= Locality Name (eg, city)
0.organizationName		= Organization Name (eg, company)
organizationalUnitName		= Organizational Unit Name (eg, section)
commonName			= Common Name (eg, fully qualified host name)
commonName_max			= 64
emailAddress			= Email Address
emailAddress_max		= 64

[ req_attributes ]
challengePassword		= A challenge password
challengePassword_min		= 4
challengePassword_max		= 20

[ v3_ca ]
basicConstraints         = CA:TRUE
```

Then generate CA's certificate using the config file `openssl.conf`.

```bash
openssl req -new -key ca.key -nodes -out ca.csr -config openssl.conf
openssl x509 -req -days 36500 -extfile openssl.conf -extensions v3_ca -in ca.csr -signkey ca.key -out rootCA.crt
```

## Generate Server CA

Create server private key.

```bash
openssl genrsa -out sca.key 2048
```

Create cert request file

```bash
openssl req -new -key sca.key -out sca.csr
```

Create server cert.

```bash
openssl x509 -req -in sca.csr -CA ca.crt -CAkey ca.key -out sca.crt -days 36500
```
