# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Fastwyre Broadband SHAKEN Cert 0425

Tested At: 15 Nov 23 16:09 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 908 day(s)\
Subject: CN=Fastwyre Broadband SHAKEN Cert 0425, O=Fastwyre Broadband, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/4f2a4a764bced7e27a5c80067dd0410ddd1b9e02

[View certificate details](https://understandingwebpki.com/?cert=MIICXzCCAgWgAwIBAgIQY4Y4Pj6sI%2B3%2FJV1EwrR4STAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUxMjExMTE0NloXDTI2MDUxMTExMTE0NlowWDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkZhc3R3eXJlIEJyb2FkYmFuZDEsMCoGA1UEAwwjRmFzdHd5cmUgQnJvYWRiYW5kIFNIQUtFTiBDZXJ0IDA0MjUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASGItNXs3hggAopaSjYwMT1I2GB8yWE6%2B2LNIzmY7zKchi6PMQynUMJv5HQXmNRS9cPF49YhXR8K%2Bm62bQQqC2Po4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwNDI1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFOXwYNq1nH17%2BHN%2F7rrsDYW46pl6MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCbk76EAUr%2B4%2FWtfSbzQNeX8c2Cyw9MkoyI0RfQ2AWCaQIgOEPnx6erv7waud6SDSJABwJhPFMpp%2BiKlqlQtg296CY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0425' |


Generated: 15 Nov 23 17:17 UTC