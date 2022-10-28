# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate caa3abae9e9defc436e4c1208744ec71587c5e1b
Tested At: 2022-10-28 18:15:38 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 19 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/00aec434-5013-4136-9d51-04f563275d0d.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIILjEI2rsVhkQwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAxNzEzMjk0N1oXDTIyMTExNjEzMjk0N1owXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARgJGbTpmi5z9vtFRfxQvgwGdn86SKS0pThiMWEoNd%2B%2FUxnvwpGPNBgKvnKpk10gH9MooL%2FFM14fxxMr0G6Rj2io4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNIADBFAiB3pKj%2F2C0kL744JsmFKhOmVMSeT6O6VBtQZqXO3LyEiAIhANPHjSEOFw55A7phHpNUTt6vlZttqUOyKm87a5h29lkU)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 7610' |
| e_sti_subject_key_identifier | error | ATIS-1000080 | STI certificates shall contain a Subject Key Identifier extension |

Generated: 28/10/2022 at 18:15:47