# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Aeneas Communications SHAKEN Cert 2891

Tested At: 21 Nov 23 01:22 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 880 day(s)\
Subject: CN=Aeneas Communications SHAKEN Cert 2891, O=Aeneas Communications\\ , C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/25d9b4b15a9c7d0d8358bd9eb307a6ebbd8b7d00

[View certificate details](https://understandingwebpki.com/?cert=MIICZjCCAgygAwIBAgIQKytCe%2FkEPa14r5tjs12Z8zAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDQxOTE0MDgxNloXDTI2MDQxODE0MDgxNlowXzELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFkFlbmVhcyBDb21tdW5pY2F0aW9ucyAxLzAtBgNVBAMMJkFlbmVhcyBDb21tdW5pY2F0aW9ucyBTSEFLRU4gQ2VydCAyODkxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEuR4iEsvGDqzGQO8wHMWigfOLVwVX%2B0gpm%2BIgOsASj8M3XvUyKGJ5nIzUKVJAyNEc286wk1Gm8SZY9YIYOgbhqKOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMjg5MTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRgQMGhnX4XGjOLXwCtBD%2BKdNb0%2FTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNIADBFAiEAwdAouQQleHw%2BWjQQuiaq2iNX7t3X2Dpdb5NAm1GlKEwCIHKL7osNttHmJA7GMm5D3fwM%2F3XlBPbtJaK7vZpjZYYT)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2891' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 21 Nov 23 01:55 UTC