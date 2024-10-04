# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IDT America, Corp 363A

Tested At: 04 Oct 24 15:58 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -71 day(s)\
Subject: CN=SHAKEN IDT America\\, Corp 363A, O=IDT America\\, Corp, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/363A/429C7C70711E3820F0B8E1DEAE6FF3262264A11E.pem

[View certificate details](https://x509.io/?cert=MIICyzCCAnCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkoR4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDYyNDIyMDEzNloXDTI0MDcyNDIyMDEzNlowZjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxGjAYBgNVBAoMEUlEVCBBbWVyaWNhLCBDb3JwMSYwJAYDVQQDDB1TSEFLRU4gSURUIEFtZXJpY2EsIENvcnAgMzYzQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPQM8yPpWKxm0BJ%2BjGRI%2F%2Fn99wUDOlM0pPkwuiTqWy8TLzNxSa1qR7pJYQDwtTuzur7Db7X4HGB5yZxyIiXZWL6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDM2M0EwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMB0GA1UdDgQWBBRtxUb554hu2Cjs4EgzW1ORcFzRVzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhALoWzljlA%2FGejFL0aU6AjgqmrJhJGM%2BoraWuYs%2BV%2BDvQAiEA7axM8u0eTEu5ITU2Cd0Px9ajySSrOhb5UZMg000vR44%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 363A', but common name is 'SHAKEN IDT America, Corp 363A' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 04 Oct 24 16:29 UTC