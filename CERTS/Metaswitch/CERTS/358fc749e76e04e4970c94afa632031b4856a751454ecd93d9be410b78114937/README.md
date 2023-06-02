# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Race Communications SHAKEN Cert 973E

Tested At: 02 Jun 23 01:11 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 680 day(s)\
Subject: CN=Race Communications SHAKEN Cert 973E, O=Race Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/7aef3caf659e78d58aaf4a541abdaf284692a168

[View certificate details](https://understandingwebpki.com/?cert=MIICYDCCAgegAwIBAgIQLKGBIrXH5b4XbQnxl6VnYTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDQxMjA5NTMzNFoXDTI1MDQxMTA5NTMzNFowWjELMAkGA1UEBhMCVVMxHDAaBgNVBAoME1JhY2UgQ29tbXVuaWNhdGlvbnMxLTArBgNVBAMMJFJhY2UgQ29tbXVuaWNhdGlvbnMgU0hBS0VOIENlcnQgOTczRTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOM%2Bd3twsjYOuVa30Uj0DCp6F4V0CToPQo3G9yW5TYM25%2BIcRFm9QKkI5cR0T%2F23TvfFcYMsiQ4tpwIkm7reAmKjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDk3M0UwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU8NwD57891kpF9xPUGVO6UrB5wFAwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgWcaZUxs39KoxxhII5UXX7%2FGEs54awI22BZkEKHGPNtICIDhiiRUyIZzOWX5XqqVsU9wbyk37L837%2Ff0HofkDKSAt)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 973E' |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |


Generated: 02 Jun 23 01:12 UTC