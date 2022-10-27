# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate 7adb08c0422dfdbc535f11c8f108e6110df249ea
Tested At: 2022-10-27 18:56:20 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 24 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/790cde70-2d5d-4e3d-88dc-84283ad70f47.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIIJgbNK1fSkkUwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAyMTEzMjk0N1oXDTIyMTEyMDEzMjk0N1owXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATjw3L0t9M0sYlFAZCQshIF1GQnbBFg9stWJ4y9rDzIzQiP7Zp13aw7QynPuNyYKb2EKPHujAAV1rlhgGyZNYnJo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNIADBFAiEAhbVycrmnaB%2Fkoak6%2Bz1ajXihM3ln%2FiYcZkifI7J5Y5oCIHfusuDvaDI2LSlnDkyCb%2BdJz5SCgFNjFn2W%2BMxpiu%2Bh)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 7610' |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_key_identifier | error | ATIS-1000080v4 | STI certificates shall contain a Subject Key Identifier extension |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

Generated: 27/10/2022 at 18:57:26