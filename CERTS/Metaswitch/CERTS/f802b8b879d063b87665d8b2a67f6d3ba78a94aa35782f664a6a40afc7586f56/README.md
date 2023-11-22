# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Segra SHAKEN Cert 1784

Tested At: 22 Nov 23 03:16 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 178 day(s)\
Subject: CN=Segra SHAKEN Cert 1784, O=Segra, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/86728e9104b39e125eb3c8a4f8a8224325991d00

[View certificate details](https://understandingwebpki.com/?cert=MIICRDCCAeugAwIBAgIQf9xpkL2Mokj4grfHR0sBhjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUxODEzMTQ1M1oXDTI0MDUxNzEzMTQ1M1owPjELMAkGA1UEBhMCVVMxDjAMBgNVBAoMBVNlZ3JhMR8wHQYDVQQDDBZTZWdyYSBTSEFLRU4gQ2VydCAxNzg0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEPJex5fFmeiPsNlgDKKODQnE%2FY7ZW4RhPxOiqJ3Ek7x4fPQplgkZRrzI2U%2BoUkHGXC10AIEcLDpR3XYtHLoS1paOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTc4NDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSEGd0m5DBxh9E5EpZXGvqcEW2rxTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiBz%2FfJZkOT3479tYHFk1ZmXa2L%2BnC1QRnavSskdjbFMRwIgD7h7tY4HwUKs%2FhvzXfNeCalbQzVHLS65Ws8TYNJ8qX0%3D)

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


Generated: 22 Nov 23 03:38 UTC