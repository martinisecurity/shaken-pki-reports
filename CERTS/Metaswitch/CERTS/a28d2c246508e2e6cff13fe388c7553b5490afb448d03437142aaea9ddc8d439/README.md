# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Eatel SHAKEN Cert 8839

Tested At: 21 Aug 23 20:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 948 day(s)\
Subject: CN=Eatel SHAKEN Cert 8839, O=Eatel, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/711741801195f6f10a57f3bb7369c2024e8e4933

[View certificate details](https://understandingwebpki.com/?cert=MIICRjCCAeugAwIBAgIQd3%2BvYhi6rJioD%2BSSupC66TAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDMyNjIyNTY0NVoXDTI2MDMyNTIyNTY0NVowPjELMAkGA1UEBhMCVVMxDjAMBgNVBAoMBUVhdGVsMR8wHQYDVQQDDBZFYXRlbCBTSEFLRU4gQ2VydCA4ODM5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEVoZ2SU5IsJAs75HlqTkK6i7%2BvOEpeKF9quHtDGs%2BgaOF78ZHf5za4FMADNGG7LekfJ3AzXFIv21zeJMiFKAQkKOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEODgzOTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBR%2BfoOsMMS%2BHY12prvX61Nc4eCMIzAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEAjcfkNur%2F7V7dU%2Brq2qNn0YasUBIaYRO%2BwP%2BT3tYV5TcCIQDBO3%2F5bfCwU%2B2Zl%2FOgnU59GJJPJcOB6qLtnsBwCkGBVA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 8839' |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


Generated: 21 Aug 23 20:18 UTC