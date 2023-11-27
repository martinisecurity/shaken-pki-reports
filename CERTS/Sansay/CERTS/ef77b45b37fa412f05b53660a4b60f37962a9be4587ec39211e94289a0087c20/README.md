# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Nextiva, Inc 914H

Tested At: 27 Nov 23 23:20 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 178 day(s)\
Subject: CN=SHAKEN Nextiva\\, Inc 914H, OU=UCaaS, O=Nextiva\\, Inc, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Nextiva_Shaken

[View certificate details](https://understandingwebpki.com/?cert=MIICzTCCAnOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkcaUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDUyNDE0NDkxN1oXDTI0MDUyMzE0NDkxN1owaTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0FyaXpvbmExFTATBgNVBAoMDE5leHRpdmEsIEluYzEOMAwGA1UECwwFVUNhYVMxITAfBgNVBAMMGFNIQUtFTiBOZXh0aXZhLCBJbmMgOTE0SDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAdYfVWHof3n%2Bv4mgD2M5rKj7MchoK7T%2BLmyN1YtomASnkMjjBCCSRwOPEMSzyLyfFdwkPg0AM99gWweZ2hJp42jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDkxNEgwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTHDafH2cITxJnXTEUPC3QpR0fkjjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgLeT8HSUYGKSrN3Quf2z0D%2FHuVqPhEd%2FO9JvrAlN8j5cCIQD%2B6q4iOlsjLP%2F9hJme0%2BiTyx6jTAvXsH45hJ9udNCM%2Fg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 914H', but common name is 'SHAKEN Nextiva, Inc 914H' |


Generated: 27 Nov 23 23:28 UTC