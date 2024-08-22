# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Segra SHAKEN Cert 1784

Tested At: 22 Aug 24 15:19 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 914 day(s)\
Subject: CN=Segra SHAKEN Cert 1784, O=Segra, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/cc9fae904b0b72958a4bbbcb4c068c738dd67f1c

[View certificate details](https://x509.io/?cert=MIICRjCCAeugAwIBAgIQID5sAE2rGQUB1ezAGuU1tjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDIyMzExMTI1NloXDTI3MDIyMjExMTI1NlowPjELMAkGA1UEBhMCVVMxDjAMBgNVBAoMBVNlZ3JhMR8wHQYDVQQDDBZTZWdyYSBTSEFLRU4gQ2VydCAxNzg0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAERGorjbEKZ5URsTXtb9vib8eID4eBYM5E9KaTQ2sOw8yQNM4W5XAhVc65t0KvlueOMgbs1d395jfgsOjwHxqtlqOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTc4NDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBQpoXdjEDOb9G7S7ZkvG4DOCsUCJDAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEAqbLlSsCw6HUBUMrqvPZarZDugB3%2FQm7ONjXYfA6uGW8CIQC08oFOmfdTIF9gXc9z%2BBNf0hlAwLr5yw3o3mIkK456zg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1784', but common name is 'Segra SHAKEN Cert 1784' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 22 Aug 24 16:06 UTC