# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Nuwave Communications SHAKEN 620J

Tested At: 24 Nov 23 11:11 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -208 day(s)\
Subject: CN=Nuwave Communications SHAKEN 620J, OU=Nuwave Communications, O=Nuwave Communications, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-cr.rbbnidhub.com/Jx6yR-yMgz/620J202204-2c7d5c55a20834b031681dbd3e2eb9f0

[View certificate details](https://understandingwebpki.com/?cert=MIIDETCCAragAwIBAgIQLH1cVaIINLAxaB29Pi658DAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwNDI5MTc0MzAwWhcNMjMwNDI5MTc0MjU5WjB5MQswCQYDVQQGEwJVUzEeMBwGA1UECgwVTnV3YXZlIENvbW11bmljYXRpb25zMR4wHAYDVQQLDBVOdXdhdmUgQ29tbXVuaWNhdGlvbnMxKjAoBgNVBAMMIU51d2F2ZSBDb21tdW5pY2F0aW9ucyBTSEFLRU4gNjIwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKQcN14DpTdbPoHaOy6cVJljOZPyQ%2BgyX9kIcftnhDG18ZcaWvrgXq21a6A4Y3vyMPMBhNFH4oB4qergCJqbgdmjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUqbMilZVs0XNrwwIFe%2BAsB%2Fa6tHowGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MjBKMAoGCCqGSM49BAMCA0kAMEYCIQCiqUKrea18gAnjqSzfXDL3r3Dv%2BE31IMIKw8ctF0nttwIhALOBc%2BpGfwh08qg%2Bf8qmg5wCPmMKr%2BKaWCLysLzlVer6)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 24 Nov 23 11:17 UTC