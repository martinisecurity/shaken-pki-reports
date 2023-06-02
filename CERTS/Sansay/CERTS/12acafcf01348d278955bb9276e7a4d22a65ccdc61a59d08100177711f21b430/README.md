# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Technology Innovation Lab 599J

Tested At: 02 Jun 23 01:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 232 day(s)\
Subject: CN=SHAKEN Technology Innovation Lab 599J, OU=Research and Development, O=Technology Innovation Lab, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://certificates.smartcarrier.io/ho5cfohtbkxwglxbacgnm4qmjnbs9pwu9z772a1n.cer

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FDCCAqKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkXhMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDExOTIyNTAyN1oXDTI0MDExOTIyNTAyN1owgZcxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEiMCAGA1UECgwZVGVjaG5vbG9neSBJbm5vdmF0aW9uIExhYjEhMB8GA1UECwwYUmVzZWFyY2ggYW5kIERldmVsb3BtZW50MS4wLAYDVQQDDCVTSEFLRU4gVGVjaG5vbG9neSBJbm5vdmF0aW9uIExhYiA1OTlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEt6pApUTylivd3Q%2F5ycars6siQgLP8MQID3wNGEnnG3x20w9YXuYhJpftVm7nMf25hqkp8yYn10%2FR6%2Buqx1SyqKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTk5SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHq4Hb25TgadyNQJfLpM9X8OMCjLMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA8%2B46JYQdRxEM4vzd1gElyM5QnuwITiQpAMPj4N4X8vECIAth7xGojLLO2tl3H0Vt5My%2FCog94D8MbGIDM5U%2BZVfa)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 599J' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 02 Jun 23 01:12 UTC