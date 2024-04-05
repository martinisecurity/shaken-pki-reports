# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Telesystem SHAKEN Cert  786E

Tested At: 05 Apr 24 18:39 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 1009 day(s)\
Subject: CN=Telesystem SHAKEN Cert  786E, O=Telesystem, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/1119f8a66803e30896e315d1e9766b55ab10e7f9

[View certificate details](https://x509.io/?cert=MIICUDCCAfagAwIBAgIQTMcXQh%2B9U9bawchIV8WvFDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDExMDEwNDExNFoXDTI3MDEwOTEwNDExNFowSTELMAkGA1UEBhMCVVMxEzARBgNVBAoMClRlbGVzeXN0ZW0xJTAjBgNVBAMMHFRlbGVzeXN0ZW0gU0hBS0VOIENlcnQgIDc4NkUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAThFDteKNuzN4L2ekLDDrwGWiI5444Qatw%2FdreGNeF0uT30EA5wjsCFJ5Htk%2FdA8738Mv9IBcGmeEGsVTISMOXIo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3ODZFMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFGHvRqFdvIvgVUCOc6U7Ba9eahKkMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQC3qdI8ZgGsspTSoiE2JDo4tIgNiR9rSGnkNz8fR2w7WAIgMBgWkxoYavtKUPNcaqQv61uTNEebp9EyCLgsG0OOzgo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 786E', but common name is 'Telesystem SHAKEN Cert  786E' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 05 Apr 24 19:04 UTC