# STIR/SHAKEN CA Ecosystem Compliance

## Certificate West Kentucky Rural SHAKEN Cert 0421

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 548 day(s)\
Subject: CN=West Kentucky Rural SHAKEN Cert 0421, O=West Kentucky Rural Telephone Cooperative, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/e8ac880da62a959a70d1886399c240bdd0a37078

[View certificate details](https://x509.io/?cert=MIICdzCCAh2gAwIBAgIQeZhOON8Klyz79JN4I1pI6TAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDQwNjEwMTgxM1oXDTI2MDQwNTEwMTgxM1owcDELMAkGA1UEBhMCVVMxMjAwBgNVBAoMKVdlc3QgS2VudHVja3kgUnVyYWwgVGVsZXBob25lIENvb3BlcmF0aXZlMS0wKwYDVQQDDCRXZXN0IEtlbnR1Y2t5IFJ1cmFsIFNIQUtFTiBDZXJ0IDA0MjEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASXMEraouzmnJdfE5YfSJ13aIkujjt3jUXtKInL1VmocLegrGW%2FBO%2BUI%2B1DQLJavo758YQjQy2Mv%2BOkCjDzE95Zo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwNDIxMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFJ2KkG7HfA3Hh4FWDQASSAQONhsNMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQDdkBLhQXpkQzsDCybHgmQF21PvmM%2FLqGVVocbriCPHFAIgH2x8t5iTAkayiL4g0r%2F2Ow%2FFR1T0KitIFRNtsd46PQ0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0421', but common name is 'West Kentucky Rural SHAKEN Cert 0421' |


Generated: 04 Oct 24 16:29 UTC