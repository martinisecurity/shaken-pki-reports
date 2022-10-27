# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 4627f95dd2ca4c2d4d33dd81994a19375739171c
Tested At: 2022-10-27 18:56:01 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 349 day(s)\
Subject: CN=SHAKEN Carrier One Inc. 705J, OU=Voice NOC, O=Carrier One Inc., ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Carrier_One_Inc._705J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC2jCCAn%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTywwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjAwMVoXDTIzMTAxMTE3MjAwMVowdTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGTAXBgNVBAoMEENhcnJpZXIgT25lIEluYy4xEjAQBgNVBAsMCVZvaWNlIE5PQzElMCMGA1UEAwwcU0hBS0VOIENhcnJpZXIgT25lIEluYy4gNzA1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABB424KM1Fpxx%2FW8QXNzarsuHIfoz1dq4yQZYg2%2FVKqjOnjgP8Ki2ySz0EnsRrod4mu82CD4vsUjdyvxwpjRzFN6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDcwNUowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBR%2Bnltw3EIgRj6WVHpzxsgz%2BBGBIjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAK0Sw7uURG8l3TkCMO4NrFRDsb4IhJWQ148%2B5Cv8w596AiEA8%2ByToCLBLCx1t6vjIXriHWKwwyEwboOVOWGEu04TT90%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 705J' |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 27/10/2022 at 18:57:26