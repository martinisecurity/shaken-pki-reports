# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate 2e37e6bc959af5b790ba4727baa232d8a81f6559
Tested At: 2022-10-27 22:31:52 +0000 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 478 day(s)\
Subject: C=US, ST=TX, L=Dallas, O=Logix, OU=Logix Fiber Network, CN=Logix-Keystore\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1

Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11159.10178

View: [Click to view](https://understandingwebpki.com/?cert=MIID%2FTCCAuWgAwIBAgIUFM3nikLOdll3rIAWIUJLa6tYEr0wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAyMTYxNTEyNDZaFw0yNDAyMTcxNTEyNDZaMHIxFzAVBgNVBAMMDkxvZ2l4LUtleXN0b3JlMRwwGgYDVQQLDBNMb2dpeCBGaWJlciBOZXR3b3JrMQ4wDAYDVQQKDAVMb2dpeDEPMA0GA1UEBwwGRGFsbGFzMQswCQYDVQQIDAJUWDELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ40Qm72bQdA49LPXKFZPJtOdgycKyhv8VowQFX7TaDiGJzI7opZ4cYSz6o77BJKQ9108RJrVlhWeTX9sHXg3HNo4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBQ%2F1niJ%2BDHAIUso3cD1kvuJUaoWqjAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDcwNzYwDQYJKoZIhvcNAQELBQADggEBADo269fCDPF8p2KnG%2F%2FZ5QsORygi60XCrR8d8IDNIDlJDpVuviBOPtUaLW1olVxw7ryRrjriOmVpyjplAvCrJp1n9j3KRYhJiCqOubH%2FQvMdWBrnzodYjCQjaV7hNm0FCJXg2WqdNmxS0wvOsYaReEnLMx9hroZe2AFSkwbR6So6LpHr6fOvTDJ%2FydSbclCwnhLnkjI2GpdYSpjAK1OUomc6W2LGpQ0dMF0AybolIaV9J2U2lbwkBg5xV07A7U26UGg3xM5yWDZxj9AkNxzGBFZXerBbNTGxO10GU665TF5zblGgW5HGCNARLReYIx9ExfRnZyUDWScDmgOJL%2BDsjLo%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_sti_crl_distribution_not_reachable | error | ATIS-1000080 | Unable to retrieve CRL specified in CRLdp from allow listed IP address |

### Not Effective

- e_sti_subject_cn
- e_sti_signature_algorithm
- e_sti_serial_number
- e_cp1_3_ambiguous_identifier
- e_sti_extension_unknown
- w_cp_1_3_subject_email
- w_cp1_3_subject_rdn_unknown
- e_cp1_3_subject_sn

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 22:33:03