# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Bek Communications Cooperative 1604

Tested At: 05 Apr 24 18:52 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 82 day(s)\
Subject: CN=SHAKEN Bek Communications Cooperative 1604, OU=Enterprise Technology Department, O=Bek Communications Cooperative, ST=North Dakota, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/BEK_1064

[View certificate details](https://x509.io/?cert=MIIDETCCArigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkdrQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYyNzExNDMyMloXDTI0MDYyNjExNDMyMlowga0xCzAJBgNVBAYTAlVTMRUwEwYDVQQIDAxOb3J0aCBEYWtvdGExJzAlBgNVBAoMHkJlayBDb21tdW5pY2F0aW9ucyBDb29wZXJhdGl2ZTEpMCcGA1UECwwgRW50ZXJwcmlzZSBUZWNobm9sb2d5IERlcGFydG1lbnQxMzAxBgNVBAMMKlNIQUtFTiBCZWsgQ29tbXVuaWNhdGlvbnMgQ29vcGVyYXRpdmUgMTYwNDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOY6WP%2FMCUgHRqu4S5os%2BfwN4VjVt2fEKh3zS7LidowVDfcJOr74N%2B%2BY0d87019o16WLISRuhzXnX5DrTD7x0SqjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDE2MDQwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBQbh2jKAbh2ZVlE4z9bd%2FCi%2FeRyTjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgeP2uu7GNpF7MBe0YuXApw6Khnuwu4ZAIP%2BBx%2FkJDEr8CIE%2BIuSrYPAvLLsQjv4NgyCan5%2Fe62HXv8YtQLJZyG3Cv)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1604', but common name is 'SHAKEN Bek Communications Cooperative 1604' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 05 Apr 24 19:04 UTC