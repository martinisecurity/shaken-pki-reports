# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 9c29e43ad0d0a8b99fab79bd5286de3be29036d2
Tested At: 2022-10-27 21:26:28 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 349 day(s)\
Subject: CN=SHAKEN Sangoma 777G, OU=Network Engineering, O=Sangoma, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Sangoma_777G

View: [Click to view](https://understandingwebpki.com/?cert=MIIC0DCCAnegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT0MwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjQ1MVoXDTIzMTAxMTE3MjQ1MVowbTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExEDAOBgNVBAoMB1NhbmdvbWExHDAaBgNVBAsME05ldHdvcmsgRW5naW5lZXJpbmcxHDAaBgNVBAMME1NIQUtFTiBTYW5nb21hIDc3N0cwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT2nggMVgZxEsFFh%2F4UyuEkgLM7jRyW%2B8qFSoEHqN1QJDKMUKcXmoeP%2FrDpl6efZQzqycHvMWabna1nKZNr48bTo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3NzdHMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU6vRL7HFC8a2OdafOCrsLTv9vZuYwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIF7zdh7OadhUrBQB07FUNLXPlEqsjD4Nf5LJ%2BAczvqPGAiAr6iAXzbcus2LO7ofSXjr30KVMscYHMKdYUmCX%2BxwzzQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 777G' |

Generated: 27/10/2022 at 21:27:34