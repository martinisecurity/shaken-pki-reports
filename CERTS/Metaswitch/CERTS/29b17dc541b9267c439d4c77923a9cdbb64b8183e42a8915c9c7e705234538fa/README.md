# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Nex-Tech Wireless SHAKEN Cert 122D

Tested At: 18 Aug 25 20:05 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 579 day(s)\
Subject: CN=Nex-Tech Wireless SHAKEN Cert 122D, O=Nex-Tech Wireless, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/80d8452d00cdf3a5975cb2301a8dcaab3715ae57

[View certificate details](https://x509.io/?cert=MIICXjCCAgOgAwIBAgIQIITW9yjZC3jRidGN0Anm6DAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDMyMDE5NDYyMFoXDTI3MDMyMDE5NDYyMFowVjELMAkGA1UEBhMCVVMxGjAYBgNVBAoMEU5leC1UZWNoIFdpcmVsZXNzMSswKQYDVQQDDCJOZXgtVGVjaCBXaXJlbGVzcyBTSEFLRU4gQ2VydCAxMjJEMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE1Dv%2B4fvieT%2B%2F9L1NHpyQimD1bkaIN4KziDnXITebcadU6Mo72cjpA1%2BNw2IBbeCZKGGfWOCjZlIFr5UeeURFDqOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTIyRDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRuJyC5Vtm2RhKF%2FdRKL1iSQiIx1DAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEAuykFZ8lh7tAXpi8d5XDl%2ByD55byXceLg0ggUhdAPEI8CIQCpfMSlNVcX19z7k0QjETizdng61kL6yKdIwsX5n4NC7g%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 122D', but common name is 'Nex-Tech Wireless SHAKEN Cert 122D' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 18 Aug 25 21:13 UTC