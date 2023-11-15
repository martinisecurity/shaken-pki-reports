# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Mercury Network Corporation 046K

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 260 day(s)\
Remaining Validity Period: 211 day(s)\
Subject: CN=SHAKEN Mercury Network Corporation 046K, OU=Phone Pro, O=Mercury Network Corporation, ST=Michigan, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Mercury_Network_Corporation_046K

[View certificate details](https://understandingwebpki.com/?cert=MIIC8DCCApegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkgfswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDkyNjIwNTE0N1oXDTI0MDYxMjIwNTE0N1owgYwxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhNaWNoaWdhbjEkMCIGA1UECgwbTWVyY3VyeSBOZXR3b3JrIENvcnBvcmF0aW9uMRIwEAYDVQQLDAlQaG9uZSBQcm8xMDAuBgNVBAMMJ1NIQUtFTiBNZXJjdXJ5IE5ldHdvcmsgQ29ycG9yYXRpb24gMDQ2SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPn8NRdOlPWC3n9MzWlIG%2BqByjQPygbtsohO5qvDKcZoK2kUY7ypc8HnMPZoKxFQVL0AvD1%2BPYZaQ2F7utDZzbOjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDA0NkswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRTKTydMPyyo2qm9OcjR5bIBmzZBjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgEU7yImYrPVW%2BKxECZYLvqhnnF8h5kx0XHCDVP9hRrtoCIFugeNqP3wxrYxbPNNFQVb9IknUKsBGlRpR52%2Br46wN1)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 046K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 17:17 UTC