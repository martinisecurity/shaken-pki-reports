# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Socket Telecom LLC 554a

Tested At: 04 Oct 24 16:06 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -348 day(s)\
Subject: emailAddress=nelsonh@socket.net, CN=SHAKEN Socket Telecom LLC 554a, OU=Operations, O=Socket Telecom LLC, ST=Missouri, C=US, emailAddress=nelsonh@socket.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/554a/order/114_554a_24

[View certificate details](https://x509.io/?cert=MIIDBDCCAqmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkgVgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDkyMjEwMTI1NloXDTIzMTAyMjEwMTI1NlowgZ4xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhNaXNzb3VyaTEbMBkGA1UECgwSU29ja2V0IFRlbGVjb20gTExDMRMwEQYDVQQLDApPcGVyYXRpb25zMScwJQYDVQQDDB5TSEFLRU4gU29ja2V0IFRlbGVjb20gTExDIDU1NGExITAfBgkqhkiG9w0BCQEWEm5lbHNvbmhAc29ja2V0Lm5ldDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJVzNFwGlHsxP3OQ7SqTs35lJw8Dkpv8MuvRgqt0kyrR4AXno9%2FeRzyx36RDh1fC8jK%2BiHSt4Vh2hIW4RdU1zVKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDU1NGEwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRWElWJAtp3FKmg4TeSVyQW%2BjwfbjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAJkcBfaYTtpa4KkQfGI9COuQystXr07AJbH7A5qXZybLAiEA4GWKOoUjFHbNm4bix6Flg8bqGyPwGuUlVHsBo1%2FLR3I%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 554a', but common name is 'SHAKEN Socket Telecom LLC 554a' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '554a' contains characters other than uppercase letters and numbers |


Generated: 04 Oct 24 16:29 UTC