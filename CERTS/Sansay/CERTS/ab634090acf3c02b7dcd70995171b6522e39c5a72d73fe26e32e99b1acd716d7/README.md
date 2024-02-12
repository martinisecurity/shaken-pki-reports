# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Bulk Solutions, LLC 644J

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 297 day(s)\
Subject: CN=SHAKEN Bulk Solutions\\, LLC 644J, OU=Voice Engineering, O=Bulk Solutions\\, LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://bulkvs-cr.s3.amazonaws.com/644J_2023120501.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6DCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkiM4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTIwNjAzMjE1NloXDTI0MTIwNTAzMjE1NlowgYQxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhEZWxhd2FyZTEcMBoGA1UECgwTQnVsayBTb2x1dGlvbnMsIExMQzEaMBgGA1UECwwRVm9pY2UgRW5naW5lZXJpbmcxKDAmBgNVBAMMH1NIQUtFTiBCdWxrIFNvbHV0aW9ucywgTExDIDY0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASWx5c1SoctyQ8kqKTtSSzEmJIkSx5HjA5uUoAoPVihV5%2B%2F%2BX0knqR2ArOpoGx0kvhTIyjwpsHcKavy9jc2%2FfZgo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2NDRKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU6rwhO4acns6ThqvOZ57HtPEDUPAwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIEsPkpUadiy0npUe%2F6OHx3N6DcM4bX9ELeaeGftFiGBNAiB%2FmwD%2FMMaA1Zr3oMsy0lwAmwXsg5ghCoiyDYTSKbw3Hg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 644J', but common name is 'SHAKEN Bulk Solutions, LLC 644J' |


Generated: 12 Feb 24 17:02 UTC