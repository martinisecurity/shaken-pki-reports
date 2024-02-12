# STIR/SHAKEN CA Ecosystem Compliance

## Certificate CAS Cable SHAKEN Cert 875F

Tested At: 12 Feb 24 18:54 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 860 day(s)\
Subject: CN=CAS Cable SHAKEN Cert 875F, O=CAS Cable, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/78696462dd41a95543996df2ba61d01c0d7274f9

[View certificate details](https://understandingwebpki.com/?cert=MIICTDCCAfOgAwIBAgIQXcgY%2FAkGwMUiTcPhCjdJADAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDYyMjEwMjAzMloXDTI2MDYyMTEwMjAzMlowRjELMAkGA1UEBhMCVVMxEjAQBgNVBAoMCUNBUyBDYWJsZTEjMCEGA1UEAwwaQ0FTIENhYmxlIFNIQUtFTiBDZXJ0IDg3NUYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARQcNAR3BpZHDuZ2qun4dQY1KBJYmTlmmmPxMDNa%2Fl0oFPH5jQaCVGMq7wvW6l%2BGGJzY%2FGOWEu8UMaRNFYhgg3So4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ4NzVGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFCYwIOXszyTz7cs0QDJ17wsi8fbRMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIDMTWyKLC10o8RwP%2BFYpGSXbKrczHKY0AhoptJ4W1ipGAiBX5VF0uAWQQs%2B4qRkqxEaS8wEWes4BdbKP0pAzHnk%2BZg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 875F', but common name is 'CAS Cable SHAKEN Cert 875F' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 12 Feb 24 19:26 UTC