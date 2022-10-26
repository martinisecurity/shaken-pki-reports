# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 9c29e43ad0d0a8b99fab79bd5286de3be29036d2
Tested At: 2022-10-26 21:14:18 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 350 day(s)\
Subject: CN=SHAKEN Sangoma 777G, OU=Network Engineering, O=Sangoma, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Sangoma_777G

View: [Click to view](https://understandingwebpki.com/?cert=MIIC0DCCAnegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT0MwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjQ1MVoXDTIzMTAxMTE3MjQ1MVowbTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExEDAOBgNVBAoMB1NhbmdvbWExHDAaBgNVBAsME05ldHdvcmsgRW5naW5lZXJpbmcxHDAaBgNVBAMME1NIQUtFTiBTYW5nb21hIDc3N0cwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT2nggMVgZxEsFFh%2F4UyuEkgLM7jRyW%2B8qFSoEHqN1QJDKMUKcXmoeP%2FrDpl6efZQzqycHvMWabna1nKZNr48bTo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3NzdHMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU6vRL7HFC8a2OdafOCrsLTv9vZuYwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIF7zdh7OadhUrBQB07FUNLXPlEqsjD4Nf5LJ%2BAczvqPGAiAr6iAXzbcus2LO7ofSXjr30KVMscYHMKdYUmCX%2BxwzzQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 777G' |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 26/10/2022 at 21:14:23