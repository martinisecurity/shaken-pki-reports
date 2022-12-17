# STIR/SHAKEN CA Ecosystem Compliance

## Certificate VOCALTRANSIT SHAKEN 783J

Tested At: 17 Dec 22 00:05 UTC\
Initial Validity Period: 151 day(s)\
Remaining Validity Period: 88 day(s)\
Subject: CN=VOCALTRANSIT SHAKEN 783J, OU=VOCALTRANSIT, O=IT Vocal LLC, ST=NV, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://shaken.vocaltransit.com/783J/vocaltransit-STIRSHAKEN.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDJDCCAsmgAwIBAgIQKXP1BcOxrhYoPkErqokLmTAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjEwMTUwMDAwMDBaFw0yMzAzMTUwMDAwMDBaMGsxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJOVjEVMBMGA1UECgwMSVQgVm9jYWwgTExDMRUwEwYDVQQLDAxWT0NBTFRSQU5TSVQxITAfBgNVBAMMGFZPQ0FMVFJBTlNJVCBTSEFLRU4gNzgzSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNNBY46aTyPUAghhTzazuhGYFuXgRwprVsrXjm8b9WhWmtcsCIS9ZP0C8JGgTSQmWXXh32xbQ%2B2BVt%2FtCTEE0MmjggE8MIIBODAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUzBN1jNq9%2FuSb8Skf385Xf8EYAiUwHwYDVR0jBBgwFoAUrqFzUYgpVxHKDKn0sQpuTrhLTQcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMBYGCCsGAQUFBwEaBAowCKAGFgQ3ODNKMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAKBggqhkjOPQQDAgNJADBGAiEA6DNLWv5eTy%2FiBAi2dlW6%2Bw1psW8wH%2B2nySmBNGn63YACIQC2OdWDYBveiV%2FDhhsrsqLR9PYHDiw0bRzRK8e40DgumA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 17 Dec 22 00:12 UTC