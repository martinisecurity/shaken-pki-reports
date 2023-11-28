# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Bulk Solutions, LLC 644J

Tested At: 28 Nov 23 10:17 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -76 day(s)\
Subject: CN=SHAKEN Bulk Solutions\\, LLC 644J, OU=Voice Engineering, O=Bulk Solutions\\, LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://bulkvs-cr.s3.amazonaws.com/644J_20210615001.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDmDCCAz2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkSZAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDkxMjE5Mzc1N1oXDTIzMDkxMjE5Mzc1N1owgYQxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhEZWxhd2FyZTEcMBoGA1UECgwTQnVsayBTb2x1dGlvbnMsIExMQzEaMBgGA1UECwwRVm9pY2UgRW5naW5lZXJpbmcxKDAmBgNVBAMMH1NIQUtFTiBCdWxrIFNvbHV0aW9ucywgTExDIDY0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQJ%2BfKPjj1zc6BLNOXT3Te%2BjxZMizox8kcU5hoHmgDd9ZKvAPdIn2FpygHxROC0YUlbpWmyNyfv4QM2RmgceC8Jo4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDY0NEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSEqLJYf%2BOMsVCAvZ3RzON79P2shjCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpowRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCVLALwPgav0e7V44he6Z6TGNFPAtEfvQ6KYeK7GUllpwIhALBcdowOlk44wBjZ%2BvIC5coOAYeiTOMralcYF4DSdqEX)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 644J', but common name is 'SHAKEN Bulk Solutions, LLC 644J' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 10:53 UTC