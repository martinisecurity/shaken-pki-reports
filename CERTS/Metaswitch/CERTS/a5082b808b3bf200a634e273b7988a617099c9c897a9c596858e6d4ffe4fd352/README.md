# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Sonic Telecom SHAKEN cert 433E

Tested At: 01 Nov 22 07:32 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 607 day(s)\
Subject: CN=Sonic Telecom SHAKEN cert 433E, O=Sonic Telecom, ST=CA, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://crs.qcall.sonic.net/certs/soniccertchain.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICYjCCAgigAwIBAgIQPCOrN3MeUy1qrCEGFArOhTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYzMDE2NTUwOVoXDTI0MDYyOTE2NTUwOVowWzELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQKDA1Tb25pYyBUZWxlY29tMScwJQYDVQQDDB5Tb25pYyBUZWxlY29tIFNIQUtFTiBjZXJ0IDQzM0UwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ6GeyQqFKg0VoXenwReqAbSPAUj0yc%2Fvbfb1UXbA6lGYGq%2BmgF%2FddnFitygzQwUygMVXcFwHcLVeZxl%2BtPIx1Ho4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ0MzNFMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFAhGRE8YHUsjai5BOfmK6U7QxtQvMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQD%2FtOUImTEmlFitBgjZSo4r7YVRuZrUlpvr9NfEOe3IqAIgGxrMGbQn9Ri8u8jnmfRsBE6btOt662Jtr97R9FRk63g%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 07:33:04