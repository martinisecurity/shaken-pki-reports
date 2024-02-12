# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Appalachian Wireless SHAKEN Cert 6940

Tested At: 12 Feb 24 18:54 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 98 day(s)\
Subject: CN=Appalachian Wireless SHAKEN Cert 6940, O=Appalachian Wireless, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/590f3a17396e9e874d158eb60337f24c3ac4812b

[View certificate details](https://understandingwebpki.com/?cert=MIICZDCCAgmgAwIBAgIQREUxHsDn00q3Ok0S0nOG1zAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUyMDIyMDQxOFoXDTI0MDUxOTIyMDQxOFowXDELMAkGA1UEBhMCVVMxHTAbBgNVBAoMFEFwcGFsYWNoaWFuIFdpcmVsZXNzMS4wLAYDVQQDDCVBcHBhbGFjaGlhbiBXaXJlbGVzcyBTSEFLRU4gQ2VydCA2OTQwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzMw7rc1DbIoXd7fTVbxb66o54fuYqYlZl%2FYzzsB8b1Uc%2FAbU5DcY7bD1z74yIwl1rihnxNkg0h26hyb1lCNn3KOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENjk0MDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQ7wi8ppqCiTw3RgWX669i3lxaXGTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEA4txV1IfsMqIMdvS7l2U%2BO%2FoCNtcaSc%2BNjVi%2BWUQIwmsCIQC9TyJVBzoq4J7K6I8QU8D6pChQD8DE0IYViTLQcF%2B9mQ%3D%3D)

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


Generated: 12 Feb 24 19:26 UTC