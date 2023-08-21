# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Union Telephone Company SHAKEN Cert 2297

Tested At: 21 Aug 23 20:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 287 day(s)\
Subject: CN=Union Telephone Company SHAKEN Cert 2297, O=Union Telephone Company, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/e3e7946f99318b6966f972fa5d0b688753a050ca

[View certificate details](https://understandingwebpki.com/?cert=MIICaDCCAg%2BgAwIBAgIQWtcD97nVXHfi6iZKz68X7zAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYwNDE3Mjk0MloXDTI0MDYwMzE3Mjk0MlowYjELMAkGA1UEBhMCVVMxIDAeBgNVBAoMF1VuaW9uIFRlbGVwaG9uZSBDb21wYW55MTEwLwYDVQQDDChVbmlvbiBUZWxlcGhvbmUgQ29tcGFueSBTSEFLRU4gQ2VydCAyMjk3MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEv1OmMBpD132hNy5Q7nizlUAECBURCb%2F2JGZ9SDxJND9r7GtXQK1V4tJGUEI6NOUnJVvUohw33weK7%2FXGqpEVkaOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMjI5NzBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQF0UgAT2%2BDvHvRAiTFNCjaTcxRDTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiASB0iAetW3eIpPhBoTYxncX1s3%2BVhfZgFsk5kTTHQCmQIgZqFv1wG0YeExhw784Q3XKWZLD0cIXb2szKvERdISBPY%3D)

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


Generated: 21 Aug 23 20:18 UTC