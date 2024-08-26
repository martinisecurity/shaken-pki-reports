# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Altafiber SHAKEN Cert 600F

Tested At: 26 Aug 24 18:01 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 924 day(s)\
Subject: CN=Altafiber SHAKEN Cert 600F, O=Altafiber, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/5bf43a66fdefa14ec681ce439ba5f70dce880ac3

[View certificate details](https://x509.io/?cert=MIICTjCCAfOgAwIBAgIQTAbtl6KEXOqIEqhTL2fNBzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDMwODE2MjcxNloXDTI3MDMwODE2MjcxNlowRjELMAkGA1UEBhMCVVMxEjAQBgNVBAoMCUFsdGFmaWJlcjEjMCEGA1UEAwwaQWx0YWZpYmVyIFNIQUtFTiBDZXJ0IDYwMEYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ0yB6R0fK1NPedBw19%2Br%2BddxETfcugBreaj4UB4EeMPOSt%2B%2BPmyGsu8irwCxOQypiegMi8UxbsJjDRmqo0kXU6o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ2MDBGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFN3MibCB5PiaOKPTuLg1jcMOlWZYMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQCIGRgT8zlnDzW%2Bms9GBMCkhq9Gn1Y0wC15Voy8rVbcvgIhAOuOHyMWgRGIEolENuo5VE4sbCvUenB2baPIK6664Sei)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 600F', but common name is 'Altafiber SHAKEN Cert 600F' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 26 Aug 24 18:49 UTC