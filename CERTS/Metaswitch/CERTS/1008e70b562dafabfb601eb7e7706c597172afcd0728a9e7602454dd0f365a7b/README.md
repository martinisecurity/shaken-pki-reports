# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Alma Tel SHAKEN 0344

Tested At: 21 Aug 23 20:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 1008 day(s)\
Subject: CN=Alma Tel SHAKEN 0344, O=Alma Tel, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/50fce21c2814c12146d9f3e9420b586cd6c12566

[View certificate details](https://understandingwebpki.com/?cert=MIICRjCCAeygAwIBAgIQVfk6cVoEn4TabjLMY2YGZzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUyNTIyMTc1M1oXDTI2MDUyNDIyMTc1M1owPzELMAkGA1UEBhMCVVMxETAPBgNVBAoMCEFsbWEgVGVsMR0wGwYDVQQDDBRBbG1hIFRlbCBTSEFLRU4gMDM0NDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPnVd1mkxdO5tWU1yNtZ1IC%2F%2Fpf2EFBMmCIjWuSmnMMcDdb3qj1eBY3xGSnEnP1N%2BqWuDATirbXqcp%2BaTFsfM9ujgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDAzNDQwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUcJzHq87ttMwiQv78GsgSs2mm04owHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAM5lcmAWKELxuFfH4fcvkV8rKy%2BCLaqayNyt5%2FsDyc8BAiAHjZww9sGrbNVsbypyTEJFCTSdR%2FIETXjgw4m5Mqm1Dw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 21 Aug 23 20:18 UTC