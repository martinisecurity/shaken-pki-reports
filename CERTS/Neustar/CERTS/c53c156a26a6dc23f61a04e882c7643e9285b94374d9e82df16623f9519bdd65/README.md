# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Logix-Keystore

Tested At: 12 Apr 23 21:48 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 311 day(s)\
Subject: C=US, ST=TX, L=Dallas, O=Logix, OU=Logix Fiber Network, CN=Logix-Keystore\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11159.10178

[View certificate details](https://understandingwebpki.com/?cert=MIID%2FTCCAuWgAwIBAgIUFM3nikLOdll3rIAWIUJLa6tYEr0wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAyMTYxNTEyNDZaFw0yNDAyMTcxNTEyNDZaMHIxFzAVBgNVBAMMDkxvZ2l4LUtleXN0b3JlMRwwGgYDVQQLDBNMb2dpeCBGaWJlciBOZXR3b3JrMQ4wDAYDVQQKDAVMb2dpeDEPMA0GA1UEBwwGRGFsbGFzMQswCQYDVQQIDAJUWDELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ40Qm72bQdA49LPXKFZPJtOdgycKyhv8VowQFX7TaDiGJzI7opZ4cYSz6o77BJKQ9108RJrVlhWeTX9sHXg3HNo4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBQ%2F1niJ%2BDHAIUso3cD1kvuJUaoWqjAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDcwNzYwDQYJKoZIhvcNAQELBQADggEBADo269fCDPF8p2KnG%2F%2FZ5QsORygi60XCrR8d8IDNIDlJDpVuviBOPtUaLW1olVxw7ryRrjriOmVpyjplAvCrJp1n9j3KRYhJiCqOubH%2FQvMdWBrnzodYjCQjaV7hNm0FCJXg2WqdNmxS0wvOsYaReEnLMx9hroZe2AFSkwbR6So6LpHr6fOvTDJ%2FydSbclCwnhLnkjI2GpdYSpjAK1OUomc6W2LGpQ0dMF0AybolIaV9J2U2lbwkBg5xV07A7U26UGg3xM5yWDZxj9AkNxzGBFZXerBbNTGxO10GU665TF5zblGgW5HGCNARLReYIx9ExfRnZyUDWScDmgOJL%2BDsjLo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Apr 23 22:02 UTC