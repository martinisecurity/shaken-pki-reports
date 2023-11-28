# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 28 Nov 23 16:10 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -76 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/50c7a58d-a438-40ea-a340-92149a160d70.cer

[View certificate details](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIIctCh6UBq7J0wCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIzMDgxNDEzNTUwM1oXDTIzMDkxMzEzNTUwM1owXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATPacT9Gl75Mut3YTou8JTjMUubIkeIlH2JlHq33AWw%2F4ltpW%2B2OHhgocRfOfRh3iTNcb6zKUhJ0c31rWmKRUKbo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYEMzE4SjAKBggqhkjOPQQDAgNIADBFAiEA82bGgXXrX7D42c48xpNd7MoxbB0QkrUdy%2BD%2F6vwtTRkCIEKyDWhdyJ6SN%2BzbJOn2desUMR2Z%2F%2FSaHfWBQYjIRtXZ)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |
| [e_atis_subject_key_identifier_size](../../ISSUES/e_atis_subject_key_identifier_size/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 63 |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 318J', but common name is 'SHAKEN' |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |


Generated: 28 Nov 23 16:15 UTC