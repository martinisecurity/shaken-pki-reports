# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Carrier One Inc. 705J

Tested At: 31 Oct 22 16:41 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 346 day(s)\
Subject: CN=SHAKEN Carrier One Inc. 705J, OU=Voice NOC, O=Carrier One Inc., ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Carrier_One_Inc._705J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC2jCCAn%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTywwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjAwMVoXDTIzMTAxMTE3MjAwMVowdTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGTAXBgNVBAoMEENhcnJpZXIgT25lIEluYy4xEjAQBgNVBAsMCVZvaWNlIE5PQzElMCMGA1UEAwwcU0hBS0VOIENhcnJpZXIgT25lIEluYy4gNzA1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABB424KM1Fpxx%2FW8QXNzarsuHIfoz1dq4yQZYg2%2FVKqjOnjgP8Ki2ySz0EnsRrod4mu82CD4vsUjdyvxwpjRzFN6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDcwNUowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBR%2Bnltw3EIgRj6WVHpzxsgz%2BBGBIjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAK0Sw7uURG8l3TkCMO4NrFRDsb4IhJWQ148%2B5Cv8w596AiEA8%2ByToCLBLCx1t6vjIXriHWKwwyEwboOVOWGEu04TT90%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 705J' |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31/10/2022 at 16:43:22