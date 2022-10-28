# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate f51dbe5b93ee71bc386bb2c3e0c6fb2052ca579b
Tested At: 2022-10-28 18:15:45 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 16 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/fd36683c-6292-4ab0-879d-e59fe379bfbf.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICWDCCAf2gAwIBAgIILjEI2rsVhkQwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAxNDEyNDYyOVoXDTIyMTExMzEyNDYyOVowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASoNAC%2FCwkQqnsCVmAngVOJhVDZ4XrVIkeLKCrlhWLoN%2Fp0ag2ZaMPP5k6521DkiXEJyQsarwoPeUAnmAYOEBJto4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNJADBGAiEAxcdgfkDo0TYH1cT7mbhDOkuWjFu3bUfBuH600vjKJJkCIQC%2FInwr6E784lHvGBGJAeqne%2F9Vbs%2B3GnC79G2xv9Gjqw%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_subject_key_identifier | error | ATIS-1000080 | STI certificates shall contain a Subject Key Identifier extension |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 7610' |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 18:15:47