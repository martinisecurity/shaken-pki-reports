# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN-6744

Tested At: 31 Oct 22 20:40 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 497 day(s)\
Subject: C=US, ST=GA, L=Atlanta, O=Southern Linc, OU=Voice and Signaling, CN=SHAKEN-6744\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://prod001-cr.rbbnidhub.com/frvFJbsMgz/sign-cert3

View: [Click to view](https://understandingwebpki.com/?cert=MIIEAzCCAuugAwIBAgIUYVVb6nzj7tDwHJG%2FOFgAaDCCUTgwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAzMTExODE4NDRaFw0yNDAzMTExODE4NDRaMHgxFDASBgNVBAMMC1NIQUtFTi02NzQ0MRwwGgYDVQQLDBNWb2ljZSBhbmQgU2lnbmFsaW5nMRYwFAYDVQQKDA1Tb3V0aGVybiBMaW5jMRAwDgYDVQQHDAdBdGxhbnRhMQswCQYDVQQIDAJHQTELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARqrLi6cvEjAQY8kdFvhSOajGuf2HaLQKHM%2Fo%2FX5681zdbmnsyaCdHQQN%2FDXSAncgbzmq8h8IloDmbf7X2EHeHTo4IBSDCCAUQwFgYIKwYBBQUHARoECjAIoAYWBDY3NDQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBR4Yx2HeDngkDecHQIsGdmh7zVttzAOBgNVHQ8BAf8EBAMCB4AwDQYJKoZIhvcNAQELBQADggEBAFMUekf2AqW9Ug1T7Tz%2F5gnr%2FWTrOUezYfxXgJIZn2zDdAnLHy9sOC%2FbhT40bMsKcqPUul%2FkRiXcDO3zpZE6y8VMJyeO%2FAr9TUdNq0%2Fqfe7GJqJUOcIXNC8mEZpiDP3M2SncD%2BXsA2WWPrIFdi%2BFoauu9S4mZwp8KJjOsx56QAoWI1EBRfJF90EPnt%2FgQUaSYRBMDwNWGz5vTUIiCO4hnMgWkJK8w2VTT4N9sWLx64v2XM4d9loclBVmmBm6%2BSwc7GdDc4wruE6Vhm3JAK%2FSNifzIh6XSsX%2BfvUPpGRj43fGA%2Fr2V0qc86cKpjPlIlrbjJAHNrmiGjZ6ifyrPSolfiU%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


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