# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Summit Broadband SHAKEN Cert 7857

Tested At: 15 Nov 23 17:58 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 837 day(s)\
Subject: CN=Summit Broadband SHAKEN Cert 7857, O=Summit Broadband, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/47c6d4e4af62db63cec43e632b0c057700d9704d

[View certificate details](https://understandingwebpki.com/?cert=MIICWjCCAgGgAwIBAgIQJNPP%2B%2FySec22mLpAX6rGajAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDMwMjE3NTc1NloXDTI2MDMwMTE3NTc1NlowVDELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEFN1bW1pdCBCcm9hZGJhbmQxKjAoBgNVBAMMIVN1bW1pdCBCcm9hZGJhbmQgU0hBS0VOIENlcnQgNzg1NzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMyLqXdhK8EeOHfMuIyitpdGUh%2B8KPo5wBTTcckEOldIezCFViQh%2Fy88IO3snKfACJOvTy9HOUsHo1iB4stgC%2FujgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDc4NTcwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUq%2BI7o4sgMIvI7CGHFo8Shk12SOIwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgcbGpqWhu2AOAi8%2BY4vFqPUqlqIgpetzdflOorEZOgUkCIA2qU5uZJvo0Ly9GUKqZayLrVVqjvu5gXrLRFHrs5f7i)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7857' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 18:10 UTC