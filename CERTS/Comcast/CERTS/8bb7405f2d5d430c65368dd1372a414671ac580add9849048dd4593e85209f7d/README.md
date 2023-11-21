# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 21 Nov 23 01:52 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -52 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/c56cbab6-41d8-4ca1-bdef-0b209d699a5e.cer

[View certificate details](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIIPASVGqQmVdkwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIzMDgzMDAyMzg1OFoXDTIzMDkyOTAyMzg1OFowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR3%2F7Hj%2F4ta4IYj7r1UFvmZUziywuc1AoaaXQcOtv022p1kxTUE56W8GafbzgCz61Eaw%2F0qPw31bwAngQSOVAOgo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYEMzE4SjAKBggqhkjOPQQDAgNIADBFAiEA4PCEw2Kd3sKZVpFNgFCnXieUv7QPxU3Yjhjpc4lXdPoCIFNJwtHUC100RSsS3nzKwKvBO3KRqTljKoyE%2FmJj6dfK)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 318J' |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_subject_key_identifier_size](../../ISSUES/e_atis_subject_key_identifier_size/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | the Certificate Policies extension is not present |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |


Generated: 21 Nov 23 01:55 UTC