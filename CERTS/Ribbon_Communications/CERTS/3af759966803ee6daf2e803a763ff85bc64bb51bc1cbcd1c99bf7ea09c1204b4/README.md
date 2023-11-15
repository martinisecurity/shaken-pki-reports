# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SMU SHAKEN 3390

Tested At: 15 Nov 23 16:37 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 230 day(s)\
Subject: CN=SMU SHAKEN 3390, OU=STI, O=Spencer Municipal Utilities, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/o_W93Cgngz/STI-202307-3390-7148553e5c7080fef0ac81ea7235258e

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApigAwIBAgIQcUhVPlxwgP7wrIHqcjUljjAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNzAzMDUyOTI4WhcNMjQwNzAyMDUyOTI3WjBbMQswCQYDVQQGEwJVUzEkMCIGA1UECgwbU3BlbmNlciBNdW5pY2lwYWwgVXRpbGl0aWVzMQwwCgYDVQQLDANTVEkxGDAWBgNVBAMMD1NNVSBTSEFLRU4gMzM5MDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA4L9Izm0IeMRB6p0E7U4ULlnBhvKAiNcLusZeeC5oS54NW6H1re114zcSkU1%2BdJW4k4s6XRUqN43pU0TwjUU3%2BjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUiNUIlA6qSsqUVDdN45sQWKF5JagwGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQzMzkwMAoGCCqGSM49BAMCA0kAMEYCIQDqfW3nMdFP2PmQM5ZTrj%2FnRsFUWPkiHy7DnLSFRwYumAIhALyU8gfyRErhWTtvl%2BAgvF2Ybak768H%2B2GdO6VKQUvm1)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 15 Nov 23 17:17 UTC