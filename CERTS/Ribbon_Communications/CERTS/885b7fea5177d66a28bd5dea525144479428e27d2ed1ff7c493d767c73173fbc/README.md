# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Nuwave Communications SHAKEN 620J

Tested At: 31 Oct 22 16:42 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 181 day(s)\
Subject: CN=Nuwave Communications SHAKEN 620J, OU=Nuwave Communications, O=Nuwave Communications, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-cr.rbbnidhub.com/Jx6yR-yMgz/620J202204-2c7d5c55a20834b031681dbd3e2eb9f0

View: [Click to view](https://understandingwebpki.com/?cert=MIIDETCCAragAwIBAgIQLH1cVaIINLAxaB29Pi658DAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwNDI5MTc0MzAwWhcNMjMwNDI5MTc0MjU5WjB5MQswCQYDVQQGEwJVUzEeMBwGA1UECgwVTnV3YXZlIENvbW11bmljYXRpb25zMR4wHAYDVQQLDBVOdXdhdmUgQ29tbXVuaWNhdGlvbnMxKjAoBgNVBAMMIU51d2F2ZSBDb21tdW5pY2F0aW9ucyBTSEFLRU4gNjIwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKQcN14DpTdbPoHaOy6cVJljOZPyQ%2BgyX9kIcftnhDG18ZcaWvrgXq21a6A4Y3vyMPMBhNFH4oB4qergCJqbgdmjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUqbMilZVs0XNrwwIFe%2BAsB%2Fa6tHowGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MjBKMAoGCCqGSM49BAMCA0kAMEYCIQCiqUKrea18gAnjqSzfXDL3r3Dv%2BE31IMIKw8ctF0nttwIhALOBc%2BpGfwh08qg%2Bf8qmg5wCPmMKr%2BKaWCLysLzlVer6)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_extension_unknown](../../ISSUES/e_sti_extension_unknown/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall not include extensions that are not specified |


Generated: 31/10/2022 at 16:43:22