# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Dalton Utilities 3139

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 174 day(s)\
Remaining Validity Period: 48 day(s)\
Subject: CN=SHAKEN Dalton Utilities 3139, OU=Dalton Utilities 3139, O=Dalton Utilities, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Dalton_Utilities_3139

[View certificate details](https://understandingwebpki.com/?cert=MIIC5jCCAoygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkeVUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDcxMTIxMTE1OFoXDTI0MDEwMTIxMTE1OFowgYExCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdHZW9yZ2lhMRkwFwYDVQQKDBBEYWx0b24gVXRpbGl0aWVzMR4wHAYDVQQLDBVEYWx0b24gVXRpbGl0aWVzIDMxMzkxJTAjBgNVBAMMHFNIQUtFTiBEYWx0b24gVXRpbGl0aWVzIDMxMzkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATb1cq08SNiVy%2FIIak4puTrkSe8lIH9ecty6SaYaNBmpnfPzmcuMVpC2LrRmrP2m4TL7nc85EmjClBdRLIc4HMEo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMTM5MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUN68H9SqAMceH5z5Da6Sdyx3O3UgwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIAnYvIa665lbO2BvEKaHW57NDhHxzO2wcEMcfr0FQWU%2BAiEAylYgqVQopHoVWqhCUxRUXx%2BiI6I6DAEnvt86po3pSsE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 3139' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 18:10 UTC