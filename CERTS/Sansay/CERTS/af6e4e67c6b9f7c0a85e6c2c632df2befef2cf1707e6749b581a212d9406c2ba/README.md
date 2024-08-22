# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Voiply, LLC 987J

Tested At: 22 Aug 24 15:53 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -120 day(s)\
Subject: CN=SHAKEN Voiply\\, LLC 987J, OU=Voiply, O=Voiply\\, LLC, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/987J/429C7C70711E3820F0B8E1DEAE6FF32622649698.pem

[View certificate details](https://x509.io/?cert=MIIC0DCCAnegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJklpgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDMyNTAyMjEyOVoXDTI0MDQyNDAyMjEyOVowbTELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEUMBIGA1UECgwLVm9pcGx5LCBMTEMxDzANBgNVBAsMBlZvaXBseTEgMB4GA1UEAwwXU0hBS0VOIFZvaXBseSwgTExDIDk4N0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATS3v1QQ4zeNEMWp31JMZQz%2Bzx6l2lCFEN3lA8Fu2tTE4yc5awgMCPi6366huwm14aea6PJJ6LXRev1I5L9UvpZo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ5ODdKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUoc%2FfIh%2BFNNGdZOCL7ZfdmOov4VswHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIBjXyZb6LEbZ3GFOqm7rva8hZgqg%2BF1h6vm9BydQsS%2FqAiAzJjmhJHTQBvaP3s1Ydd0xn7oKpBoHvLFjskVZkSZZ7g%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 987J', but common name is 'SHAKEN Voiply, LLC 987J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 22 Aug 24 16:06 UTC