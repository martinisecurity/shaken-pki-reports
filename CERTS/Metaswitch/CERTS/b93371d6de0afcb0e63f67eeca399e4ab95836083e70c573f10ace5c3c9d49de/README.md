# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Aeneas Communications SHAKEN Cert 2891

Tested At: 15 Nov 23 17:58 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 885 day(s)\
Subject: CN=Aeneas Communications SHAKEN Cert 2891, O=Aeneas Communications\\ , C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/25d9b4b15a9c7d0d8358bd9eb307a6ebbd8b7d00

[View certificate details](https://understandingwebpki.com/?cert=MIICZjCCAgygAwIBAgIQKytCe%2FkEPa14r5tjs12Z8zAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDQxOTE0MDgxNloXDTI2MDQxODE0MDgxNlowXzELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFkFlbmVhcyBDb21tdW5pY2F0aW9ucyAxLzAtBgNVBAMMJkFlbmVhcyBDb21tdW5pY2F0aW9ucyBTSEFLRU4gQ2VydCAyODkxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEuR4iEsvGDqzGQO8wHMWigfOLVwVX%2B0gpm%2BIgOsASj8M3XvUyKGJ5nIzUKVJAyNEc286wk1Gm8SZY9YIYOgbhqKOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMjg5MTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRgQMGhnX4XGjOLXwCtBD%2BKdNb0%2FTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNIADBFAiEAwdAouQQleHw%2BWjQQuiaq2iNX7t3X2Dpdb5NAm1GlKEwCIHKL7osNttHmJA7GMm5D3fwM%2F3XlBPbtJaK7vZpjZYYT)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2891' |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


Generated: 15 Nov 23 18:10 UTC