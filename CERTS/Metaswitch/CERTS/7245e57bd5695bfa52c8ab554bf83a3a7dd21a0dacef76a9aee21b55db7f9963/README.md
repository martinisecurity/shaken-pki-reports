# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Northland Networks SHAKEN Cert 7556

Tested At: 22 Aug 24 15:16 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 430 day(s)\
Subject: CN=Northland Networks SHAKEN Cert 7556, O=Northland Networks, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/37febf7ba183fa2f44fa353cf7f32aa6e1135a94

[View certificate details](https://x509.io/?cert=MIICXzCCAgWgAwIBAgIQfAwbaiRI9dxl4AG7C2uHXzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMTAyNjE2MTIyM1oXDTI1MTAyNTE2MTIyM1owWDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEk5vcnRobGFuZCBOZXR3b3JrczEsMCoGA1UEAwwjTm9ydGhsYW5kIE5ldHdvcmtzIFNIQUtFTiBDZXJ0IDc1NTYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARwIh4Qlr7Pkv1M16fS41Cwc4T92f3K%2FuBAzYg3RPS7up0E0%2FjtUr7o4WVvGrGflh8Wd%2FLU%2BVrSCIAVDcs0XLJMo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3NTU2MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFM8WyqqY4lM0FblFImPn0TsRwjNvMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCbDHkDX39HTe1w%2FjygNqUOGHsiReHUQlwa%2B3I3yycVsgIgEB%2FhzDShkBSYaMjPFrJVlOJmI9Kjiz1%2Fg6xmFjUwj50%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7556', but common name is 'Northland Networks SHAKEN Cert 7556' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 15:44 UTC