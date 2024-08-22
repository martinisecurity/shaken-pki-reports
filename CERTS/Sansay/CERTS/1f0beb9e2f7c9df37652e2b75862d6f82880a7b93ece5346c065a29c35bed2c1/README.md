# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Primo Dialler LLC 249K

Tested At: 22 Aug 24 15:15 UTC\
Initial Validity Period: 40 day(s)\
Remaining Validity Period: -523 day(s)\
Subject: CN=SHAKEN Primo Dialler LLC 249K, OU=Service Provider, O=Primo Dialler LLC, ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: http://5.161.152.107/191c4c42dd7fa6115e84100637e42c99.cer

[View certificate details](https://x509.io/?cert=MIIC4TCCAoigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYJ8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIwNjEyNTUwMloXDTIzMDMxODEyNTUwMlowfjELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGjAYBgNVBAoMEVByaW1vIERpYWxsZXIgTExDMRkwFwYDVQQLDBBTZXJ2aWNlIFByb3ZpZGVyMSYwJAYDVQQDDB1TSEFLRU4gUHJpbW8gRGlhbGxlciBMTEMgMjQ5SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEX0eCLPJmDebub9O3tCmF%2FA3E2ov0mT0WdnYxu5b%2F3qPs4rkHiDEmJoSNcRzfGkyNRBm%2FdHn02Fb93jbJ%2Fxo%2B6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDI0OUswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBSFegELszR6DWADvE2pV9%2BBE2Tj%2BjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgO7WyJuJagr9N99CUmZqah%2BfQvvZzw0eLgnALpFpugHsCICGTrki4zf5SBvKxxb6Z1RT8lHdIYaYp9sT%2FTVV%2BW%2Bze)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 249K', but common name is 'SHAKEN Primo Dialler LLC 249K' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 15:44 UTC