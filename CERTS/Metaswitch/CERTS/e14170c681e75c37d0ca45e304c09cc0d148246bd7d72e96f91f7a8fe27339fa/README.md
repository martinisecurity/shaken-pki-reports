# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Appalachian Wireless SHAKEN Cert 6940

Tested At: 31 Oct 22 18:33 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 567 day(s)\
Subject: CN=Appalachian Wireless SHAKEN Cert 6940, O=Appalachian Wireless, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/590f3a17396e9e874d158eb60337f24c3ac4812b

View: [Click to view](https://understandingwebpki.com/?cert=MIICZDCCAgmgAwIBAgIQREUxHsDn00q3Ok0S0nOG1zAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUyMDIyMDQxOFoXDTI0MDUxOTIyMDQxOFowXDELMAkGA1UEBhMCVVMxHTAbBgNVBAoMFEFwcGFsYWNoaWFuIFdpcmVsZXNzMS4wLAYDVQQDDCVBcHBhbGFjaGlhbiBXaXJlbGVzcyBTSEFLRU4gQ2VydCA2OTQwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzMw7rc1DbIoXd7fTVbxb66o54fuYqYlZl%2FYzzsB8b1Uc%2FAbU5DcY7bD1z74yIwl1rihnxNkg0h26hyb1lCNn3KOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENjk0MDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQ7wi8ppqCiTw3RgWX669i3lxaXGTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEA4txV1IfsMqIMdvS7l2U%2BO%2FoCNtcaSc%2BNjVi%2BWUQIwmsCIQC9TyJVBzoq4J7K6I8QU8D6pChQD8DE0IYViTLQcF%2B9mQ%3D%3D)


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


Generated: 31/10/2022 at 18:34:12