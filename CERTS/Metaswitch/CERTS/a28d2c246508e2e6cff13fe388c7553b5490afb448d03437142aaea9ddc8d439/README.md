# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Eatel SHAKEN Cert 8839

Tested At: 26 Aug 24 17:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 577 day(s)\
Subject: CN=Eatel SHAKEN Cert 8839, O=Eatel, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/3dc4ef07987f8fb0892c9b426513c57173ed8c3e

[View certificate details](https://x509.io/?cert=MIICRjCCAeugAwIBAgIQd3%2BvYhi6rJioD%2BSSupC66TAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDMyNjIyNTY0NVoXDTI2MDMyNTIyNTY0NVowPjELMAkGA1UEBhMCVVMxDjAMBgNVBAoMBUVhdGVsMR8wHQYDVQQDDBZFYXRlbCBTSEFLRU4gQ2VydCA4ODM5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEVoZ2SU5IsJAs75HlqTkK6i7%2BvOEpeKF9quHtDGs%2BgaOF78ZHf5za4FMADNGG7LekfJ3AzXFIv21zeJMiFKAQkKOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEODgzOTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBR%2BfoOsMMS%2BHY12prvX61Nc4eCMIzAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEAjcfkNur%2F7V7dU%2Brq2qNn0YasUBIaYRO%2BwP%2BT3tYV5TcCIQDBO3%2F5bfCwU%2B2Zl%2FOgnU59GJJPJcOB6qLtnsBwCkGBVA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 8839', but common name is 'Eatel SHAKEN Cert 8839' |


Generated: 26 Aug 24 18:03 UTC