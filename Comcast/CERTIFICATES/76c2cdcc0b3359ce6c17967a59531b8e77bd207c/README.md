# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate 76c2cdcc0b3359ce6c17967a59531b8e77bd207c
Tested At: 2022-10-28 18:15:26 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 29 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/32780845-7305-4c38-b97b-e613e8cfb314.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIIFAjSrCLE0pQwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAyNzExMTIzOFoXDTIyMTEyNjExMTIzOFowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQJnb1B9FdNwikpbUbScMfidmatD9Uzmjx3ppRIkUbew8x5K0jEvpjw6cAD35XnJ%2FeEf1Yf256DXgWj1y%2FN4W%2Fjo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYEMzE4SjAKBggqhkjOPQQDAgNIADBFAiEArpCF73peC9kMHhL3KC%2Bu2%2B0Orz9S34fIeyLLLAvcbdwCIFtBhBms35N6StykBrit0IbZN3G%2B%2FSEdIaulOUiN1Vf3)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_key_identifier | error | ATIS-1000080 | STI certificates shall contain a Subject Key Identifier extension |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 318J' |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |

Generated: 28/10/2022 at 18:15:47