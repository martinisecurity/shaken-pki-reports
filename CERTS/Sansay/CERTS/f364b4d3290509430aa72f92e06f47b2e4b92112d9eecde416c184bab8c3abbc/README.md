# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN XCast Labs 689J

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 32 day(s)\
Remaining Validity Period: 9 day(s)\
Subject: CN=SHAKEN XCast Labs 689J, OU=XCast Labs, O=XCast Labs, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.xcastlabs.net/1700784000/xclsshaken.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC0DCCAnegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhREwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAyMzAwMDAwMFoXDTIzMTEyNDAwMDAwMFowbTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAoMClhDYXN0IExhYnMxEzARBgNVBAsMClhDYXN0IExhYnMxHzAdBgNVBAMMFlNIQUtFTiBYQ2FzdCBMYWJzIDY4OUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQJJpEzyhOIdbNnGdZfTtUr0KUyt%2FKC7rn2oLPf810hl62Bgqa%2Fb%2BL7%2FdtEfiOO%2Bjkg2aG1OUwxZaeRDRP6f1Hio4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2ODlKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUQu1fhldX0A%2ByhR0kbkEsERXho%2B4wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIBASJcrig6fHYuJPBsd5927V9qjRHSNA1Pdn3u6w%2B1b9AiAn%2Fer7Osf6GQP9dPGOtYkPRd7jjCt7ZgIc5CRsoC%2BgLQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 689J' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 18:10 UTC