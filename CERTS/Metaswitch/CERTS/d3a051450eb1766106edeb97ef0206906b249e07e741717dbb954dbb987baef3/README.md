# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Buckeye SHAKEN Cert 7608

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 1070 day(s)\
Subject: CN=Buckeye SHAKEN Cert 7608, O=Buckeye\\ , C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/80ac3a7475bf118d59c0c864b846ca67cf17d5e9

[View certificate details](https://understandingwebpki.com/?cert=MIICSzCCAfCgAwIBAgIQcFEXxy4OUOKypaDdAWBYtDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDExNzE3MjgzOFoXDTI3MDExNjE3MjgzOFowQzELMAkGA1UEBhMCVVMxETAPBgNVBAoMCEJ1Y2tleWUgMSEwHwYDVQQDDBhCdWNrZXllIFNIQUtFTiBDZXJ0IDc2MDgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQXQ5iAbekWx%2Fu8C3Ae1uLEuBOOyQ00NnkNeB4QYkd6WaEXdOZ%2Fo3ZMxv%2F8JHC7SilmNiiVPBc%2BSeLtPehRMSX6o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3NjA4MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFMo5xxOnCg9KvVqUVLHfpGYegk5FMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQC0jn79jobRpxKr3oCDday%2FQnGrxRBj1jkRhPVWgwFYEgIhAL1ItHb3NJKViJtYw%2B8Q8nK5K6V0cF%2BvANGMBFUtu7IZ)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7608', but common name is 'Buckeye SHAKEN Cert 7608' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 12 Feb 24 17:02 UTC