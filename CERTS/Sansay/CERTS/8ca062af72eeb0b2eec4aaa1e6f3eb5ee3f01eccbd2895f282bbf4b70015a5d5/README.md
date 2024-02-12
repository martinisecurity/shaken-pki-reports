# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Voneto 485K

Tested At: 12 Feb 24 16:52 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 17 day(s)\
Subject: CN=SHAKEN Voneto 485K, OU=Voneto, O=Voneto, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Voneto_485K

[View certificate details](https://understandingwebpki.com/?cert=MIICxzCCAm2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYugwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMwMTE0MTQ1M1oXDTI0MDIyOTE0MTQ1M1owYzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEPMA0GA1UECgwGVm9uZXRvMQ8wDQYDVQQLDAZWb25ldG8xGzAZBgNVBAMMElNIQUtFTiBWb25ldG8gNDg1SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPZ6v5ZOAnEKIifzRZ2Tu7xLqKpRj%2BIkwUWEONAsqcqj71ZLu8e0ws4GWL%2FgoDzzXso4%2BQaxe%2FBbfgGAgwTfTmKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQ4NUswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRfi0v3Iw0lutaW7EgaavSzHUSA6jAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAPYvyRcCZPMuh%2FV0dN7GsMZi1PRwI3cYmzlzKCNUpeEtAiBxxWOJ0nIVdYXfhxwXgTPmNn6bYmUl52Aos4HqR%2BftJA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 485K', but common name is 'SHAKEN Voneto 485K' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC