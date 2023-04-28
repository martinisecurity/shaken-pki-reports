# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Nuwave Communications SHAKEN 620J

Tested At: 28 Apr 23 02:11 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 341 day(s)\
Subject: CN=Nuwave Communications SHAKEN 620J, OU=Nuwave Communications, O=Nuwave Communications, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/Jx6yR-yMgz/620J202304-401f156c7342aa701f6e8bb3d26c0ecf

[View certificate details](https://understandingwebpki.com/?cert=MIIDDzCCAragAwIBAgIQQB8VbHNCqnAfbouz0mwOzzAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNDAzMjIwMjQwWhcNMjQwNDAyMjIwMjM5WjB5MQswCQYDVQQGEwJVUzEeMBwGA1UECgwVTnV3YXZlIENvbW11bmljYXRpb25zMR4wHAYDVQQLDBVOdXdhdmUgQ29tbXVuaWNhdGlvbnMxKjAoBgNVBAMMIU51d2F2ZSBDb21tdW5pY2F0aW9ucyBTSEFLRU4gNjIwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABBYhyJ%2FILpAtFGTQ0HjnJ13D87YysIZnU8wTIYGQaitqb41n3LLMP8RJ9HtOCVxt0V4oETbY0qTgpUeHOoXb7BKjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQU43QNC1vmWsST5%2FynSaN%2BPNdFIdMwGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MjBKMAoGCCqGSM49BAMCA0cAMEQCIDNBfnhM1d1h3SPNcpPccfE2BzrlOXjqdOjxlhv1ngabAiBbWhAqRebHNmqnHS2ubTUf%2FlivVdTb2iD2FalV71xk2g%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 28 Apr 23 02:17 UTC