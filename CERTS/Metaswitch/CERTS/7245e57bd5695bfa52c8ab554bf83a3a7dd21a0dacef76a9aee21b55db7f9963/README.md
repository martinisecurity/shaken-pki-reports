# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Northland Networks SHAKEN Cert 7556

Tested At: 15 Nov 23 17:58 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 710 day(s)\
Subject: CN=Northland Networks SHAKEN Cert 7556, O=Northland Networks, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/37febf7ba183fa2f44fa353cf7f32aa6e1135a94

[View certificate details](https://understandingwebpki.com/?cert=MIICXzCCAgWgAwIBAgIQfAwbaiRI9dxl4AG7C2uHXzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMTAyNjE2MTIyM1oXDTI1MTAyNTE2MTIyM1owWDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEk5vcnRobGFuZCBOZXR3b3JrczEsMCoGA1UEAwwjTm9ydGhsYW5kIE5ldHdvcmtzIFNIQUtFTiBDZXJ0IDc1NTYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARwIh4Qlr7Pkv1M16fS41Cwc4T92f3K%2FuBAzYg3RPS7up0E0%2FjtUr7o4WVvGrGflh8Wd%2FLU%2BVrSCIAVDcs0XLJMo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3NTU2MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFM8WyqqY4lM0FblFImPn0TsRwjNvMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCbDHkDX39HTe1w%2FjygNqUOGHsiReHUQlwa%2B3I3yycVsgIgEB%2FhzDShkBSYaMjPFrJVlOJmI9Kjiz1%2Fg6xmFjUwj50%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7556' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 18:10 UTC