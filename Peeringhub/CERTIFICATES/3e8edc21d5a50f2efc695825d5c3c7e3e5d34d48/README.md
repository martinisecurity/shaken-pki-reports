# STIR/SHAKEN CA Ecosystem Compliance
## Peeringhub

### Certificate 3e8edc21d5a50f2efc695825d5c3c7e3e5d34d48
Tested At: 2022-10-28 18:54:32 +0000 UTC\
Initial Validity Period: 151 day(s)\
Remaining Validity Period: 138 day(s)\
Subject: CN=VOCALTRANSIT SHAKEN 783J, OU=VOCALTRANSIT, O=IT Vocal LLC, ST=NV, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US

Link: https://shaken.vocaltransit.com/783J/vocaltransit-STIRSHAKEN.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDJDCCAsmgAwIBAgIQKXP1BcOxrhYoPkErqokLmTAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjEwMTUwMDAwMDBaFw0yMzAzMTUwMDAwMDBaMGsxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJOVjEVMBMGA1UECgwMSVQgVm9jYWwgTExDMRUwEwYDVQQLDAxWT0NBTFRSQU5TSVQxITAfBgNVBAMMGFZPQ0FMVFJBTlNJVCBTSEFLRU4gNzgzSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNNBY46aTyPUAghhTzazuhGYFuXgRwprVsrXjm8b9WhWmtcsCIS9ZP0C8JGgTSQmWXXh32xbQ%2B2BVt%2FtCTEE0MmjggE8MIIBODAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUzBN1jNq9%2FuSb8Skf385Xf8EYAiUwHwYDVR0jBBgwFoAUrqFzUYgpVxHKDKn0sQpuTrhLTQcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMBYGCCsGAQUFBwEaBAowCKAGFgQ3ODNKMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAKBggqhkjOPQQDAgNJADBGAiEA6DNLWv5eTy%2FiBAi2dlW6%2Bw1psW8wH%2B2nySmBNGn63YACIQC2OdWDYBveiV%2FDhhsrsqLR9PYHDiw0bRzRK8e40DgumA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 28/10/2022 at 18:55:01