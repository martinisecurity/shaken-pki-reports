# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Socket Telecom LLC 554a

Tested At: 21 Aug 23 20:12 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -25 day(s)\
Subject: emailAddress=nelsonh@socket.net, CN=SHAKEN Socket Telecom LLC 554a, OU=Operations, O=Socket Telecom LLC, ST=Missouri, C=US, emailAddress=nelsonh@socket.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/554a/order/26_554a_24

[View certificate details](https://understandingwebpki.com/?cert=MIIDAjCCAqmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkdrcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYyNzEyNDIyNFoXDTIzMDcyNzEyNDIyNFowgZ4xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhNaXNzb3VyaTEbMBkGA1UECgwSU29ja2V0IFRlbGVjb20gTExDMRMwEQYDVQQLDApPcGVyYXRpb25zMScwJQYDVQQDDB5TSEFLRU4gU29ja2V0IFRlbGVjb20gTExDIDU1NGExITAfBgkqhkiG9w0BCQEWEm5lbHNvbmhAc29ja2V0Lm5ldDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJVzNFwGlHsxP3OQ7SqTs35lJw8Dkpv8MuvRgqt0kyrR4AXno9%2FeRzyx36RDh1fC8jK%2BiHSt4Vh2hIW4RdU1zVKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDU1NGEwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRWElWJAtp3FKmg4TeSVyQW%2BjwfbjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgFceoUQ18%2BwJokHuPTO4CoNPdqr3aH34yFWTJd7q9W4UCIFYf9A183XN1akj7Vnak7m5L0ATe7zJPuc0pxp44aLzZ)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 554a' |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 21 Aug 23 20:18 UTC