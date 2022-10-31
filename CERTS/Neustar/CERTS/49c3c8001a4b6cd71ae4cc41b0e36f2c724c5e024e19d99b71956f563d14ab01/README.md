# STIR/SHAKEN CA Ecosystem Compliance

## Certificate intelepeer.com

Tested At: 31 Oct 22 20:40 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 551 day(s)\
Subject: C=US, ST=CO, L=Centennial, O=IntelePeer CC LLC, OU=IntelePeer CC LLC, CN=intelepeer.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11238.10188

View: [Click to view](https://understandingwebpki.com/?cert=MIIEDDCCAvSgAwIBAgIUVoi1GKC9sJ41Jlc%2BIyAmSLOE1q8wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MDQxMTQ2MThaFw0yNDA1MDQxMTQ2MThaMIGAMRcwFQYDVQQDDA5pbnRlbGVwZWVyLmNvbTEaMBgGA1UECwwRSW50ZWxlUGVlciBDQyBMTEMxGjAYBgNVBAoMEUludGVsZVBlZXIgQ0MgTExDMRMwEQYDVQQHDApDZW50ZW5uaWFsMQswCQYDVQQIDAJDTzELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ8XdF0h2EtK3SUp2cCblrqI3bUE%2B78sPpAXv8TCvf1U8QxdQhkK46jvPeEgBGTT1laZvYkbuV6DZq876hGm6nHo4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBROF1abaQbMaRRTiKvHRr7MH88pODAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDE3OEgwDQYJKoZIhvcNAQELBQADggEBAFALP74FjE7PIvgGoXnLmhCeg4omo3nRDpdJOJqkYCUBe04BHnRmUYJyltDvq%2B6myjlmhf4p4NZxPuSKqv2Yx4EtIYIhxhZuvcsgRcdAWvTLQB1wUhAHPEUhn1D7e4NqyURY4MIWKeKlYlFnoqHZU5VqRhZxfNftVeHIgmKMvCjEE5A4e28I0%2Bx%2Bkk4O4Czpij4GOiXvD85Mc4IvBHGM7WNY%2BLLJkukg7mpSSQZwqCxI5OIWj3M2D062yd3kCovYyX92QEuYHnue%2FwI%2Bewn3rEjOx77U%2FhMH2%2FqrrSHH8mVhwgM%2Bwhgbi3N%2Br4j94%2Bhu9ncLXbr72s7mZClFEwA%2FXMU%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 20:47:45