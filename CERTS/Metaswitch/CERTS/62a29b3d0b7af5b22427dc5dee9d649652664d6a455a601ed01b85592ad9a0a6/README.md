# STIR/SHAKEN CA Ecosystem Compliance

## Certificate West Kentucky Rural SHAKEN Cert 0421

Tested At: 21 Aug 23 20:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 958 day(s)\
Subject: CN=West Kentucky Rural SHAKEN Cert 0421, O=West Kentucky Rural Telephone Cooperative, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/fecae0c1f85acf27b0d6f8afbe4c90593fd6ecc7

[View certificate details](https://understandingwebpki.com/?cert=MIICdzCCAh2gAwIBAgIQeZhOON8Klyz79JN4I1pI6TAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDQwNjEwMTgxM1oXDTI2MDQwNTEwMTgxM1owcDELMAkGA1UEBhMCVVMxMjAwBgNVBAoMKVdlc3QgS2VudHVja3kgUnVyYWwgVGVsZXBob25lIENvb3BlcmF0aXZlMS0wKwYDVQQDDCRXZXN0IEtlbnR1Y2t5IFJ1cmFsIFNIQUtFTiBDZXJ0IDA0MjEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASXMEraouzmnJdfE5YfSJ13aIkujjt3jUXtKInL1VmocLegrGW%2FBO%2BUI%2B1DQLJavo758YQjQy2Mv%2BOkCjDzE95Zo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwNDIxMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFJ2KkG7HfA3Hh4FWDQASSAQONhsNMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQDdkBLhQXpkQzsDCybHgmQF21PvmM%2FLqGVVocbriCPHFAIgH2x8t5iTAkayiL4g0r%2F2Ow%2FFR1T0KitIFRNtsd46PQ0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0421' |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 21 Aug 23 20:18 UTC