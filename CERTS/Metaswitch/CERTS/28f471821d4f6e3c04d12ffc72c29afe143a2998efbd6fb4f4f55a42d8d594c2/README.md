# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Ritter Communications SHAKEN cert 095A

Tested At: 21 Aug 23 20:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 833 day(s)\
Subject: CN=Ritter Communications SHAKEN cert 095A, O=Ritter Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/ed02c72bad69b41f5a672ea7e1eb3f67a7ad3be0

[View certificate details](https://understandingwebpki.com/?cert=MIICZTCCAgugAwIBAgIQWA5bG%2FwZeoV4TopjVTLO%2FjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMTIwMjE3MzMyM1oXDTI1MTIwMTE3MzMyM1owXjELMAkGA1UEBhMCVVMxHjAcBgNVBAoMFVJpdHRlciBDb21tdW5pY2F0aW9uczEvMC0GA1UEAwwmUml0dGVyIENvbW11bmljYXRpb25zIFNIQUtFTiBjZXJ0IDA5NUEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT1qn%2F9cBWWRfx%2BCvrvY5Gd8wAEk8JRaKHkJVtMofYx%2BfE5oRu%2FYkK50R9obpilgpoyh7CE0tCrc8g95aehPSn7o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwOTVBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFGBeZS66YMdnR7zup7WSQlzWpoxSMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCVcp8uT51xRLS3IWFj5v1b6oS6kZN6nDEWEpcQBINakwIgKQi3Su7b7XX6rPGCeEi%2B4uizXcVp5dQNxpfKcXPVSt4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 095A' |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


Generated: 21 Aug 23 20:18 UTC