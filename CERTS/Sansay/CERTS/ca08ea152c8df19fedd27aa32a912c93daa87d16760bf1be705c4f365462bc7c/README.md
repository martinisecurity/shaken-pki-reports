# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN InteractiveTel, LLC 920J

Tested At: 04 Oct 24 16:14 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -460 day(s)\
Subject: emailAddress=jfindley@interactivetel.com, CN=SHAKEN InteractiveTel\\, LLC 920J, OU=NOC, O=InteractiveTel\\, LLC, ST=Texas, C=US, emailAddress=jfindley@interactivetel.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/920J/order/350_920J_66

[View certificate details](https://x509.io/?cert=MIIDBDCCAqqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkcwIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYwMjA0MjIyNVoXDTIzMDcwMjA0MjIyNVowgZ8xCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVUZXhhczEcMBoGA1UECgwTSW50ZXJhY3RpdmVUZWwsIExMQzEMMAoGA1UECwwDTk9DMSgwJgYDVQQDDB9TSEFLRU4gSW50ZXJhY3RpdmVUZWwsIExMQyA5MjBKMSowKAYJKoZIhvcNAQkBFhtqZmluZGxleUBpbnRlcmFjdGl2ZXRlbC5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASMcpDf9PfrQJvZ2I%2F245uHQmrHEDVHkPOeja8uEeyKGGQmCO5Z%2BIyec1ynTksWHveCDCq4BKha7Ft4sRx5xmLUo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ5MjBKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUTOmRudIlU1mpp%2FGRM%2B2axh62DtowHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIE9vyggshJ%2FTLOYLTjJrynOlx8A8FVuW9D8ZTeL7nWpQAiEAhU4UK1iPr4bA2%2BWpvxDH9z0P9p%2Bx6O6rMwnbs0Vfo4w%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 920J', but common name is 'SHAKEN InteractiveTel, LLC 920J' |


Generated: 04 Oct 24 16:29 UTC