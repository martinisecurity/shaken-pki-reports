# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN-406H

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 438 day(s)\
Subject: C=US, ST=CA, L=Tarzana, O=Cybernet Communications\\, Inc., OU=Cybernet1, CN=SHAKEN-406H\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/37.202

[View certificate details](https://understandingwebpki.com/?cert=MIIECTCCAvGgAwIBAgIUKin04RKV%2BB7D2428ad0IvzDvuNswDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA0MTMxMzEyNDFaFw0yNDA0MTMxMzEyNDFaMH4xFDASBgNVBAMMC1NIQUtFTi00MDZIMRIwEAYDVQQLDAlDeWJlcm5ldDExJjAkBgNVBAoMHUN5YmVybmV0IENvbW11bmljYXRpb25zLCBJbmMuMRAwDgYDVQQHDAdUYXJ6YW5hMQswCQYDVQQIDAJDQTELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ0QHHEO9QUqZroFH4GUaXF9sjsJXM%2FCf6e98bDgK7iuTuZHYKaHiu6pSujtjGJ4l1aIOejgX%2BFGeluy%2BgcuTRXo4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBRPS2T5%2FVbX6yRR7TAoXUDetQLNJDAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDQwNkgwDQYJKoZIhvcNAQELBQADggEBAHbQHtlcB4OQ1p9RtKT7ABtJxQHTs1HDEAVHPYJo2SnYI5AS4SdRP3dHpO4TEsw9Qre2u5FKhu%2Bl6RL71D63hzRS9ujQ4RHnloMxqgfw15HZ5TFoh4nQlxuB%2FUxtairhk4u7ux1nxC2ZZDKVRUv%2Bh6FFLr4E1uoBp9lsnSnYkDEENLiRm%2FJ0j0NwZN9ag3Gsq%2BFC9K6mJTXSUui3ATo%2BlBnqL91Whh2jDeHTL1Lj8bCDIE5xOuWchTyrUo80oQxJzodY5ovtKoh2HCMTLCN6gEXokTi%2BMEcolMXPPYXfc2iwiU2VWtuF%2BnrLImM3Z2goFXBYKr4a8jjP4%2FeYYcIh%2Bx4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31 Jan 23 21:50 UTC