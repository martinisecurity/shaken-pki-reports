# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate dc9e764717b56faa963a0b026603f04b33542673
Tested At: 2022-10-26 20:32:07 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 9 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/045bff20-7369-413e-a5d1-fd103a7f0648.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIIG2z%2ForpReTYwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAwNTEzMjk1MFoXDTIyMTEwNDEzMjk1MFowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASjao6yclzWguhvEfx0PfdUgSlEpjipCF5KsrfGF%2FpcSoqC4IOt4HNIpGIvZDDBdvRtrioK%2B5nY319L5FQMMo%2B5o4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNIADBFAiEA4aDN%2FFwVSXWxMTpjSonCdcinFtezoCAdAujdPB4Qu1ICICADR5OypPD8YlGfEO1Qbg2iism6BrVEtrIOZ8n59jpo)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_key_identifier | error | ATIS-1000080v4 | STI certificates shall contain a Subject Key Identifier extension |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 7610' |

Generated: 26/10/2022 at 20:32:17