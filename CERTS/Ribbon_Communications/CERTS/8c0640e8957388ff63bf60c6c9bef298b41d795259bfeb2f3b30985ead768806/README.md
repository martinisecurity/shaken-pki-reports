# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Netfortris SHAKEN 8886

Tested At: 06 Jul 23 14:06 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 337 day(s)\
Subject: CN=Netfortris SHAKEN 8886, OU=VOIP, O=Netfortris, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/t0CiOIjnRz/NFJune072023-3a5cc337262fa27b6e6a3d8948fc365e

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQOlzDNyYvontuaj2JSPw2XjAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNjA3MjAyNTIzWhcNMjQwNjA2MjAyNTIyWjBSMQswCQYDVQQGEwJVUzETMBEGA1UECgwKTmV0Zm9ydHJpczENMAsGA1UECwwEVk9JUDEfMB0GA1UEAwwWTmV0Zm9ydHJpcyBTSEFLRU4gODg4NjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGTpSsSeFsWKpfjLhUAeYrXYfXIDS6vpwtugqdIEVBcWrdminS5ts2AF91uVm7JDcoqA51cNi8CR08yzisrPLgWjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQULkodxs1bf3IRKr3bvmIiNIAtKRcwGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQ4ODg2MAoGCCqGSM49BAMCA0gAMEUCIQDVUQ7th4xjB%2FdCeQNsWO2wWoxXXQ%2BayGg2hjExxC%2ByMAIgCFZlg2NsMut2ECiBjckUDluRvzqmr1bgm3B6jf%2F2BS4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 06 Jul 23 14:08 UTC