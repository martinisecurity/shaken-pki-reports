# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Matrix 9451

Tested At: 04 Oct 24 16:14 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -63 day(s)\
Subject: CN=SHAKEN Matrix 9451, OU=Engineering, O=Matrix, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/9451/429C7C70711E3820F0B8E1DEAE6FF3262264A20E.pem

[View certificate details](https://x509.io/?cert=MIICxjCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkog4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDcwMjE3MTQzNFoXDTI0MDgwMTE3MTQzNFowYTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZNYXRyaXgxFDASBgNVBAsMC0VuZ2luZWVyaW5nMRswGQYDVQQDDBJTSEFLRU4gTWF0cml4IDk0NTEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATKK19BR8TbDyCPk5h2YzHgsh%2BJs%2BPyP%2FPJ%2F4T5HwXs7u43XyL9Z%2BZZ2RNErU5qraZmexjDBbnguUi0kNUn%2B59So4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ5NDUxMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUK5Ds9QZFxG%2F%2BcfvTtwaOpOm5qd0wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQDqznTzfPq5eoKqnBzh828asX1bkkUG6xF3FQs4H4Ll5wIhALQk7IxrpGgD1K8ayBYbw%2FdyXLViyIfJZDJQD9twbMGY)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 9451', but common name is 'SHAKEN Matrix 9451' |


Generated: 04 Oct 24 16:29 UTC