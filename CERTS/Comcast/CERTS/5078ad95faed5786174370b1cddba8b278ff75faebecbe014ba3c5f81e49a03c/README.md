# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 31 Oct 22 16:42 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 13 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/39f48e26-2e26-42a8-8816-150688ced7c4.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVjCCAf2gAwIBAgIIA1yoDXJSGpAwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAxNDEwMDU0OVoXDTIyMTExMzEwMDU0OVowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ5EaHX83jowCpamvqSyfACPrgNUgPAsXS57G2UuFjL1xurSoKXIIAxs3fwffOwPTkU5f98DmQtpIYKgc0B7mZxo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNHADBEAiBi6NWoilr7OZ0maQ9IfVzF7jxsj4kZ0FcW9bJ2eOdAJQIgGzI6v2oJogEu7prPdX66aDnTV2tGPoptNyIYiXYVEVc%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_subject_key_identifier](../../ISSUES/e_sti_subject_key_identifier/README.md) | error | ATIS&#x2011;1000080 | STI certificates shall contain a Subject Key Identifier extension |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 7610' |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |


Generated: 31/10/2022 at 16:43:22