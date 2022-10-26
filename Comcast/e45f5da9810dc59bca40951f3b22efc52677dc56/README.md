# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate e45f5da9810dc59bca40951f3b22efc52677dc56
Tested At: 2022-10-26 20:32:10 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 7 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/1827d9ff-facd-4d30-9ccb-16cd9bc20fc6.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICWDCCAf2gAwIBAgIIQfJ8xvOHXQQwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAwMzEyNDYyMVoXDTIyMTEwMjEyNDYyMVowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQvv%2FrENfYftD4tfsoBXT2n8o2ZgBlWCM%2BlBQWQCVzG1wrJL3Ou3380brt2mvA6eDrUT6M94KYZcd0xQyOlbvWLo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNJADBGAiEAqqjAvdp51kRyjvizGqwUJ9o%2BcKLuN4iRoxY%2BQx9cwCYCIQDNacrS9n1zVcQzekvS7K9OLSR2DFwAkfb2pmIQunq5Kw%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_subject_key_identifier | error | ATIS-1000080v4 | STI certificates shall contain a Subject Key Identifier extension |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 7610' |

Generated: 26/10/2022 at 20:32:17