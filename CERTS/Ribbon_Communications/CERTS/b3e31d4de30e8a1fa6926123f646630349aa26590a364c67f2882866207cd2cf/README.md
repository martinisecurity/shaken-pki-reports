# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ThreeriverTelco SHAKEN 1525

Tested At: 15 Nov 23 16:37 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 169 day(s)\
Subject: CN=ThreeriverTelco SHAKEN 1525, OU=STI, O=ThreeriverTelco, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/0X3WXxsVRz/STI-202305-1525-5f4ab88c046b1cfd58ff229698fff463

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQX0q4jARrHP1Y%2FyKWmP%2F0YzAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNTAzMTUxODEwWhcNMjQwNTAyMTUxODA5WjBbMQswCQYDVQQGEwJVUzEYMBYGA1UECgwPVGhyZWVyaXZlclRlbGNvMQwwCgYDVQQLDANTVEkxJDAiBgNVBAMMG1RocmVlcml2ZXJUZWxjbyBTSEFLRU4gMTUyNTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNnjwnvymka7omLqKhEvi%2Bn3fJF1Do7jQOmuPvmS7c41PW1pOdRpUIm%2BO8hpkWE3mG2QaF18RM5DQOmIN8OwFmSjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUoNp5k3vi85I08BGaOTWB0BRv8YowGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQxNTI1MAoGCCqGSM49BAMCA0gAMEUCIQDMxK0pAH84NlJfKy1nnjqF4rwOM3I%2FTstP7H0r9%2B2MwQIgNx0PJ4kXeA%2BSZa1cveWQOpGlIFTjb311y2adkSEoy%2F0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 17:17 UTC