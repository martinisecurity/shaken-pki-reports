# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN SouthPoint Communications 205k

Tested At: 28 Apr 23 02:07 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 0 day(s)\
Subject: emailAddress=chris.duncan@soundcurve.com, CN=SHAKEN SouthPoint Communications 205k, OU=Operations, O=SouthPoint Communications, ST=California, C=US, emailAddress=chris.duncan@soundcurve.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/205k/order/247_205k_99

[View certificate details](https://understandingwebpki.com/?cert=MIIDHDCCAsKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZ5MwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMyODE3MDg1OVoXDTIzMDQyNzE3MDg1OVowgbcxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMSIwIAYDVQQKDBlTb3V0aFBvaW50IENvbW11bmljYXRpb25zMRMwEQYDVQQLDApPcGVyYXRpb25zMS4wLAYDVQQDDCVTSEFLRU4gU291dGhQb2ludCBDb21tdW5pY2F0aW9ucyAyMDVrMSowKAYJKoZIhvcNAQkBFhtjaHJpcy5kdW5jYW5Ac291bmRjdXJ2ZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASfJ3OhB3xnPprzTOKiNwTougqlsCZMeV6M0kcwk8F%2BAkWyorpqFfoEGE65rnxAU2fZJu74W5lCPZtiPEyGdxs1o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyMDVrMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU9%2BIQO2rwmus8MaQutkdnPhX%2F4aswHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIB%2BbxSvWBOrq%2F%2BFCrT%2FfY127EX%2BugO8mYbCmL%2BhO5jHtAiEA%2BQYeAh829kbbBkT7%2FO%2FzBZMXUBiyGKKLSh9eOlELbB0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 205k' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 28 Apr 23 02:17 UTC