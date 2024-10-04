# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Consolidated Telcom ND SHAKEN Cert 7008

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 526 day(s)\
Subject: CN=Consolidated Telcom ND SHAKEN Cert 7008, O=Consolidated Telcom ND, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/1e9efaedff05b3fc118ca5c6c413a10ce0a84b98

[View certificate details](https://x509.io/?cert=MIICaDCCAg2gAwIBAgIQA4Nbzvf0PP8Gqq%2BBFXwlJDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDMxNDIzMzcwMVoXDTI2MDMxMzIzMzcwMVowYDELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFkNvbnNvbGlkYXRlZCBUZWxjb20gTkQxMDAuBgNVBAMMJ0NvbnNvbGlkYXRlZCBUZWxjb20gTkQgU0hBS0VOIENlcnQgNzAwODBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGuJWHPR%2FmcPDQfenvYNAKakHcIWY5%2F%2FnsCGO6sUWJ%2F27hiPwipGARMDF0LVZ8XDrCbc5q0Bk%2BxnyzhXthJPpqOjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDcwMDgwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUFTuWZLHuUFLL%2FuuMGQyoKLozX1wwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAMW1MAJrkCOWcX3KQ8q8vJIRqj38SwxSuE3oq0XJaExEAiEA3lcAgvx63M0hhYyFz9ze4PrqSKMYz2PQksoNe3hxe2w%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7008', but common name is 'Consolidated Telcom ND SHAKEN Cert 7008' |


Generated: 04 Oct 24 16:29 UTC