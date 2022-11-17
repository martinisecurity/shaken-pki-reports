# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Star2Star Communications, LLC 590J

Tested At: 17 Nov 22 19:11 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 310 day(s)\
Subject: CN=SHAKEN Star2Star Communications\\, LLC 590J, OU=BVPROD, O=Star2Star Communications\\, LLC, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/star2star-prod-stirshaken

[View certificate details](https://understandingwebpki.com/?cert=MIIDoDCCA0WgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkSvQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDkyMzAxMTAzNVoXDTIzMDkyMzAxMTAzNVowgYwxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMSYwJAYDVQQKDB1TdGFyMlN0YXIgQ29tbXVuaWNhdGlvbnMsIExMQzEPMA0GA1UECwwGQlZQUk9EMTIwMAYDVQQDDClTSEFLRU4gU3RhcjJTdGFyIENvbW11bmljYXRpb25zLCBMTEMgNTkwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDmML9Ds%2Bs4rHgp6h2fVFdvO%2FmXDpLLYprajtfIT8Mjl0kEsB4G1QmZPJaI6Y7J9L00kbNcLB3vo3ZaSuQqu3ImjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENTkwSjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFL2tlsl2Nv4L6YqZqQ0ddTEZi1RLMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BmjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhANyCYijHGt3wGvYOU4X%2B69oDTIpTD63pXiQxKHHEzRUFAiEA7BjUerLiB1mVVK1f94kAicCvz35SK%2FgzAg569HPjUPc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 590J' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 17 Nov 22 19:21 UTC