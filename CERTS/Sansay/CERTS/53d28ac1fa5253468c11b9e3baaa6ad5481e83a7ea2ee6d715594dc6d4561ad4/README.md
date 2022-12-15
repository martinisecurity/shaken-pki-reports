# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Sangoma 777G

Tested At: 15 Dec 22 18:29 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 300 day(s)\
Subject: CN=SHAKEN Sangoma 777G, OU=Network Engineering, O=Sangoma, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Sangoma_777G

[View certificate details](https://understandingwebpki.com/?cert=MIIC0DCCAnegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT0MwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjQ1MVoXDTIzMTAxMTE3MjQ1MVowbTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExEDAOBgNVBAoMB1NhbmdvbWExHDAaBgNVBAsME05ldHdvcmsgRW5naW5lZXJpbmcxHDAaBgNVBAMME1NIQUtFTiBTYW5nb21hIDc3N0cwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT2nggMVgZxEsFFh%2F4UyuEkgLM7jRyW%2B8qFSoEHqN1QJDKMUKcXmoeP%2FrDpl6efZQzqycHvMWabna1nKZNr48bTo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3NzdHMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU6vRL7HFC8a2OdafOCrsLTv9vZuYwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIF7zdh7OadhUrBQB07FUNLXPlEqsjD4Nf5LJ%2BAczvqPGAiAr6iAXzbcus2LO7ofSXjr30KVMscYHMKdYUmCX%2BxwzzQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 777G' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Dec 22 18:35 UTC