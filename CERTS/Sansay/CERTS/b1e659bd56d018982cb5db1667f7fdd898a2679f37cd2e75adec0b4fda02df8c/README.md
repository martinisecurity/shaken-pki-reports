# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Nextiva, Inc 914H

Tested At: 22 Aug 24 15:52 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 259 day(s)\
Subject: CN=SHAKEN Nextiva\\, Inc 914H, OU=UCaaS, O=Nextiva\\, Inc, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/914H/429C7C70711E3820F0B8E1DEAE6FF32622649B8F.pem

[View certificate details](https://x509.io/?cert=MIICzjCCAnOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkm48wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDUwODAwMDAwMFoXDTI1MDUwODAwMDAwMFowaTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0FyaXpvbmExFTATBgNVBAoMDE5leHRpdmEsIEluYzEOMAwGA1UECwwFVUNhYVMxITAfBgNVBAMMGFNIQUtFTiBOZXh0aXZhLCBJbmMgOTE0SDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAdYfVWHof3n%2Bv4mgD2M5rKj7MchoK7T%2BLmyN1YtomASnkMjjBCCSRwOPEMSzyLyfFdwkPg0AM99gWweZ2hJp42jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDkxNEgwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTHDafH2cITxJnXTEUPC3QpR0fkjjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAOkKMBExZI5Ywm7YL%2BNzvcpzYAbKbWBiQgWI5%2FWz%2BUzTAiEAxFEOW%2BI40fwYvRFUF7%2BtkQOjjuL9zZpla8yjJ1j6l8A%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 914H', but common name is 'SHAKEN Nextiva, Inc 914H' |


Generated: 22 Aug 24 16:06 UTC