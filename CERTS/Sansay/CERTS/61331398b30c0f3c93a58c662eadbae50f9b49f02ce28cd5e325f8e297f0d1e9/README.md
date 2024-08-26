# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Matrix 9451

Tested At: 26 Aug 24 18:37 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -199 day(s)\
Subject: CN=SHAKEN Matrix 9451, OU=Engineering, O=Matrix, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Lingo-9451

[View certificate details](https://x509.io/?cert=MIICxDCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkjVkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDEwOTIxMDQzNFoXDTI0MDIwODIxMDQzNFowYTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZNYXRyaXgxFDASBgNVBAsMC0VuZ2luZWVyaW5nMRswGQYDVQQDDBJTSEFLRU4gTWF0cml4IDk0NTEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATKK19BR8TbDyCPk5h2YzHgsh%2BJs%2BPyP%2FPJ%2F4T5HwXs7u43XyL9Z%2BZZ2RNErU5qraZmexjDBbnguUi0kNUn%2B59So4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ5NDUxMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUK5Ds9QZFxG%2F%2BcfvTtwaOpOm5qd0wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIHv7R6N%2B8YRnbd1cr%2FYaYVGnT%2BMreJE7PuJgk7ptTP3IAiB53bE%2B0B2cwbLEVrB7nSxUonKArdFYt7j8GD0U0f66vg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 9451', but common name is 'SHAKEN Matrix 9451' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 26 Aug 24 18:49 UTC