# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate 2a8742b8e2c37f409a0c9c229b8b901248442e37
Tested At: 2022-10-27 21:25:30 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 17 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/39f48e26-2e26-42a8-8816-150688ced7c4.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVjCCAf2gAwIBAgIIA1yoDXJSGpAwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAxNDEwMDU0OVoXDTIyMTExMzEwMDU0OVowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ5EaHX83jowCpamvqSyfACPrgNUgPAsXS57G2UuFjL1xurSoKXIIAxs3fwffOwPTkU5f98DmQtpIYKgc0B7mZxo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNHADBEAiBi6NWoilr7OZ0maQ9IfVzF7jxsj4kZ0FcW9bJ2eOdAJQIgGzI6v2oJogEu7prPdX66aDnTV2tGPoptNyIYiXYVEVc%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_sti_subject_key_identifier | error | ATIS-1000080 | STI certificates shall contain a Subject Key Identifier extension |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 7610' |

Generated: 27/10/2022 at 21:27:34