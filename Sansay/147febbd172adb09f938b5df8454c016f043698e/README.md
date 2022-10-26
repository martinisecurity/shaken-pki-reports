# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 147febbd172adb09f938b5df8454c016f043698e
Tested At: 2022-10-26 20:19:46 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 350 day(s)\
Subject: CN=SHAKEN OneStream Networks\\, LLC 630J, OU=OneStream Networks, O=OneStream Networks\\, LLC, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/OneStream_Networks_LLC_630J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTzswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjI1MFoXDTIzMTAxMTE3MjI1MFowgY0xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEgMB4GA1UECgwXT25lU3RyZWFtIE5ldHdvcmtzLCBMTEMxGzAZBgNVBAsMEk9uZVN0cmVhbSBOZXR3b3JrczEsMCoGA1UEAwwjU0hBS0VOIE9uZVN0cmVhbSBOZXR3b3JrcywgTExDIDYzMEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASa11qvWcQ9DavvtESM1uR3AZMEpp%2BkFrLAxVgyINKUMeqH%2BkMfm5Y6%2FHw4EjwOJPLgclhrrvzU935N8NQqRNNVo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MzBKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUHprFro9TA3WtLfg3l%2FJJk7X4Yg4wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCt2hq1OH2QFB5tajeiKkb8cylv7Bt5SQdJQifzn7NGmAIgH0%2FtCGf9Xa3gYeBl6LdwZOh2FUrAgGUpD5q48Wdx7so%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 630J' |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 26/10/2022 at 20:21:30