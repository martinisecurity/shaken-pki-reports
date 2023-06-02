# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Intelegrated, LLC 325K

Tested At: 02 Jun 23 01:08 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 146 day(s)\
Subject: CN=SHAKEN Intelegrated\\, LLC 325K, OU=Intelegrated\\, LLC, O=Intelegrated\\, LLC, ST=South Dakota, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Intelegrated_LLC_325K102522

[View certificate details](https://understandingwebpki.com/?cert=MIIC6DCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUwswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNTIxMTEyOVoXDTIzMTAyNTIxMTEyOVowgYQxCzAJBgNVBAYTAlVTMRUwEwYDVQQIDAxTb3V0aCBEYWtvdGExGjAYBgNVBAoMEUludGVsZWdyYXRlZCwgTExDMRowGAYDVQQLDBFJbnRlbGVncmF0ZWQsIExMQzEmMCQGA1UEAwwdU0hBS0VOIEludGVsZWdyYXRlZCwgTExDIDMyNUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARnKFiF00QjKfzUYQ3rWSQ3yn%2F8LuOuDcHwmMr4MZpmdJ5R1R9IFyQDE6hXV13tis8Nvqaef6mO%2F1l4vI9KQf7Vo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMjVLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUPpD6ttoZsaH%2BNf7AyHC5OEwwVdEwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIE1r5egjJFuR3MAnA033%2Fr1qldv%2BPiHDGANr5MmW17wnAiARPw02kpJvTRhvyHiYYMDmGwJ0%2FkZOUIfHOhUYx%2BH4ww%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 325K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 02 Jun 23 01:12 UTC