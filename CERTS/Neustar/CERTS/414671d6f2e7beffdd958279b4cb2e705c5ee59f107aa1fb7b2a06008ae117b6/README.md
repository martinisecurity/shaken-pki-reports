# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN-6744

Tested At: 02 Nov 22 17:24 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 496 day(s)\
Subject: C=US, ST=GA, L=Atlanta, O=Southern Linc, OU=Voice and Signaling, CN=SHAKEN-6744\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://prod001-cr.rbbnidhub.com/frvFJbsMgz/sign-cert3

[View certificate details](https://understandingwebpki.com/?cert=MIIEAzCCAuugAwIBAgIUYVVb6nzj7tDwHJG%2FOFgAaDCCUTgwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAzMTExODE4NDRaFw0yNDAzMTExODE4NDRaMHgxFDASBgNVBAMMC1NIQUtFTi02NzQ0MRwwGgYDVQQLDBNWb2ljZSBhbmQgU2lnbmFsaW5nMRYwFAYDVQQKDA1Tb3V0aGVybiBMaW5jMRAwDgYDVQQHDAdBdGxhbnRhMQswCQYDVQQIDAJHQTELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARqrLi6cvEjAQY8kdFvhSOajGuf2HaLQKHM%2Fo%2FX5681zdbmnsyaCdHQQN%2FDXSAncgbzmq8h8IloDmbf7X2EHeHTo4IBSDCCAUQwFgYIKwYBBQUHARoECjAIoAYWBDY3NDQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBR4Yx2HeDngkDecHQIsGdmh7zVttzAOBgNVHQ8BAf8EBAMCB4AwDQYJKoZIhvcNAQELBQADggEBAFMUekf2AqW9Ug1T7Tz%2F5gnr%2FWTrOUezYfxXgJIZn2zDdAnLHy9sOC%2FbhT40bMsKcqPUul%2FkRiXcDO3zpZE6y8VMJyeO%2FAr9TUdNq0%2Fqfe7GJqJUOcIXNC8mEZpiDP3M2SncD%2BXsA2WWPrIFdi%2BFoauu9S4mZwp8KJjOsx56QAoWI1EBRfJF90EPnt%2FgQUaSYRBMDwNWGz5vTUIiCO4hnMgWkJK8w2VTT4N9sWLx64v2XM4d9loclBVmmBm6%2BSwc7GdDc4wruE6Vhm3JAK%2FSNifzIh6XSsX%2BfvUPpGRj43fGA%2Fr2V0qc86cKpjPlIlrbjJAHNrmiGjZ6ifyrPSolfiU%3D)

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


Generated: 02 Nov 22 17:25 UTC