# STIR/SHAKEN CA Ecosystem Compliance

## Certificate CTS Telecom, Inc SHAKEN Cert 8331

Tested At: 27 Nov 23 23:10 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 293 day(s)\
Subject: CN=CTS Telecom\\, Inc SHAKEN Cert 8331, O=CTS Telecom\\, Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/6db90427e977d872d1d8aa0a64c3760b043f7edb

[View certificate details](https://understandingwebpki.com/?cert=MIICWzCCAgGgAwIBAgIQEdo%2BggoZj3D1TCHWBIBp%2FjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDkxNjEzMDk0OFoXDTI0MDkxNTEzMDk0OFowVDELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEENUUyBUZWxlY29tLCBJbmMxKjAoBgNVBAMMIUNUUyBUZWxlY29tLCBJbmMgU0hBS0VOIENlcnQgODMzMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCpjJKepeD3NlMFV3eTQUM8iK18cfJhx558BVVMNaZ75LQ0yxLzqGkG1iSGoX%2FfW%2BvmnsD%2FxQnfWimFUAa3Ufp2jgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDgzMzEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUB3E3M6QttsNcUGuLfQrgMFKKKwEwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAPSI83GR599xo5HjS%2F9kfIBvpuJ7HqQVUtRGDSFGX74xAiBuNuIL5mkynzOZVUVQIHV3msnnuLHEifp8%2F%2BQAa0zmkg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 23:28 UTC