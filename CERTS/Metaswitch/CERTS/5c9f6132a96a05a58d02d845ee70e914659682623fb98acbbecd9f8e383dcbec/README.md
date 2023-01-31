# STIR/SHAKEN CA Ecosystem Compliance

## Certificate CTS Telecom, Inc SHAKEN Cert 8331

Tested At: 31 Jan 23 21:46 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 593 day(s)\
Subject: CN=CTS Telecom\\, Inc SHAKEN Cert 8331, O=CTS Telecom\\, Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/6db90427e977d872d1d8aa0a64c3760b043f7edb

[View certificate details](https://understandingwebpki.com/?cert=MIICWzCCAgGgAwIBAgIQEdo%2BggoZj3D1TCHWBIBp%2FjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDkxNjEzMDk0OFoXDTI0MDkxNTEzMDk0OFowVDELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEENUUyBUZWxlY29tLCBJbmMxKjAoBgNVBAMMIUNUUyBUZWxlY29tLCBJbmMgU0hBS0VOIENlcnQgODMzMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCpjJKepeD3NlMFV3eTQUM8iK18cfJhx558BVVMNaZ75LQ0yxLzqGkG1iSGoX%2FfW%2BvmnsD%2FxQnfWimFUAa3Ufp2jgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDgzMzEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUB3E3M6QttsNcUGuLfQrgMFKKKwEwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAPSI83GR599xo5HjS%2F9kfIBvpuJ7HqQVUtRGDSFGX74xAiBuNuIL5mkynzOZVUVQIHV3msnnuLHEifp8%2F%2BQAa0zmkg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31 Jan 23 21:50 UTC