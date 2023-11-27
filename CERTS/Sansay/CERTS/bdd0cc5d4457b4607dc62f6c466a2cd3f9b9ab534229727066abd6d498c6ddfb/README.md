# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Swift Telco LLC 452K

Tested At: 27 Nov 23 23:20 UTC\
Initial Validity Period: 2 day(s)\
Remaining Validity Period: -284 day(s)\
Subject: CN=SHAKEN Swift Telco LLC 452K, OU=SWT, O=Swift Telco LLC, ST=South Carolina, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Swift_telco_2_D

[View certificate details](https://understandingwebpki.com/?cert=MIIC2DCCAn6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYXEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIxNDIxMjUwNVoXDTIzMDIxNjIxMjUwNVowdDELMAkGA1UEBhMCVVMxFzAVBgNVBAgMDlNvdXRoIENhcm9saW5hMRgwFgYDVQQKDA9Td2lmdCBUZWxjbyBMTEMxDDAKBgNVBAsMA1NXVDEkMCIGA1UEAwwbU0hBS0VOIFN3aWZ0IFRlbGNvIExMQyA0NTJLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEkT6gzQw0zjt%2FAH6COkEIOERte%2BTkPETA6lsKPDX7u6K9pdvdHUW3bguglpgIOAA8XQFl31HnyTXUj%2BsE%2BXl20qOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENDUySzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFNgDrM%2BnSFrweJo5XIbccO7IoE5hMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiA6mGc8f6xrTrjCckvmYAWnKBau4j%2BpSuVjiqQnvD7gjQIhAISg0qaRS4Ik%2FeIOyE4LddPJC%2BMzMh2%2BRXIEBUQ3ZxFs)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 452K', but common name is 'SHAKEN Swift Telco LLC 452K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 23:28 UTC