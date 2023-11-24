# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Threshold Communications Inc 563J

Tested At: 24 Nov 23 11:09 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 14 day(s)\
Subject: CN=SHAKEN Threshold Communications Inc 563J, OU=NOC, O=Threshold Communications Inc, ST=Washington, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Threshold_Communications_Inc_563J

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApWgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhfMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTEwNzE2MjIzOFoXDTIzMTIwNzE2MjIzOFowgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMSUwIwYDVQQKDBxUaHJlc2hvbGQgQ29tbXVuaWNhdGlvbnMgSW5jMQwwCgYDVQQLDANOT0MxMTAvBgNVBAMMKFNIQUtFTiBUaHJlc2hvbGQgQ29tbXVuaWNhdGlvbnMgSW5jIDU2M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARMkuoAaHJ0IZKIi3iifdRfyX26s7V3xg0p9qopk0q0CYe1hyb1jEicNMVhXBk%2BWX0GcRj50k81FqyGwPoEvOXSo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ1NjNKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUCtMVhpspSBvylbdCt200kXSDv7UwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIGVP5p5MoMdyifFeoXaezUUg0hGWcuZLkcCzZsKK7BX1AiBr3WvHykLGpU%2BBQ5Z7ctfHBnczZdXaPXYw68HCqycpLw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 563J', but common name is 'SHAKEN Threshold Communications Inc 563J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 24 Nov 23 11:17 UTC