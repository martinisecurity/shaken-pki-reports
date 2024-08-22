# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Northeast Communications of Wisconsin SHAKEN Cert 6692

Tested At: 22 Aug 24 15:19 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 904 day(s)\
Subject: CN=Northeast Communications of Wisconsin SHAKEN Cert 6692, O=Northeast Communications of Wisconsin\\ , C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/b169a4a9dbb9ca311b83e2c9b0c04f1bd80e00d3

[View certificate details](https://x509.io/?cert=MIIChjCCAiygAwIBAgIQGc5gJGav4NdAXtfCJtLnozAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDIxMzEwMzkxNVoXDTI3MDIxMjEwMzkxNVowfzELMAkGA1UEBhMCVVMxLzAtBgNVBAoMJk5vcnRoZWFzdCBDb21tdW5pY2F0aW9ucyBvZiBXaXNjb25zaW4gMT8wPQYDVQQDDDZOb3J0aGVhc3QgQ29tbXVuaWNhdGlvbnMgb2YgV2lzY29uc2luIFNIQUtFTiBDZXJ0IDY2OTIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATL4PA%2BfEhBkJ64BPg81zhg7RBG6UuYi2CLxJRKkObXzpeOjGY8Is8tM7DOLvyKeM2Z17AOIaIlo7VDnBjUNEIno4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ2NjkyMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFLnn3RmXyv1W3o8BLasQoX0PKIIuMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQD4pnUfe5C8EdQiePfZiKwXAntIVP9GInjA0lMoES%2BzYwIgJow8Sr1%2B8BMwvIEISdsEkTzGh%2FGpKlgvvC1P55lWrfs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 6692', but common name is 'Northeast Communications of Wisconsin SHAKEN Cert 6692' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 22 Aug 24 16:06 UTC