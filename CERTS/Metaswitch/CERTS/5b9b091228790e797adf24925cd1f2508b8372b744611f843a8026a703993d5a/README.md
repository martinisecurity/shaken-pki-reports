# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Syringa Networks SHAKEN Cert 5869

Tested At: 21 Nov 23 16:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 252 day(s)\
Subject: CN=Syringa Networks SHAKEN Cert 5869, O=Syringa Networks, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/f9a476014f036865318edcafd2156667825b4092

[View certificate details](https://understandingwebpki.com/?cert=MIICXDCCAgGgAwIBAgIQco1lFBkU4%2FQbNYOm7P1rdjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDczMDE3MTY0NloXDTI0MDcyOTE3MTY0NlowVDELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEFN5cmluZ2EgTmV0d29ya3MxKjAoBgNVBAMMIVN5cmluZ2EgTmV0d29ya3MgU0hBS0VOIENlcnQgNTg2OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLbXr2NlOuGH%2FPtLogryTVaQo0M2guy65WfPLy6CGbVI1jrszYZcSYDKRZbIHQhty2FdaseGPP8oU%2FhtIfq6n%2BejgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDU4NjkwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUP1V5l7u2nVM1vCmXLFyuz8A5E6MwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAOHnsSStivj2j71rewA2d28bT4BlOj3BRZyuyjrcRsvsAiEA7jeYWI5RgITuLx4ie9lMGKFyxd6XBZ%2Fukkv2YYIIDRg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC