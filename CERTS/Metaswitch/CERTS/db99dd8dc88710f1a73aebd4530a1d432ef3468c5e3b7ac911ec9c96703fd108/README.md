# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Five Area Telephone SHAKEN Cert 2071

Tested At: 18 Aug 25 20:05 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 292 day(s)\
Subject: CN=Five Area Telephone SHAKEN Cert 2071, O=Five Area Telephone Cooperative, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/738466bb39439e11398346bd5b0156a486803683

[View certificate details](https://x509.io/?cert=MIICbDCCAhOgAwIBAgIQLo6LXocGN4lmnpdNJX9kHzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDYwNzE2NTU0MFoXDTI2MDYwNjE2NTU0MFowZjELMAkGA1UEBhMCVVMxKDAmBgNVBAoMH0ZpdmUgQXJlYSBUZWxlcGhvbmUgQ29vcGVyYXRpdmUxLTArBgNVBAMMJEZpdmUgQXJlYSBUZWxlcGhvbmUgU0hBS0VOIENlcnQgMjA3MTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDvu9oqIdDJHkX0qk%2FVpDMS4Pz49jP4Aoq2Haa3uJqvoIMozSoSKtfI9Tz05zwa3YDUMqR50UqHojHAsEL0R82SjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDIwNzEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUQc6LhlueiQgcvR2tPOdub6fMk20wHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgQqTo83X0RWKFs8AF4AijSBTpOkBuoJpFpWnIsVWNgskCIC221BgoAYbrNRLo%2B4eX%2Fygfh4VraKpoYV7lS%2B4RXLQ2)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2071', but common name is 'Five Area Telephone SHAKEN Cert 2071' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 18 Aug 25 21:13 UTC