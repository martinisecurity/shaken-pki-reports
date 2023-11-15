# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Fulton SHAKEN 0455

Tested At: 15 Nov 23 16:37 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 220 day(s)\
Subject: CN=Fulton SHAKEN 0455, OU=STI, O=Fulton Telephone Company, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/7crfAxXVRz/STI-0455-2023-5d0904544e3802776d08740599a8aa12

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApigAwIBAgIQXQkEVE44AndtCHQFmaiqEjAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNjIyMTkxMTI1WhcNMjQwNjIxMTkxMTI0WjBbMQswCQYDVQQGEwJVUzEhMB8GA1UECgwYRnVsdG9uIFRlbGVwaG9uZSBDb21wYW55MQwwCgYDVQQLDANTVEkxGzAZBgNVBAMMEkZ1bHRvbiBTSEFLRU4gMDQ1NTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNQJUDUrqmO%2FU8xEhK%2BA1%2BRiwdUQD7O29JHZAVx2rxRqpIgZQVu3A0nzyFSyiWGoANfqbOuXAlNGJdLpEdX9NayjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUybiD8jsXx54mo6FGVw7Nwr2U2B0wGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQwNDU1MAoGCCqGSM49BAMCA0kAMEYCIQCJg09JPlcghezofCYF%2FH%2FFmz6TT%2FFsNYa3097UOoSvggIhAKU1PWWMdRQm5wtTLECzSGq1s%2B3EonNscjaoSics5Idm)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 17:17 UTC