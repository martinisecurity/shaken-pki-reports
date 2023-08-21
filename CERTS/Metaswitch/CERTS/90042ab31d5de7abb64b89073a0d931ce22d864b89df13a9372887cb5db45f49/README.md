# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Northeast Communications of Wisconsin SHAKEN Cert 6692

Tested At: 21 Aug 23 20:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 202 day(s)\
Subject: CN=Northeast Communications of Wisconsin SHAKEN Cert 6692, O=Northeast Communications of Wisconsin, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/1a7ab760b70a56725786e626f157ba4512f909ec

[View certificate details](https://understandingwebpki.com/?cert=MIIChjCCAiugAwIBAgIQNWXGgeC2lKjQl3tQ8dSa6jAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxMDIwNTAyOVoXDTI0MDMwOTIwNTAyOVowfjELMAkGA1UEBhMCVVMxLjAsBgNVBAoMJU5vcnRoZWFzdCBDb21tdW5pY2F0aW9ucyBvZiBXaXNjb25zaW4xPzA9BgNVBAMMNk5vcnRoZWFzdCBDb21tdW5pY2F0aW9ucyBvZiBXaXNjb25zaW4gU0hBS0VOIENlcnQgNjY5MjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHfWdQKcMn7l%2FfoGVxmDuiwVQNNSr5H1a3THZ1C3bblV%2F%2BSf2FxPKASGNFXL61oGtFiY%2FTdz4IWgevMV6e3BZd6jgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDY2OTIwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUGMzXpYwh3CqUQ2%2FlB75ulad5MRAwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAPjcNrT76chSB%2FORIjnwwAjcfqEUOISvv83t%2FcbV68T6AiEA%2F6WC0sTRarIchFIrvF6cPxBWkucv%2FGQiBOQEC3KQYVQ%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Aug 23 20:18 UTC