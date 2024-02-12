# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Telcentris Inc. dba Voxox 696J

Tested At: 12 Feb 24 16:52 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -372 day(s)\
Subject: CN=SHAKEN Telcentris Inc. dba Voxox 696J, OU=NOC, O=Telcentris Inc. dba Voxox, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Voxox_696J

[View certificate details](https://understandingwebpki.com/?cert=MIIC6jCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkW1swCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDEwNjE2MTMzN1oXDTIzMDIwNTE2MTMzN1owgYQxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMSIwIAYDVQQKDBlUZWxjZW50cmlzIEluYy4gZGJhIFZveG94MQwwCgYDVQQLDANOT0MxLjAsBgNVBAMMJVNIQUtFTiBUZWxjZW50cmlzIEluYy4gZGJhIFZveG94IDY5NkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATbHNz%2B%2Bg%2B%2F6gYnvUjtzRASJlG0qJ7Mu1NeSe3Wr5t15kO7M%2BARvH7%2BoaYbfls69EoqBa97VOyAjrFOdqDH%2FOuEo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2OTZKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU0NP78Nxc95rRRl%2FkhaK3%2FB93ZawwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCqYB9vw8VlkXFnFD6VlMvRiAOwl6HlXS3a0TlzxzK4dgIhAJB1tgp1%2FIIwkO2diE%2FNPm4EyEKYRMrkq%2FTUXsR%2FvnlN)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 696J', but common name is 'SHAKEN Telcentris Inc. dba Voxox 696J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC