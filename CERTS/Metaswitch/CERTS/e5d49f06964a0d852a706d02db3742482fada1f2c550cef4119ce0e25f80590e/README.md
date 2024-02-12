# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Union Telephone Company SHAKEN Cert 2297

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 113 day(s)\
Subject: CN=Union Telephone Company SHAKEN Cert 2297, O=Union Telephone Company, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/79b04ced7278d55340ae248b065543b91d3c6ed3

[View certificate details](https://understandingwebpki.com/?cert=MIICaDCCAg%2BgAwIBAgIQWtcD97nVXHfi6iZKz68X7zAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYwNDE3Mjk0MloXDTI0MDYwMzE3Mjk0MlowYjELMAkGA1UEBhMCVVMxIDAeBgNVBAoMF1VuaW9uIFRlbGVwaG9uZSBDb21wYW55MTEwLwYDVQQDDChVbmlvbiBUZWxlcGhvbmUgQ29tcGFueSBTSEFLRU4gQ2VydCAyMjk3MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEv1OmMBpD132hNy5Q7nizlUAECBURCb%2F2JGZ9SDxJND9r7GtXQK1V4tJGUEI6NOUnJVvUohw33weK7%2FXGqpEVkaOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMjI5NzBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQF0UgAT2%2BDvHvRAiTFNCjaTcxRDTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiASB0iAetW3eIpPhBoTYxncX1s3%2BVhfZgFsk5kTTHQCmQIgZqFv1wG0YeExhw784Q3XKWZLD0cIXb2szKvERdISBPY%3D)

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


Generated: 12 Feb 24 17:02 UTC