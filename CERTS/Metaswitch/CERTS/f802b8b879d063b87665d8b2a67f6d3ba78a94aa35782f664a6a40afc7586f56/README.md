# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Segra SHAKEN Cert 1784

Tested At: 01 Nov 22 07:32 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 564 day(s)\
Subject: CN=Segra SHAKEN Cert 1784, O=Segra, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/97ebbd6324f0d938d5e150ecb217c5ead6156cfa

View: [Click to view](https://understandingwebpki.com/?cert=MIICRDCCAeugAwIBAgIQf9xpkL2Mokj4grfHR0sBhjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUxODEzMTQ1M1oXDTI0MDUxNzEzMTQ1M1owPjELMAkGA1UEBhMCVVMxDjAMBgNVBAoMBVNlZ3JhMR8wHQYDVQQDDBZTZWdyYSBTSEFLRU4gQ2VydCAxNzg0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEPJex5fFmeiPsNlgDKKODQnE%2FY7ZW4RhPxOiqJ3Ek7x4fPQplgkZRrzI2U%2BoUkHGXC10AIEcLDpR3XYtHLoS1paOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTc4NDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSEGd0m5DBxh9E5EpZXGvqcEW2rxTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiBz%2FfJZkOT3479tYHFk1ZmXa2L%2BnC1QRnavSskdjbFMRwIgD7h7tY4HwUKs%2FhvzXfNeCalbQzVHLS65Ws8TYNJ8qX0%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 07:33:04