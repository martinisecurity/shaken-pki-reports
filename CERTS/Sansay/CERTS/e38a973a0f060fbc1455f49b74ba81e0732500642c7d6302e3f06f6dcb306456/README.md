# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN InteractiveTel, LLC 920J

Tested At: 21 Nov 23 01:45 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -383 day(s)\
Subject: emailAddress=jfindley@interactivetel.com, CN=SHAKEN InteractiveTel\\, LLC 920J, OU=NOC, O=InteractiveTel\\, LLC, ST=Texas, C=US, emailAddress=jfindley@interactivetel.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/920J/order/107_920J_66

[View certificate details](https://understandingwebpki.com/?cert=MIIDsjCCA1igAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTGAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwMzE4NDQxNFoXDTIyMTEwMjE4NDQxNFowgZ8xCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVUZXhhczEcMBoGA1UECgwTSW50ZXJhY3RpdmVUZWwsIExMQzEMMAoGA1UECwwDTk9DMSgwJgYDVQQDDB9TSEFLRU4gSW50ZXJhY3RpdmVUZWwsIExMQyA5MjBKMSowKAYJKoZIhvcNAQkBFhtqZmluZGxleUBpbnRlcmFjdGl2ZXRlbC5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASMcpDf9PfrQJvZ2I%2F245uHQmrHEDVHkPOeja8uEeyKGGQmCO5Z%2BIyec1ynTksWHveCDCq4BKha7Ft4sRx5xmLUo4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDkyMEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRM6ZG50iVTWamn8ZEz7ZrGHrYO2jCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpwwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCm30PHecoqo7W701nju01p4iT5TBXeihvHa7LNTgd5GwIgX%2B5IymELj4qOXfo6FShekJlQFRrQMj1y%2BXUPOuqldFg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 920J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: [2.16.840.1.114569.1.1.3 2.16.840.1.114569.1.1.4] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC