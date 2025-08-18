# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Blackfoot Communications SHAKEN Cert 2235

Tested At: 18 Aug 25 20:05 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 243 day(s)\
Subject: CN=Blackfoot Communications SHAKEN Cert 2235, O=Blackfoot Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/73aa28faf4546c63b3b20c530d38004b43bbecd4

[View certificate details](https://x509.io/?cert=MIICbDCCAhGgAwIBAgIQZhf9PrscAJGLvlYsldQNIzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDQxOTE0MTkzNFoXDTI2MDQxODE0MTkzNFowZDELMAkGA1UEBhMCVVMxITAfBgNVBAoMGEJsYWNrZm9vdCBDb21tdW5pY2F0aW9uczEyMDAGA1UEAwwpQmxhY2tmb290IENvbW11bmljYXRpb25zIFNIQUtFTiBDZXJ0IDIyMzUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS3TreOkJtIXQpz5XTLheeQ8tYiLUCpmf3y7rcxh7ZL1dM9yCasBGTGDg5JNz3bxGckJn%2FkFdx7bmSwH9bD2Ko2o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQyMjM1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFNqS%2BfOypQJemYXsY5hH%2F3bbEMGEMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQC%2BOBDFkTSOFJaP7288spgh6LPDwUua%2Fz%2FHsOBvNPkgTAIhAMYD6JFa%2FcWoZFLg%2FSs6O74PQlQoZRgdmzlT4kWts9FB)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2235', but common name is 'Blackfoot Communications SHAKEN Cert 2235' |


Generated: 18 Aug 25 21:13 UTC