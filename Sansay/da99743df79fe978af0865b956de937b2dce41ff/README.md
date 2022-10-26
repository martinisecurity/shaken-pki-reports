# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate da99743df79fe978af0865b956de937b2dce41ff
Tested At: 2022-10-26 22:31:16 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 15 day(s)\
Subject: CN=SHAKEN Noble Systems Communications LLC 187J, OU=NOC, O=Noble Systems Communications LLC, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/NobleSys_187J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8zCCApqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTzcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjIwN1oXDTIyMTExMDE3MjIwN1owgY8xCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdHZW9yZ2lhMSkwJwYDVQQKDCBOb2JsZSBTeXN0ZW1zIENvbW11bmljYXRpb25zIExMQzEMMAoGA1UECwwDTk9DMTUwMwYDVQQDDCxTSEFLRU4gTm9ibGUgU3lzdGVtcyBDb21tdW5pY2F0aW9ucyBMTEMgMTg3SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABP%2FmGlJuj%2FstdPV1ajhIPkNsLxwZQGX6idMO%2Fp8KZhhsDFAajP4jO2H2H5v6dRG3QkxY459oHsT7Ovh3AZPOq46jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDE4N0owFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBROIbMJwncdkQ3fRdiI1rOMaa87hDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgM8qB4qYdTfjFV0wJXMR%2FXZ0TpOf3nWa2jphxBjSuy20CIFYXibJEKk7AfVMcJJg9UXrV%2Bo%2BY3FrMCJSCIGDN1bfx)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 187J' |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 26/10/2022 at 22:31:35