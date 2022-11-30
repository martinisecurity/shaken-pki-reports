# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Sonic Telecom SHAKEN cert 433E

Tested At: 30 Nov 22 15:54 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 578 day(s)\
Subject: CN=Sonic Telecom SHAKEN cert 433E, O=Sonic Telecom, ST=CA, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://crs.qcall.sonic.net/certs/soniccertchain.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICYjCCAgigAwIBAgIQPCOrN3MeUy1qrCEGFArOhTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYzMDE2NTUwOVoXDTI0MDYyOTE2NTUwOVowWzELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQKDA1Tb25pYyBUZWxlY29tMScwJQYDVQQDDB5Tb25pYyBUZWxlY29tIFNIQUtFTiBjZXJ0IDQzM0UwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ6GeyQqFKg0VoXenwReqAbSPAUj0yc%2Fvbfb1UXbA6lGYGq%2BmgF%2FddnFitygzQwUygMVXcFwHcLVeZxl%2BtPIx1Ho4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ0MzNFMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFAhGRE8YHUsjai5BOfmK6U7QxtQvMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQD%2FtOUImTEmlFitBgjZSo4r7YVRuZrUlpvr9NfEOe3IqAIgGxrMGbQn9Ri8u8jnmfRsBE6btOt662Jtr97R9FRk63g%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 30 Nov 22 16:07 UTC