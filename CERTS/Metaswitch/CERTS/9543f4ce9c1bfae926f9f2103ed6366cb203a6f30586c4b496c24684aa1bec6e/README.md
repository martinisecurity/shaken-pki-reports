# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Crosstel Tandem Inc SEPB SHAKEN Cert 357H

Tested At: 28 Nov 23 10:18 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 889 day(s)\
Subject: CN=Crosstel Tandem Inc SEPB SHAKEN Cert 357H, O=Crosstel Tandem Inc SEPB, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/15812c8f761ade70a623766b81fb851b88f684b3

[View certificate details](https://understandingwebpki.com/?cert=MIICajCCAhGgAwIBAgIQcm2T%2BLp3lCO98qy5M12gJTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUwNTE1NDgyNloXDTI2MDUwNDE1NDgyNlowZDELMAkGA1UEBhMCVVMxITAfBgNVBAoMGENyb3NzdGVsIFRhbmRlbSBJbmMgU0VQQjEyMDAGA1UEAwwpQ3Jvc3N0ZWwgVGFuZGVtIEluYyBTRVBCIFNIQUtFTiBDZXJ0IDM1N0gwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARpewb3LmLno%2Ftuh1YY9wwBuGHaS%2B6OuIFush5FNA1NTUQOsh%2Bh383KZ7OFt%2F%2F2dkOBDqyMleN235FE3%2BLXzozAo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQzNTdIMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFKtqeek7VXNZQns%2BPlO7dmGEWg9OMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIAwvIXEG8%2FSwEqlfBi9402u9pG6%2Fkz4oLhVGBmL%2BdxShAiBTjzN%2B49TwTemHJwLgSSua9T0Fev8jGwLzTkwU0DZjxA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 357H', but common name is 'Crosstel Tandem Inc SEPB SHAKEN Cert 357H' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 28 Nov 23 10:53 UTC