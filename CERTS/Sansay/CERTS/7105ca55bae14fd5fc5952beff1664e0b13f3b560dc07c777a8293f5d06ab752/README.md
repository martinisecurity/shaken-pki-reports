# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Kloud 7 LLC 214K

Tested At: 22 Aug 24 15:54 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 44 day(s)\
Subject: CN=SHAKEN Kloud 7 LLC 214K, OU=Kloud 7 LLC, O=Kloud 7 LLC, ST=Alabama, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Kloud_7_LLC_214K

[View certificate details](https://x509.io/?cert=MIIC0jCCAnegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkgwwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAwNjEzMjUwNFoXDTI0MTAwNTEzMjUwNFowbTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0FsYWJhbWExFDASBgNVBAoMC0tsb3VkIDcgTExDMRQwEgYDVQQLDAtLbG91ZCA3IExMQzEgMB4GA1UEAwwXU0hBS0VOIEtsb3VkIDcgTExDIDIxNEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARdX6kte0L7IPZ3cViHRR4V9%2FAmDh53MgtCLTUp8V00p36Ia0dkPb2WJ1uNDbZubO%2BaCiaFnr2j74cn1d4Yv%2BcZo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyMTRLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUeChyfCFT8YKaszNP4RQNW1mc5z8wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCHcgBxXBM1Wn8zmpWUy3Me6s6yG26o9wStw8C%2BP93OzgIhAPARteNClLPBjVY3Sl6Q1o6giRxv%2FGbDT%2BfT%2BDFYPRIX)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 214K', but common name is 'SHAKEN Kloud 7 LLC 214K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 22 Aug 24 16:06 UTC