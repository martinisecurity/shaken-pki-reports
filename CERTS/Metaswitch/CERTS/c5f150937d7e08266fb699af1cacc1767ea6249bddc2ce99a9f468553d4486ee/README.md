# STIR/SHAKEN CA Ecosystem Compliance

## Certificate 702 Communications SHAKEN Cert 2737

Tested At: 28 Nov 23 15:59 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 154 day(s)\
Subject: CN=702 Communications SHAKEN Cert 2737, O=702 Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/dd6b11ece7734be2f2b5742bfee300774c703397

[View certificate details](https://understandingwebpki.com/?cert=MIICXzCCAgWgAwIBAgIQViLpXZyBporn3%2FDFYMTVNzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQzMDE2NTE0NloXDTI0MDQyOTE2NTE0NlowWDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEjcwMiBDb21tdW5pY2F0aW9uczEsMCoGA1UEAwwjNzAyIENvbW11bmljYXRpb25zIFNIQUtFTiBDZXJ0IDI3MzcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ6bRz3A7Z1ZjPt8cbiwm3WBe%2ByWRKQx1tPOcX7%2FdxauUPJzw7ewF8ZCaol%2Fu4gdU%2FruGAbe2PPP9pP1G7Xuahco4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQyNzM3MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFBoFOLZFbYAI1b97AMLYlKcHZvMxMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCICDaMJjB%2Birr3x5vzmmnQ%2B1XxcmXSmPj6QTlhVv3FMnCAiEA3ZS54YjUF4qZZE8Xab1np8QJ67teyp%2FC0SeT6ekA7Dw%3D)

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


Generated: 28 Nov 23 16:15 UTC