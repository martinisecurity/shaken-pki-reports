# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate b11ce2599d4b279750f8823c8f39c6e4540afafe
Tested At: 2022-10-27 18:56:44 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 22 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/1a35ef12-21c3-42a8-bb04-7a995edcb7a6.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICWDCCAf2gAwIBAgIIbmYeknWYBfUwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAxOTEzMjk0N1oXDTIyMTExODEzMjk0N1owXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATEyQQAZAkpP9l7s66gF1p7wApy%2FlX%2BlZOCvv%2FmpovhE2fEZpkdSZ6A96AjlLPKG3T%2F2qXLlD6kaFoA%2F81m3yEbo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNJADBGAiEA7TDUV0p6qXHaOtOvXh80tlJR0hsJyeX8bKjnFwCspXwCIQD1DEC5Q0SdRARmG8TepLL2%2FdgK0h5BYQylkUkQ2y%2FLJA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 7610' |
| e_sti_subject_key_identifier | error | ATIS-1000080v4 | STI certificates shall contain a Subject Key Identifier extension |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |

Generated: 27/10/2022 at 18:57:26