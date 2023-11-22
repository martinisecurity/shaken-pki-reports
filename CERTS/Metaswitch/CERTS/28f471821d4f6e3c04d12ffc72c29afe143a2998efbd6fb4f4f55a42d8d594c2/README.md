# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Ritter Communications SHAKEN cert 095A

Tested At: 22 Nov 23 03:16 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 741 day(s)\
Subject: CN=Ritter Communications SHAKEN cert 095A, O=Ritter Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/7721a7f2c31bb0f15ac5d46257b8c3bee1bbae6f

[View certificate details](https://understandingwebpki.com/?cert=MIICZTCCAgugAwIBAgIQWA5bG%2FwZeoV4TopjVTLO%2FjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMTIwMjE3MzMyM1oXDTI1MTIwMTE3MzMyM1owXjELMAkGA1UEBhMCVVMxHjAcBgNVBAoMFVJpdHRlciBDb21tdW5pY2F0aW9uczEvMC0GA1UEAwwmUml0dGVyIENvbW11bmljYXRpb25zIFNIQUtFTiBjZXJ0IDA5NUEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT1qn%2F9cBWWRfx%2BCvrvY5Gd8wAEk8JRaKHkJVtMofYx%2BfE5oRu%2FYkK50R9obpilgpoyh7CE0tCrc8g95aehPSn7o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwOTVBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFGBeZS66YMdnR7zup7WSQlzWpoxSMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCVcp8uT51xRLS3IWFj5v1b6oS6kZN6nDEWEpcQBINakwIgKQi3Su7b7XX6rPGCeEi%2B4uizXcVp5dQNxpfKcXPVSt4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 095A' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Nov 23 03:38 UTC