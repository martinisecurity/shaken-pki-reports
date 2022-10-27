# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate 8e62e5ab2267e22f4d55a4ea43494af79ce80900
Tested At: 2022-10-27 18:24:46 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 23 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/03ce07b4-5ab7-476a-a360-e0fdc8bf7ef0.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICWDCCAf2gAwIBAgIIJYTEfyzfW4owCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAyMDEzMjk1NloXDTIyMTExOTEzMjk1NlowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASGPZh1PTV0%2BZ6F7CQSSO35I%2FxOT%2Fv4Nr1Bnh5dsSDHi5VQgjComW6sMvyINTIWQh85RuNS%2FJpdZBvYsHSR0NNeo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNJADBGAiEAw75zNno0t9wAeiFntI%2BHWd%2F%2FvG6WZn1qm7XRSOJx0C0CIQCOPAkaBDFBwMuKfR1SVl2AY1jx%2B%2FN7iCsaWJMmHnR%2F%2FQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 7610' |
| e_sti_subject_key_identifier | error | ATIS-1000080v4 | STI certificates shall contain a Subject Key Identifier extension |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |

Generated: 27/10/2022 at 18:24:52