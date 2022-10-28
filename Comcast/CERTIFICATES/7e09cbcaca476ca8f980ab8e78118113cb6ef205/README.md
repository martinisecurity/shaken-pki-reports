# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate 7e09cbcaca476ca8f980ab8e78118113cb6ef205
Tested At: 2022-10-28 18:15:27 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 2 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/b51976bd-cc19-4104-8fcb-e0b8f3650680.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIIFAjSrCLE0pQwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMDkzMDEzMjk0N1oXDTIyMTAzMDEzMjk0N1owXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQnepuDIASxLpUNuysVfn6pSyG5WJfegT3fFUH6wuxFDIEz9fQ7A3g1jJhfmNB7QN6F1qLKdUHUT4zZw4cZL076o4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNIADBFAiAdlq6eBjqS6RfLhAo1mJIteup7f5rh5pgAuFKmGcRc%2BAIhAMvaBCdJtgau4p8b%2FYY9mSihMGV2iMQcgZEVwj43cJvC)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 7610' |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_sti_subject_key_identifier | error | ATIS-1000080 | STI certificates shall contain a Subject Key Identifier extension |

Generated: 28/10/2022 at 18:15:47