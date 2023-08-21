# STIR/SHAKEN CA Ecosystem Compliance

## Certificate GVTC SHAKEN Cert 2083

Tested At: 21 Aug 23 20:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 916 day(s)\
Subject: CN=GVTC SHAKEN Cert 2083, O=GVTC, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/90b8c01a8cc0c06e82b9517d599d24c14dbb859b

[View certificate details](https://understandingwebpki.com/?cert=MIICQzCCAemgAwIBAgIQWvc%2BUVymE9%2FSCN5C7oycxDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDIyMjIwNDU1MFoXDTI2MDIyMTIwNDU1MFowPDELMAkGA1UEBhMCVVMxDTALBgNVBAoMBEdWVEMxHjAcBgNVBAMMFUdWVEMgU0hBS0VOIENlcnQgMjA4MzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABADtLAkxWmg4BW5w%2FbMBLJD9WkM8CzDhKhUOgu3sLuzg1BGx0XeXK8ljM1KAqvzQkgInk5j0fYz38INxvl%2BcG9ajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDIwODMwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUCX%2FAxX5zTjc8vzedB0ag14YoJOEwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIgBCYJKpLTexw2OFKdS9O%2F3n%2B0s9EwwFdNt9FnO753X94CIQDcZF0YCSNsAJemwCFqvgEs3rcOHI%2FEnplFCUe8zXbF%2Bg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2083' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


Generated: 21 Aug 23 20:18 UTC