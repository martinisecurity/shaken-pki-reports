# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Threshold Communications Inc 563J

Tested At: 26 Aug 24 18:37 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -199 day(s)\
Subject: CN=SHAKEN Threshold Communications Inc 563J, OU=NOC, O=Threshold Communications Inc, ST=Washington, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Threshold_Communications_Inc_563J

[View certificate details](https://x509.io/?cert=MIIC7zCCApWgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkjUIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDEwOTIwMjkyMloXDTI0MDIwODIwMjkyMlowgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMSUwIwYDVQQKDBxUaHJlc2hvbGQgQ29tbXVuaWNhdGlvbnMgSW5jMQwwCgYDVQQLDANOT0MxMTAvBgNVBAMMKFNIQUtFTiBUaHJlc2hvbGQgQ29tbXVuaWNhdGlvbnMgSW5jIDU2M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARMkuoAaHJ0IZKIi3iifdRfyX26s7V3xg0p9qopk0q0CYe1hyb1jEicNMVhXBk%2BWX0GcRj50k81FqyGwPoEvOXSo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ1NjNKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUCtMVhpspSBvylbdCt200kXSDv7UwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIDhyojPADQqJm1P%2F7msk8kxhU3MxpY4ze%2Fxe1uiJO3HTAiEAxqxZvwsUd3x5lte8eEmbDe3uIIVL1gMu0AqyS99z5Xo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 563J', but common name is 'SHAKEN Threshold Communications Inc 563J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 26 Aug 24 18:49 UTC