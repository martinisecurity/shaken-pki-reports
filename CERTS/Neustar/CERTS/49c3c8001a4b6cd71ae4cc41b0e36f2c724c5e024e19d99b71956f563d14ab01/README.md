# STIR/SHAKEN CA Ecosystem Compliance

## Certificate intelepeer.com

Tested At: 05 Jan 23 20:56 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 485 day(s)\
Subject: C=US, ST=CO, L=Centennial, O=IntelePeer CC LLC, OU=IntelePeer CC LLC, CN=intelepeer.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11238.10188

[View certificate details](https://understandingwebpki.com/?cert=MIIEDDCCAvSgAwIBAgIUVoi1GKC9sJ41Jlc%2BIyAmSLOE1q8wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MDQxMTQ2MThaFw0yNDA1MDQxMTQ2MThaMIGAMRcwFQYDVQQDDA5pbnRlbGVwZWVyLmNvbTEaMBgGA1UECwwRSW50ZWxlUGVlciBDQyBMTEMxGjAYBgNVBAoMEUludGVsZVBlZXIgQ0MgTExDMRMwEQYDVQQHDApDZW50ZW5uaWFsMQswCQYDVQQIDAJDTzELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ8XdF0h2EtK3SUp2cCblrqI3bUE%2B78sPpAXv8TCvf1U8QxdQhkK46jvPeEgBGTT1laZvYkbuV6DZq876hGm6nHo4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBROF1abaQbMaRRTiKvHRr7MH88pODAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDE3OEgwDQYJKoZIhvcNAQELBQADggEBAFALP74FjE7PIvgGoXnLmhCeg4omo3nRDpdJOJqkYCUBe04BHnRmUYJyltDvq%2B6myjlmhf4p4NZxPuSKqv2Yx4EtIYIhxhZuvcsgRcdAWvTLQB1wUhAHPEUhn1D7e4NqyURY4MIWKeKlYlFnoqHZU5VqRhZxfNftVeHIgmKMvCjEE5A4e28I0%2Bx%2Bkk4O4Czpij4GOiXvD85Mc4IvBHGM7WNY%2BLLJkukg7mpSSQZwqCxI5OIWj3M2D062yd3kCovYyX92QEuYHnue%2FwI%2Bewn3rEjOx77U%2FhMH2%2FqrrSHH8mVhwgM%2Bwhgbi3N%2Br4j94%2Bhu9ncLXbr72s7mZClFEwA%2FXMU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 05 Jan 23 21:05 UTC