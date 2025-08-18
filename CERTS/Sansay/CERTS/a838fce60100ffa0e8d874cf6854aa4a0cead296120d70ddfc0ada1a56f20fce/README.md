# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Dalton Utilities 3139

Tested At: 18 Aug 25 20:26 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 113 day(s)\
Subject: CN=SHAKEN Dalton Utilities 3139, OU=Dalton Utilities 3139, O=Dalton Utilities, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/3139/429C7C70711E3820F0B8E1DEAE6FF3262264B5C2.pem

[View certificate details](https://x509.io/?cert=MIIC5TCCAoygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJktcIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MTIwOTE0MDczNVoXDTI1MTIwOTE0MDczNVowgYExCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdHZW9yZ2lhMRkwFwYDVQQKDBBEYWx0b24gVXRpbGl0aWVzMR4wHAYDVQQLDBVEYWx0b24gVXRpbGl0aWVzIDMxMzkxJTAjBgNVBAMMHFNIQUtFTiBEYWx0b24gVXRpbGl0aWVzIDMxMzkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATb1cq08SNiVy%2FIIak4puTrkSe8lIH9ecty6SaYaNBmpnfPzmcuMVpC2LrRmrP2m4TL7nc85EmjClBdRLIc4HMEo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMTM5MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUN68H9SqAMceH5z5Da6Sdyx3O3UgwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIFCRcdoIazy3YFmZS0B%2BvktuUzyFA1nC19O9UAQhlR2hAiBZi45eRN1NPRXkEgmtnBlGXMkCijr4p%2BURjr%2Bz9gW69w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 3139', but common name is 'SHAKEN Dalton Utilities 3139' |


Generated: 18 Aug 25 21:13 UTC