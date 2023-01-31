# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Bulk Solutions, LLC 644J

Tested At: 31 Jan 23 21:38 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 309 day(s)\
Subject: CN=SHAKEN Bulk Solutions\\, LLC 644J, OU=Voice Engineering, O=Bulk Solutions\\, LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://bulkvs-cr.s3.amazonaws.com/608300.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6DCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkV0EwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTIwNTIyMjgxOVoXDTIzMTIwNTIyMjgxOVowgYQxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhEZWxhd2FyZTEcMBoGA1UECgwTQnVsayBTb2x1dGlvbnMsIExMQzEaMBgGA1UECwwRVm9pY2UgRW5naW5lZXJpbmcxKDAmBgNVBAMMH1NIQUtFTiBCdWxrIFNvbHV0aW9ucywgTExDIDY0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATk%2B0hTZzXewsPnpEycHO8EmLHkPh16A4Jp7%2FiCZuTc%2FbHIaITlQeNNxp9j%2FXffgdoKQpsla0FGsuKGfNwvF2%2Bro4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2NDRKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUCSTvlWp7PRxJ1qqX0CR66WC7W7YwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIErTM7m7%2FaMq3i8enW74tA1U39bDWdX64mY3ZeD5LmXqAiAOkVPTXkbBO%2Fd0VVZeO%2FX%2Bq2LYOSHq%2Fl3nLiQ4JT%2FZMg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 644J' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 31 Jan 23 21:50 UTC