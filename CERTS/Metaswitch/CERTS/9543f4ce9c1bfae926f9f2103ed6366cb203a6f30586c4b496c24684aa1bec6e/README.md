# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Crosstel Tandem Inc SEPB SHAKEN Cert 357H

Tested At: 21 Aug 23 20:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 987 day(s)\
Subject: CN=Crosstel Tandem Inc SEPB SHAKEN Cert 357H, O=Crosstel Tandem Inc SEPB, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/e8dafa47260dac29a47f4ba89833e0b648af8afd

[View certificate details](https://understandingwebpki.com/?cert=MIICajCCAhGgAwIBAgIQcm2T%2BLp3lCO98qy5M12gJTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUwNTE1NDgyNloXDTI2MDUwNDE1NDgyNlowZDELMAkGA1UEBhMCVVMxITAfBgNVBAoMGENyb3NzdGVsIFRhbmRlbSBJbmMgU0VQQjEyMDAGA1UEAwwpQ3Jvc3N0ZWwgVGFuZGVtIEluYyBTRVBCIFNIQUtFTiBDZXJ0IDM1N0gwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARpewb3LmLno%2Ftuh1YY9wwBuGHaS%2B6OuIFush5FNA1NTUQOsh%2Bh383KZ7OFt%2F%2F2dkOBDqyMleN235FE3%2BLXzozAo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQzNTdIMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFKtqeek7VXNZQns%2BPlO7dmGEWg9OMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIAwvIXEG8%2FSwEqlfBi9402u9pG6%2Fkz4oLhVGBmL%2BdxShAiBTjzN%2B49TwTemHJwLgSSua9T0Fev8jGwLzTkwU0DZjxA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 357H' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |


Generated: 21 Aug 23 20:18 UTC