# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Logix-Keystore

Tested At: 28 Nov 23 19:57 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 81 day(s)\
Subject: C=US, ST=TX, L=Dallas, O=Logix, OU=Logix Fiber Network, CN=Logix-Keystore\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11159.10178

[View certificate details](https://understandingwebpki.com/?cert=MIID%2FTCCAuWgAwIBAgIUFM3nikLOdll3rIAWIUJLa6tYEr0wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAyMTYxNTEyNDZaFw0yNDAyMTcxNTEyNDZaMHIxFzAVBgNVBAMMDkxvZ2l4LUtleXN0b3JlMRwwGgYDVQQLDBNMb2dpeCBGaWJlciBOZXR3b3JrMQ4wDAYDVQQKDAVMb2dpeDEPMA0GA1UEBwwGRGFsbGFzMQswCQYDVQQIDAJUWDELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ40Qm72bQdA49LPXKFZPJtOdgycKyhv8VowQFX7TaDiGJzI7opZ4cYSz6o77BJKQ9108RJrVlhWeTX9sHXg3HNo4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBQ%2F1niJ%2BDHAIUso3cD1kvuJUaoWqjAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDcwNzYwDQYJKoZIhvcNAQELBQADggEBADo269fCDPF8p2KnG%2F%2FZ5QsORygi60XCrR8d8IDNIDlJDpVuviBOPtUaLW1olVxw7ryRrjriOmVpyjplAvCrJp1n9j3KRYhJiCqOubH%2FQvMdWBrnzodYjCQjaV7hNm0FCJXg2WqdNmxS0wvOsYaReEnLMx9hroZe2AFSkwbR6So6LpHr6fOvTDJ%2FydSbclCwnhLnkjI2GpdYSpjAK1OUomc6W2LGpQ0dMF0AybolIaV9J2U2lbwkBg5xV07A7U26UGg3xM5yWDZxj9AkNxzGBFZXerBbNTGxO10GU665TF5zblGgW5HGCNARLReYIx9ExfRnZyUDWScDmgOJL%2BDsjLo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |
| [e_atis_signature_algorithm](../../ISSUES/e_atis_signature_algorithm/README.md) | error | ATIS1000080 | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.113549.1.1.11 |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'Logix-Keystore' does not contain 'SHAKEN' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 20:21 UTC