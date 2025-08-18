# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Technology Innovation Lab 599J

Tested At: 18 Aug 25 20:15 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -576 day(s)\
Subject: CN=SHAKEN Technology Innovation Lab 599J, OU=Research and Development, O=Technology Innovation Lab, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://certificates.smartcarrier.io/ho5cfohtbkxwglxbacgnm4qmjnbs9pwu9z772a1n.cer

[View certificate details](https://x509.io/?cert=MIIC%2FDCCAqKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkXhMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDExOTIyNTAyN1oXDTI0MDExOTIyNTAyN1owgZcxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEiMCAGA1UECgwZVGVjaG5vbG9neSBJbm5vdmF0aW9uIExhYjEhMB8GA1UECwwYUmVzZWFyY2ggYW5kIERldmVsb3BtZW50MS4wLAYDVQQDDCVTSEFLRU4gVGVjaG5vbG9neSBJbm5vdmF0aW9uIExhYiA1OTlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEt6pApUTylivd3Q%2F5ycars6siQgLP8MQID3wNGEnnG3x20w9YXuYhJpftVm7nMf25hqkp8yYn10%2FR6%2Buqx1SyqKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTk5SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHq4Hb25TgadyNQJfLpM9X8OMCjLMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA8%2B46JYQdRxEM4vzd1gElyM5QnuwITiQpAMPj4N4X8vECIAth7xGojLLO2tl3H0Vt5My%2FCog94D8MbGIDM5U%2BZVfa)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 599J', but common name is 'SHAKEN Technology Innovation Lab 599J' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 18 Aug 25 21:13 UTC