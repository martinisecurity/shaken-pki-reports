# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Syringa Networks SHAKEN Cert 5869

Tested At: 28 Apr 23 02:17 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 459 day(s)\
Subject: CN=Syringa Networks SHAKEN Cert 5869, O=Syringa Networks, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/fe756139593d8af30f6d59a341bc41e4a6e21abb

[View certificate details](https://understandingwebpki.com/?cert=MIICXDCCAgGgAwIBAgIQco1lFBkU4%2FQbNYOm7P1rdjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDczMDE3MTY0NloXDTI0MDcyOTE3MTY0NlowVDELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEFN5cmluZ2EgTmV0d29ya3MxKjAoBgNVBAMMIVN5cmluZ2EgTmV0d29ya3MgU0hBS0VOIENlcnQgNTg2OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLbXr2NlOuGH%2FPtLogryTVaQo0M2guy65WfPLy6CGbVI1jrszYZcSYDKRZbIHQhty2FdaseGPP8oU%2FhtIfq6n%2BejgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDU4NjkwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUP1V5l7u2nVM1vCmXLFyuz8A5E6MwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAOHnsSStivj2j71rewA2d28bT4BlOj3BRZyuyjrcRsvsAiEA7jeYWI5RgITuLx4ie9lMGKFyxd6XBZ%2Fukkv2YYIIDRg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Apr 23 02:17 UTC