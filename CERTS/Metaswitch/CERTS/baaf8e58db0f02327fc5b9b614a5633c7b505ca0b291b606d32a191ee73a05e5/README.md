# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Fidelity Communications SHAKEN Cert 1882

Tested At: 26 Aug 24 17:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: -259 day(s)\
Subject: CN=Fidelity Communications SHAKEN Cert 1882, O=Fidelity Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/54ea30a0105002d9a139f34e6297f0d2fdf88fca

[View certificate details](https://x509.io/?cert=MIICajCCAg%2BgAwIBAgIQfO82cw9bOmxFt4Pb5z7dgzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIwMTIxMTAwMTQyOFoXDTIzMTIxMTAwMTQyOFowYjELMAkGA1UEBhMCVVMxIDAeBgNVBAoMF0ZpZGVsaXR5IENvbW11bmljYXRpb25zMTEwLwYDVQQDDChGaWRlbGl0eSBDb21tdW5pY2F0aW9ucyBTSEFLRU4gQ2VydCAxODgyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDAcaweHKFoGZfAQq4GIiXEruNbqZP%2BYwIO4Wi2Wb5FtLEoRC%2BZpS7yMppDixT4CJ15j81UbXZOJifitq1JD%2FhKOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTg4MjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBROYU8TbcIDnUnqLnrqX3CzJZIJ2zAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEA8QdtLJfBGM3FkwNrIUk0DB3UocVC3h1jDPoRfQulqPMCIQDGDnFP8yHb7B5EjOHTt8Fy1p7vxgUqrIep990KLds6Hg%3D%3D)

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


Generated: 26 Aug 24 18:03 UTC