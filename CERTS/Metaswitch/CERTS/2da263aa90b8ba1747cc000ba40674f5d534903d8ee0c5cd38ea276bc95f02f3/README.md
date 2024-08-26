# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Fastwyre Broadband SHAKEN Cert 0425

Tested At: 26 Aug 24 17:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 623 day(s)\
Subject: CN=Fastwyre Broadband SHAKEN Cert 0425, O=Fastwyre Broadband, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/4f2a4a764bced7e27a5c80067dd0410ddd1b9e02

[View certificate details](https://x509.io/?cert=MIICXzCCAgWgAwIBAgIQY4Y4Pj6sI%2B3%2FJV1EwrR4STAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUxMjExMTE0NloXDTI2MDUxMTExMTE0NlowWDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkZhc3R3eXJlIEJyb2FkYmFuZDEsMCoGA1UEAwwjRmFzdHd5cmUgQnJvYWRiYW5kIFNIQUtFTiBDZXJ0IDA0MjUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASGItNXs3hggAopaSjYwMT1I2GB8yWE6%2B2LNIzmY7zKchi6PMQynUMJv5HQXmNRS9cPF49YhXR8K%2Bm62bQQqC2Po4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwNDI1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFOXwYNq1nH17%2BHN%2F7rrsDYW46pl6MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCbk76EAUr%2B4%2FWtfSbzQNeX8c2Cyw9MkoyI0RfQ2AWCaQIgOEPnx6erv7waud6SDSJABwJhPFMpp%2BiKlqlQtg296CY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0425', but common name is 'Fastwyre Broadband SHAKEN Cert 0425' |


Generated: 26 Aug 24 18:03 UTC