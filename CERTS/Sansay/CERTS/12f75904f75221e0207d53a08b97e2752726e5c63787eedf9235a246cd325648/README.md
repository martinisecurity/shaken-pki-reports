# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN California Telecom 319K

Tested At: 15 Nov 23 16:10 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 14 day(s)\
Subject: CN=SHAKEN California Telecom 319K, OU=CaliforniaTelecom, O=California Telecom, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/California_Telecom_319K

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhdMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAzMDE0MDUzNVoXDTIzMTEyOTE0MDUzNVowgYQxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJDYWxpZm9ybmlhIFRlbGVjb20xGjAYBgNVBAsMEUNhbGlmb3JuaWFUZWxlY29tMScwJQYDVQQDDB5TSEFLRU4gQ2FsaWZvcm5pYSBUZWxlY29tIDMxOUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASARI%2FByFwiJV9tLoOKjSNGDY97dFOsdg6%2FNSu1Ve0AI6Q7yvIajg5pQwKhWqlqGDbEc3BT7cPXBoCWkglSvSv%2Fo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMTlLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUh%2F6stqcI6yhZbhr4pJsRGsNYUT0wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCHsFI4kB88jEFiXhmbKO7fFPO2cJx41yy5G2s8ZXJ7wAIgazAd8%2FJ6%2F9xvwSAjfm%2FQbOVDdshsyrJwCOOy36s2Xyc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 319K' |


Generated: 15 Nov 23 16:51 UTC