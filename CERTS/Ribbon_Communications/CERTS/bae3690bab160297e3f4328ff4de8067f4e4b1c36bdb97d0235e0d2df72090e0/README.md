# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Fulton SHAKEN 0455

Tested At: 24 Nov 23 11:11 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 211 day(s)\
Subject: CN=Fulton SHAKEN 0455, OU=STI, O=Fulton Telephone Company, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/7crfAxXVRz/STI-0455-2023-5d0904544e3802776d08740599a8aa12

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApigAwIBAgIQXQkEVE44AndtCHQFmaiqEjAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNjIyMTkxMTI1WhcNMjQwNjIxMTkxMTI0WjBbMQswCQYDVQQGEwJVUzEhMB8GA1UECgwYRnVsdG9uIFRlbGVwaG9uZSBDb21wYW55MQwwCgYDVQQLDANTVEkxGzAZBgNVBAMMEkZ1bHRvbiBTSEFLRU4gMDQ1NTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNQJUDUrqmO%2FU8xEhK%2BA1%2BRiwdUQD7O29JHZAVx2rxRqpIgZQVu3A0nzyFSyiWGoANfqbOuXAlNGJdLpEdX9NayjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUybiD8jsXx54mo6FGVw7Nwr2U2B0wGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQwNDU1MAoGCCqGSM49BAMCA0kAMEYCIQCJg09JPlcghezofCYF%2FH%2FFmz6TT%2FFsNYa3097UOoSvggIhAKU1PWWMdRQm5wtTLECzSGq1s%2B3EonNscjaoSics5Idm)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_not_specified](../../ISSUES/e_atis_ext_not_specified/README.md) | error | ATIS1000080 | Certificate contains extensions that are not specified: 1.3.6.1.5.5.7.1.1 |


Generated: 24 Nov 23 11:17 UTC