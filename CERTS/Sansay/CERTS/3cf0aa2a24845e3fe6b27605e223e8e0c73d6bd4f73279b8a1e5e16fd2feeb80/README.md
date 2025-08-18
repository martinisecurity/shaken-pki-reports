# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Bulk Solutions, LLC 644J

Tested At: 18 Aug 25 20:04 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -621 day(s)\
Subject: CN=SHAKEN Bulk Solutions\\, LLC 644J, OU=Voice Engineering, O=Bulk Solutions\\, LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://bulkvs-cr.s3.amazonaws.com/644J_2022120501.pem

[View certificate details](https://x509.io/?cert=MIIC6DCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkV0EwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTIwNTIyMjgxOVoXDTIzMTIwNTIyMjgxOVowgYQxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhEZWxhd2FyZTEcMBoGA1UECgwTQnVsayBTb2x1dGlvbnMsIExMQzEaMBgGA1UECwwRVm9pY2UgRW5naW5lZXJpbmcxKDAmBgNVBAMMH1NIQUtFTiBCdWxrIFNvbHV0aW9ucywgTExDIDY0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATk%2B0hTZzXewsPnpEycHO8EmLHkPh16A4Jp7%2FiCZuTc%2FbHIaITlQeNNxp9j%2FXffgdoKQpsla0FGsuKGfNwvF2%2Bro4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2NDRKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUCSTvlWp7PRxJ1qqX0CR66WC7W7YwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIErTM7m7%2FaMq3i8enW74tA1U39bDWdX64mY3ZeD5LmXqAiAOkVPTXkbBO%2Fd0VVZeO%2FX%2Bq2LYOSHq%2Fl3nLiQ4JT%2FZMg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 644J', but common name is 'SHAKEN Bulk Solutions, LLC 644J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 18 Aug 25 21:13 UTC