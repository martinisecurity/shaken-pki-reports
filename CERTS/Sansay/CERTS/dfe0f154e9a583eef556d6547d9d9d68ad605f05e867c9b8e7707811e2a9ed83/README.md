# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ComData Solutions 451K

Tested At: 18 Aug 25 20:48 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -279 day(s)\
Subject: CN=SHAKEN ComData Solutions 451K, OU=Comdata, O=ComData Solutions, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Comdata_451K

[View certificate details](https://x509.io/?cert=MIIC2DCCAn2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhiwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTExMzE2NDU1OVoXDTI0MTExMjE2NDU1OVowczELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRowGAYDVQQKDBFDb21EYXRhIFNvbHV0aW9uczEQMA4GA1UECwwHQ29tZGF0YTEmMCQGA1UEAwwdU0hBS0VOIENvbURhdGEgU29sdXRpb25zIDQ1MUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQFvjEQL5m2Vm1YJJV3weRoXQX3hz3cGn9XYZ6k5alrRCx1JlOUnBvkn3ScOy%2BaB%2FuV8vxVx1Yswars1Cix1m6ro4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ0NTFLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUsRucIPf7zoSeWkFN%2FYWdkx2A6pkwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQC0xC8dAE%2B%2FxtjwJCExpG9TkGVrHtg2nMA3j1udEdoxSAIhAIYFB55RJ1AQOeBaAnXIjV9kJGSZScmQVTtQHf0X9jZX)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 451K', but common name is 'SHAKEN ComData Solutions 451K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 18 Aug 25 21:13 UTC