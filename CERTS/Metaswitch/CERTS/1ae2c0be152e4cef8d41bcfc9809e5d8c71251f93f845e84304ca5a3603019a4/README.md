# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Blackfoot Communications SHAKEN Cert 2235

Tested At: 21 Aug 23 20:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 971 day(s)\
Subject: CN=Blackfoot Communications SHAKEN Cert 2235, O=Blackfoot Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/73aa28faf4546c63b3b20c530d38004b43bbecd4

[View certificate details](https://understandingwebpki.com/?cert=MIICbDCCAhGgAwIBAgIQZhf9PrscAJGLvlYsldQNIzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDQxOTE0MTkzNFoXDTI2MDQxODE0MTkzNFowZDELMAkGA1UEBhMCVVMxITAfBgNVBAoMGEJsYWNrZm9vdCBDb21tdW5pY2F0aW9uczEyMDAGA1UEAwwpQmxhY2tmb290IENvbW11bmljYXRpb25zIFNIQUtFTiBDZXJ0IDIyMzUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS3TreOkJtIXQpz5XTLheeQ8tYiLUCpmf3y7rcxh7ZL1dM9yCasBGTGDg5JNz3bxGckJn%2FkFdx7bmSwH9bD2Ko2o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQyMjM1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFNqS%2BfOypQJemYXsY5hH%2F3bbEMGEMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQC%2BOBDFkTSOFJaP7288spgh6LPDwUua%2Fz%2FHsOBvNPkgTAIhAMYD6JFa%2FcWoZFLg%2FSs6O74PQlQoZRgdmzlT4kWts9FB)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2235' |


Generated: 21 Aug 23 20:18 UTC