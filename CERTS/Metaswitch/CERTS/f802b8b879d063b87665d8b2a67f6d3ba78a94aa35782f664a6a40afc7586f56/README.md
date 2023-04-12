# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Segra SHAKEN Cert 1784

Tested At: 12 Apr 23 21:57 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 401 day(s)\
Subject: CN=Segra SHAKEN Cert 1784, O=Segra, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/97ebbd6324f0d938d5e150ecb217c5ead6156cfa

[View certificate details](https://understandingwebpki.com/?cert=MIICRDCCAeugAwIBAgIQf9xpkL2Mokj4grfHR0sBhjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUxODEzMTQ1M1oXDTI0MDUxNzEzMTQ1M1owPjELMAkGA1UEBhMCVVMxDjAMBgNVBAoMBVNlZ3JhMR8wHQYDVQQDDBZTZWdyYSBTSEFLRU4gQ2VydCAxNzg0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEPJex5fFmeiPsNlgDKKODQnE%2FY7ZW4RhPxOiqJ3Ek7x4fPQplgkZRrzI2U%2BoUkHGXC10AIEcLDpR3XYtHLoS1paOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTc4NDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSEGd0m5DBxh9E5EpZXGvqcEW2rxTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiBz%2FfJZkOT3479tYHFk1ZmXa2L%2BnC1QRnavSskdjbFMRwIgD7h7tY4HwUKs%2FhvzXfNeCalbQzVHLS65Ws8TYNJ8qX0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Apr 23 22:02 UTC