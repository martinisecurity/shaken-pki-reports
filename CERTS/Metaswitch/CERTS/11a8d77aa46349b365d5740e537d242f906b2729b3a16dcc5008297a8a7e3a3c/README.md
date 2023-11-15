# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Whidbey Telephone Company SHAKEN Cert 2452

Tested At: 15 Nov 23 17:58 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 935 day(s)\
Subject: CN=Whidbey Telephone Company SHAKEN Cert 2452, O=Whidbey Telephone Company, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/4b8c18a5a68450eb480b9727e86382fdaa4afa66

[View certificate details](https://understandingwebpki.com/?cert=MIICbTCCAhOgAwIBAgIQEYOQmammEJg%2BUMmpabTkKDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDYwODE2MDI0OFoXDTI2MDYwNzE2MDI0OFowZjELMAkGA1UEBhMCVVMxIjAgBgNVBAoMGVdoaWRiZXkgVGVsZXBob25lIENvbXBhbnkxMzAxBgNVBAMMKldoaWRiZXkgVGVsZXBob25lIENvbXBhbnkgU0hBS0VOIENlcnQgMjQ1MjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGauunR%2F1OnsVfDML9NjLok%2FgoMgo9Yd4ehuOpxMgwHEFRkSNuT35X0cVv22OFsJSoZ20WrltFrcg2VUXKvBJjSjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDI0NTIwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUpasuppRnkfQtst%2B8I54rQkB0jKowHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIgWmvdLzQrkFHOmxvpRvt%2B2TCFRGH1M8wd02hQApt0C8gCIQDaPB63Xr2fkqp9v1qDIiAWh6ifHwxfi1tV31mCFe7yIg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2452' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 18:10 UTC