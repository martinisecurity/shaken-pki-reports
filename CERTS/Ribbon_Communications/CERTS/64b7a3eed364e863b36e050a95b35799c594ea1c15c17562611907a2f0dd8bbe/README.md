# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Netfortris SHAKEN 8886

Tested At: 11 Jan 23 21:52 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 150 day(s)\
Subject: CN=Netfortris SHAKEN 8886, OU=VOIP, O=Netfortris, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/t0CiOIjnRz/NFJune102022-26092b2abd3c6bbfc676d98623fc1b25

[View certificate details](https://understandingwebpki.com/?cert=MIIC6DCCAo%2BgAwIBAgIQJgkrKr08a7%2FGdtmGI%2FwbJTAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwNjEwMTkwMDA5WhcNMjMwNjEwMTkwMDA4WjBSMQswCQYDVQQGEwJVUzETMBEGA1UECgwKTmV0Zm9ydHJpczENMAsGA1UECwwEVk9JUDEfMB0GA1UEAwwWTmV0Zm9ydHJpcyBTSEFLRU4gODg4NjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHKjVX0y5iWU%2BPWcVma0eo8bOVa6szzU%2BEyknh5%2BvU%2FrxREY8K9gnIQZwzAQC%2Fobbdb9g467rPWSJnwYFydtfCCjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUaaPiQZSfxkkc%2F0jrtzilcXSuLVUwGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQ4ODg2MAoGCCqGSM49BAMCA0cAMEQCIAiL%2FtMGhx5ilLmO11px22GSd%2FRUvCx90DyXF7PHiMZzAiBR436HM6vmoXCoSf%2FSUpihGhcjv0LwC4Gaz4xQH7d4eQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 11 Jan 23 21:59 UTC