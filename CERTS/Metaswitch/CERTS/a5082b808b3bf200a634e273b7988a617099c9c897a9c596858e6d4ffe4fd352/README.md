# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Sonic Telecom SHAKEN cert 433E

Tested At: 28 Nov 23 20:12 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 214 day(s)\
Subject: CN=Sonic Telecom SHAKEN cert 433E, O=Sonic Telecom, ST=CA, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://crs.qcall.sonic.net/certs/soniccertchain.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICYjCCAgigAwIBAgIQPCOrN3MeUy1qrCEGFArOhTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYzMDE2NTUwOVoXDTI0MDYyOTE2NTUwOVowWzELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQKDA1Tb25pYyBUZWxlY29tMScwJQYDVQQDDB5Tb25pYyBUZWxlY29tIFNIQUtFTiBjZXJ0IDQzM0UwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ6GeyQqFKg0VoXenwReqAbSPAUj0yc%2Fvbfb1UXbA6lGYGq%2BmgF%2FddnFitygzQwUygMVXcFwHcLVeZxl%2BtPIx1Ho4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ0MzNFMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFAhGRE8YHUsjai5BOfmK6U7QxtQvMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQD%2FtOUImTEmlFitBgjZSo4r7YVRuZrUlpvr9NfEOe3IqAIgGxrMGbQn9Ri8u8jnmfRsBE6btOt662Jtr97R9FRk63g%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 20:21 UTC