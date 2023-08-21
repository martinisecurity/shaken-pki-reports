# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN California Telecom 319K

Tested At: 21 Aug 23 20:14 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 66 day(s)\
Subject: CN=SHAKEN California Telecom 319K, OU=CaliforniaTelecom, O=California Telecom, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/California_Telecom_319K

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU0IwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjE5NDAzNFoXDTIzMTAyNjE5NDAzNFowgYQxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJDYWxpZm9ybmlhIFRlbGVjb20xGjAYBgNVBAsMEUNhbGlmb3JuaWFUZWxlY29tMScwJQYDVQQDDB5TSEFLRU4gQ2FsaWZvcm5pYSBUZWxlY29tIDMxOUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASARI%2FByFwiJV9tLoOKjSNGDY97dFOsdg6%2FNSu1Ve0AI6Q7yvIajg5pQwKhWqlqGDbEc3BT7cPXBoCWkglSvSv%2Fo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMTlLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUh%2F6stqcI6yhZbhr4pJsRGsNYUT0wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCroiVhbtXoRQ3z3kQlAYcrIIlooSdJHb2fN9UDT6y8SwIgXquhtmUvZRf3d5glDF8ERPNiM4ws1V%2FT2YAEDBsoOZ8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 319K' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 21 Aug 23 20:18 UTC