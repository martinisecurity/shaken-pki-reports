# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN OneStream Networks, LLC 630J

Tested At: 12 Dec 22 23:38 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 303 day(s)\
Subject: CN=SHAKEN OneStream Networks\\, LLC 630J, OU=OneStream Networks, O=OneStream Networks\\, LLC, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/OneStream_Networks_LLC_630J

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTzswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjI1MFoXDTIzMTAxMTE3MjI1MFowgY0xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEgMB4GA1UECgwXT25lU3RyZWFtIE5ldHdvcmtzLCBMTEMxGzAZBgNVBAsMEk9uZVN0cmVhbSBOZXR3b3JrczEsMCoGA1UEAwwjU0hBS0VOIE9uZVN0cmVhbSBOZXR3b3JrcywgTExDIDYzMEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASa11qvWcQ9DavvtESM1uR3AZMEpp%2BkFrLAxVgyINKUMeqH%2BkMfm5Y6%2FHw4EjwOJPLgclhrrvzU935N8NQqRNNVo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MzBKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUHprFro9TA3WtLfg3l%2FJJk7X4Yg4wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCt2hq1OH2QFB5tajeiKkb8cylv7Bt5SQdJQifzn7NGmAIgH0%2FtCGf9Xa3gYeBl6LdwZOh2FUrAgGUpD5q48Wdx7so%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 630J' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 12 Dec 22 23:45 UTC