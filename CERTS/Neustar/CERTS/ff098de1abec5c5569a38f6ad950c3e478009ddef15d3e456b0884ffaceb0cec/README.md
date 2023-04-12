# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Flowroute

Tested At: 12 Apr 23 01:37 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 336 day(s)\
Subject: C=US, ST=WA, L=Seattle, O=Intrado, OU=Cloud Collaboration, CN=Flowroute\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11338.10182

[View certificate details](https://understandingwebpki.com/?cert=MIID%2BzCCAuOgAwIBAgIUZLK3HL9Y8lSVdU4F3nszxTfs5T4wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAzMTIxNDU4MDZaFw0yNDAzMTIxNDU4MDZaMHAxEjAQBgNVBAMMCUZsb3dyb3V0ZTEcMBoGA1UECwwTQ2xvdWQgQ29sbGFib3JhdGlvbjEQMA4GA1UECgwHSW50cmFkbzEQMA4GA1UEBwwHU2VhdHRsZTELMAkGA1UECAwCV0ExCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE60M79nV6619OIca4mGxxjaVovxR%2FC5yBNfLL71qQR%2B7kurTRIJa4nw7XhEtZPbh9yrUa0mGz%2FLE3eOaZ8ORgd6OCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUdif8H5sGexOSdVwIB%2F8hJVTduuswDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ2ODNHMA0GCSqGSIb3DQEBCwUAA4IBAQC1062x05nTOzMqDincO%2BZTaZYS6dCgADStKspUKldXg6uF8ris0dYQQMTuW6JcqQpsLKGVGMKq7mLZZPtQLqtYmoTwZYDC9exwkI0MuR2oew0v%2BNEKjZnwRyKNtNUWOr7Og8wLozZQPzpN1ucQuqU4eDEigajtYzdjLpXNQ5iWfprSbLyRoVnQVXlH%2F6btBn845M8u%2FQqAT6OE18aXeTjYYHm5t5LkB52TFSCOP%2FAzSJSYZ0Wd3mVDDFqzlT0eJQ1Hdd1qlNrOQVFaPlhrdinp1oKlxIYfNT5spyNS4YZ%2FJynF%2FnQW5Flt%2FwOUqiz9tsaGzyd1Gzd6tAtz4iGjxo91)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Apr 23 01:46 UTC