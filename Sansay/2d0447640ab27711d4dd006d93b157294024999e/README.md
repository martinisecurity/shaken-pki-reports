# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 2d0447640ab27711d4dd006d93b157294024999e
Tested At: 2022-10-26 22:30:02 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 15 day(s)\
Subject: CN=SHAKEN Star2Star Communications\\, LLC 590J, OU=BVPROD, O=Star2Star Communications\\, LLC, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/star2star_stirshaken_20220922

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8DCCApegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT00wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE5MjgzNFoXDTIyMTExMDE5MjgzNFowgYwxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMSYwJAYDVQQKDB1TdGFyMlN0YXIgQ29tbXVuaWNhdGlvbnMsIExMQzEPMA0GA1UECwwGQlZQUk9EMTIwMAYDVQQDDClTSEFLRU4gU3RhcjJTdGFyIENvbW11bmljYXRpb25zLCBMTEMgNTkwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDmML9Ds%2Bs4rHgp6h2fVFdvO%2FmXDpLLYprajtfIT8Mjl0kEsB4G1QmZPJaI6Y7J9L00kbNcLB3vo3ZaSuQqu3ImjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDU5MEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBS9rZbJdjb%2BC%2BmKmakNHXUxGYtUSzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgNXU%2FZVitKIShYwRSn67WlBi5ZyFB0psPHR1T7hTLobkCIG8LTHTO8z7Cpvc1uUFznwVwRLf04DbrRaagxF%2B9kCja)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 590J' |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 26/10/2022 at 22:31:35