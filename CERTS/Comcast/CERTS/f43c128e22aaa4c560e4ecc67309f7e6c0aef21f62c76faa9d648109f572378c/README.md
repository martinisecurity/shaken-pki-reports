# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 01 Mar 23 18:22 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -105 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/00aec434-5013-4136-9d51-04f563275d0d.cer

[View certificate details](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIILjEI2rsVhkQwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAxNzEzMjk0N1oXDTIyMTExNjEzMjk0N1owXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARgJGbTpmi5z9vtFRfxQvgwGdn86SKS0pThiMWEoNd%2B%2FUxnvwpGPNBgKvnKpk10gH9MooL%2FFM14fxxMr0G6Rj2io4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNIADBFAiB3pKj%2F2C0kL744JsmFKhOmVMSeT6O6VBtQZqXO3LyEiAIhANPHjSEOFw55A7phHpNUTt6vlZttqUOyKm87a5h29lkU)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7610' |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain a Subject Key Identifier extension |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 01 Mar 23 18:22 UTC