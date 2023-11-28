# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Alma Tel SHAKEN 0344

Tested At: 28 Nov 23 10:18 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 909 day(s)\
Subject: CN=Alma Tel SHAKEN 0344, O=Alma Tel, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/50fce21c2814c12146d9f3e9420b586cd6c12566

[View certificate details](https://understandingwebpki.com/?cert=MIICRjCCAeygAwIBAgIQVfk6cVoEn4TabjLMY2YGZzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUyNTIyMTc1M1oXDTI2MDUyNDIyMTc1M1owPzELMAkGA1UEBhMCVVMxETAPBgNVBAoMCEFsbWEgVGVsMR0wGwYDVQQDDBRBbG1hIFRlbCBTSEFLRU4gMDM0NDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPnVd1mkxdO5tWU1yNtZ1IC%2F%2Fpf2EFBMmCIjWuSmnMMcDdb3qj1eBY3xGSnEnP1N%2BqWuDATirbXqcp%2BaTFsfM9ujgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDAzNDQwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUcJzHq87ttMwiQv78GsgSs2mm04owHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAM5lcmAWKELxuFfH4fcvkV8rKy%2BCLaqayNyt5%2FsDyc8BAiAHjZww9sGrbNVsbypyTEJFCTSdR%2FIETXjgw4m5Mqm1Dw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 28 Nov 23 10:53 UTC